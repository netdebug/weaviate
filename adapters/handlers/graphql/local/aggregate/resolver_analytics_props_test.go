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
 */package aggregate

import (
	"testing"

	"github.com/creativesoftwarefdn/weaviate/entities/filters"
	"github.com/creativesoftwarefdn/weaviate/entities/schema"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/usecases/config"
	"github.com/creativesoftwarefdn/weaviate/usecases/kinds"
)

func Test_ExtractAnalyticsPropsFromAggregate(t *testing.T) {
	t.Parallel()

	query :=
		`{ Aggregate { Things { 
			Car(groupBy:["madeBy", "Manufacturer", "name"], useAnalyticsEngine: true, forceRecalculate: true) { 
				horsepower { mean } 
			} 
		} } }`

	analytics := filters.AnalyticsProps{
		UseAnaltyicsEngine: true,
		ForceRecalculate:   true,
	}
	cfg := config.Config{
		AnalyticsEngine: config.AnalyticsEngine{
			Enabled: true,
		},
	}

	expectedParams := &kinds.AggregateParams{
		Kind:      kind.Thing,
		ClassName: schema.ClassName("Car"),
		Properties: []kinds.AggregateProperty{
			{
				Name:        "horsepower",
				Aggregators: []kinds.Aggregator{kinds.MeanAggregator},
			},
		},
		GroupBy:   groupCarByMadeByManufacturerName(),
		Analytics: analytics,
	}
	resolver := newMockResolver(cfg)
	resolver.On("LocalAggregate", expectedParams).
		Return([]interface{}{}, nil).Once()

	resolver.AssertResolve(t, query)
}
