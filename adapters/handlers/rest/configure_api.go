/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */

package rest

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/creativesoftwarefdn/weaviate/adapters/handlers/rest/operations"
	"github.com/creativesoftwarefdn/weaviate/adapters/handlers/rest/state"
	"github.com/creativesoftwarefdn/weaviate/adapters/locks"
	"github.com/creativesoftwarefdn/weaviate/adapters/repos/etcd"
	"github.com/creativesoftwarefdn/weaviate/entities/models"
	"github.com/creativesoftwarefdn/weaviate/entities/schema"
	"github.com/creativesoftwarefdn/weaviate/usecases/config"
	"github.com/creativesoftwarefdn/weaviate/usecases/connstate"
	dblisting "github.com/creativesoftwarefdn/weaviate/usecases/connswitch"
	"github.com/creativesoftwarefdn/weaviate/usecases/kinds"
	"github.com/creativesoftwarefdn/weaviate/usecases/network/common/peers"
	schemaUC "github.com/creativesoftwarefdn/weaviate/usecases/schema"
	"github.com/creativesoftwarefdn/weaviate/usecases/telemetry"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
)

func makeConfigureServer(appState *state.State) func(*http.Server, string, string) {
	return func(s *http.Server, scheme, addr string) {
		// Add properties to the config
		appState.ServerConfig.Hostname = addr
		appState.ServerConfig.Scheme = scheme
	}
}

func configureAPI(api *operations.WeaviateAPI) http.Handler {
	appState, etcdClient := startupRoutine()

	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.OidcAuth = func(token string, scopes []string) (*models.Principal, error) {
		return appState.OIDC.ValidateAndExtract(token, scopes)
	}

	api.Logger = func(msg string, args ...interface{}) {
		appState.Logger.WithField("action", "restapi_management").Infof(msg, args...)
	}

	schemaRepo := etcd.NewSchemaRepo(etcdClient)
	connstateRepo := etcd.NewConnStateRepo(etcdClient)

	schemaManager, err := schemaUC.NewManager(appState.Connector, schemaRepo,
		appState.Locks, appState.Network, appState.Logger, appState)
	if err != nil {
		appState.Logger.
			WithField("action", "startup").WithError(err).
			Fatal("could not initialize schema manager")
		os.Exit(1)
	}
	kindsManager := kinds.NewManager(appState.Connector, appState.Locks,
		schemaManager, appState.Network, appState.ServerConfig)
	batchKindsManager := kinds.NewBatchManager(appState.Connector, appState.Locks,
		schemaManager, appState.Network, appState.ServerConfig)
	kindsTraverser := kinds.NewTraverser(appState.Locks, appState.Connector,
		appState)
	connstateManager, err := connstate.NewManager(connstateRepo, appState.Logger)
	if err != nil {
		appState.Logger.
			WithField("action", "startup").WithError(err).
			Fatal("could not initialize connector state manager")
		os.Exit(1)
	}

	appState.Connector.SetStateManager(connstateManager)
	appState.Connector.SetLogger(appState.Logger)
	appState.Connector.SetSchema(schemaManager.GetSchema())
	initialState := connstateManager.GetInitialState()
	appState.Connector.SetState(context.Background(), initialState)
	// allow up to 2 minutes for the connected db to be ready
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	if err := appState.Connector.Connect(ctx); err != nil {
		appState.Logger.
			WithField("action", "startup").WithError(err).
			Fatal("could not connect connector")
		os.Exit(1)
	}
	if err := appState.Connector.Init(context.Background()); err != nil {
		appState.Logger.
			WithField("action", "startup").WithError(err).
			Fatal("could not init connector")
		os.Exit(1)
	}

	updateSchemaCallback := makeUpdateSchemaCall(appState.Logger, appState, kindsTraverser)
	schemaManager.RegisterSchemaUpdateCallback(updateSchemaCallback)
	schemaManager.RegisterSchemaUpdateCallback(func(updatedSchema schema.Schema) {
		appState.Connector.SetSchema(updatedSchema)
	})

	appState.Network.RegisterUpdatePeerCallback(func(peers peers.Peers) {
		schemaManager.TriggerSchemaUpdateCallbacks()
	})
	appState.Network.RegisterSchemaGetter(schemaManager)

	setupSchemaHandlers(api, appState.TelemetryLogger, schemaManager)
	setupKindHandlers(api, appState.TelemetryLogger, kindsManager)
	setupKindBatchHandlers(api, appState.TelemetryLogger, batchKindsManager)
	setupC11yHandlers(api, appState.TelemetryLogger, appState)
	setupGraphQLHandlers(api, appState.TelemetryLogger, appState)
	setupMiscHandlers(api, appState.TelemetryLogger, appState.ServerConfig, appState.Network, schemaManager)

	api.ServerShutdown = func() {}
	configureServer = makeConfigureServer(appState)
	setupMiddlewares := makeSetupMiddlewares(appState)
	setupGlobalMiddleware := makeSetupGlobalMiddleware(appState)
	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// TODO: Split up and don't write into global variables. Instead return an appState
func startupRoutine() (*state.State, *clientv3.Client) {
	appState := &state.State{}
	// context for the startup procedure. (So far the only subcommand respecting
	// the context is the schema initialization, as this uses the etcd client
	// requiring context. Nevertheless it would make sense to have everything
	// that goes on in here pay attention to the context, so we can have a
	// "startup in x seconds or fail")
	ctx := context.Background()
	// The timeout is arbitrary we have to adjust it as we go along, if we
	// realize it is to big/small
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	logger := logger()
	appState.Logger = logger

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("created startup context, nothing done so far")

	// Load the config using the flags
	serverConfig := &config.WeaviateConfig{}
	appState.ServerConfig = serverConfig
	err := serverConfig.LoadConfig(connectorOptionGroup, logger)
	if err != nil {
		logger.WithField("action", "startup").WithError(err).Error("could not load config")
		logger.Exit(1)
	}

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("config loaded")

	appState.OIDC = configureOIDC(appState)
	appState.AnonymousAccess = configureAnonymousAccess(appState)

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("configured OIDC and anonymous access client")

	rawContextionary := loadContextionary(logger, appState.ServerConfig.Config)
	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("contextionary loaded")

	appState.Network = connectToNetwork(logger, appState.ServerConfig.Config)
	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("network configured")

		// Create the database connector using the config
	dbConnector, err := dblisting.NewConnector(serverConfig.Config.Database.Name, serverConfig.Config.Database.DatabaseConfig, serverConfig.Config)
	// Could not find, or configure connector.
	if err != nil {
		logger.WithField("action", "startup").WithError(err).Error("could not load config")
		logger.Exit(1)
	}
	appState.Connector = dbConnector

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("created db connector")

	// parse config store URL
	configURL := serverConfig.Config.ConfigurationStorage.URL
	configStore, err := url.Parse(configURL)
	if err != nil || configURL == "" {
		logger.WithField("action", "startup").WithField("url", configURL).
			WithError(err).Error("cannot parse config store URL")
		logger.Exit(1)
	}

	// Construct a distributed lock
	etcdClient, err := clientv3.New(clientv3.Config{Endpoints: []string{configStore.String()}})
	if err != nil {
		logger.WithField("action", "startup").
			WithError(err).Error("cannot construct distributed lock with etcd")
		logger.Exit(1)
	}

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("created etcd client")

	appState.TelemetryLogger = configureTelemetry(appState, etcdClient, logger)

	// new lock
	etcdLock, err := locks.NewEtcdLock(etcdClient, "/weaviate/schema-connector-rw-lock", logger)
	if err != nil {
		logger.WithField("action", "startup").
			WithError(err).Error("cannot create etcd-based lock")
		logger.Exit(1)
	}
	appState.Locks = etcdLock

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("created etcd session")
		// END remove

	logger.WithField("action", "startup").WithField("startup_time_left", timeTillDeadline(ctx)).
		Debug("initialized schema")

	// initialize the contextinoary with the rawContextionary, it will get updated on each schema update
	appState.Contextionary = rawContextionary
	appState.RawContextionary = rawContextionary

	return appState, etcdClient
}

func configureTelemetry(appState *state.State, etcdClient *clientv3.Client,
	logger logrus.FieldLogger) *telemetry.RequestsLog {
	// Extract environment variables needed for logging
	mainLog := telemetry.NewLog()
	loggingInterval := appState.ServerConfig.Config.Telemetry.Interval
	loggingURL := appState.ServerConfig.Config.Telemetry.RemoteURL
	loggingDisabled := appState.ServerConfig.Config.Telemetry.Disabled
	loggingDebug := appState.ServerConfig.Config.Debug

	if loggingURL == "" {
		loggingURL = telemetry.DefaultURL
	}

	if loggingInterval == 0 {
		loggingInterval = telemetry.DefaultInterval
	}

	// Propagate the peer name (if any), debug toggle and the enabled toggle to the requestsLog
	if appState.ServerConfig.Config.Network != nil {
		mainLog.PeerName = appState.ServerConfig.Config.Network.PeerName
	}
	mainLog.Debug = loggingDebug
	mainLog.Disabled = loggingDisabled

	// Initialize a non-expiring context for the reporter
	reportingContext := context.Background()
	// Initialize the reporter
	reporter := telemetry.NewReporter(reportingContext, mainLog, loggingInterval, loggingURL, loggingDisabled, loggingDebug, etcdClient, logger)

	// Start reporting
	go func() {
		reporter.Start()
	}()

	return mainLog
}

// logger does not parse the regular config object, as logging needs to be
// configured before the configuration is even loaded/parsed. We are thus
// "manually" reading the desired env vars and set reasonable defaults if they
// are not set.
//
// Defaults to log level info and json format
func logger() *logrus.Logger {
	logger := logrus.New()
	if os.Getenv("LOG_FORMAT") != "text" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	if os.Getenv("LOG_LEVEL") == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}
