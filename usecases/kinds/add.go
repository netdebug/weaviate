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
 */package kinds

import (
	"context"

	"github.com/creativesoftwarefdn/weaviate/entities/models"
	"github.com/creativesoftwarefdn/weaviate/entities/schema"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/usecases/kinds/validation"
	"github.com/go-openapi/strfmt"
)

type addAndGetRepo interface {
	getRepo
	addRepo
}

type addRepo interface {
	AddAction(ctx context.Context, class *models.Action, id strfmt.UUID) error
	AddThing(ctx context.Context, class *models.Thing, id strfmt.UUID) error
}

// TODO: Can we use the schema manager UC here instead of the "whole thing"?
type schemaManager interface {
	UpdatePropertyAddDataType(context.Context, kind.Kind, string, string, string) error
	GetSchema() schema.Schema
}

// AddAction Class Instance to the connected DB. If the class contains a network
// ref, it has a side-effect on the schema: The schema will be updated to
// include this particular network ref class.
func (m *Manager) AddAction(ctx context.Context, class *models.Action) (*models.Action, error) {
	unlock, err := m.locks.LockSchema()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.addActionToConnectorAndSchema(ctx, class)
}

func (m *Manager) addActionToConnectorAndSchema(ctx context.Context, class *models.Action) (*models.Action, error) {
	id, err := generateUUID()
	if err != nil {
		return nil, newErrInternal("could not generate id: %v", err)
	}
	class.ID = id

	err = m.validateAction(ctx, class)
	if err != nil {
		return nil, newErrInvalidUserInput("invalid action: %v", err)
	}

	err = m.addNetworkDataTypesForAction(ctx, class)
	if err != nil {
		return nil, newErrInternal("could not update schema for network refs: %v", err)
	}

	m.repo.AddAction(ctx, class, class.ID)
	if err != nil {
		return nil, newErrInternal("could not store action: %v", err)
	}

	return class, nil
}

func (m *Manager) validateAction(ctx context.Context, class *models.Action) error {
	// Validate schema given in body with the weaviate schema
	databaseSchema := schema.HackFromDatabaseSchema(m.schemaManager.GetSchema())
	return validation.ValidateActionBody(
		ctx, class, databaseSchema, m.repo, m.network, m.config)
}

// AddThing Class Instance to the connected DB. If the class contains a network
// ref, it has a side-effect on the schema: The schema will be updated to
// include this particular network ref class.
func (m *Manager) AddThing(ctx context.Context, class *models.Thing) (*models.Thing, error) {
	unlock, err := m.locks.LockSchema()
	if err != nil {
		return nil, newErrInternal("could not aquire lock: %v", err)
	}
	defer unlock()

	return m.addThingToConnectorAndSchema(ctx, class)
}

func (m *Manager) addThingToConnectorAndSchema(ctx context.Context, class *models.Thing) (*models.Thing, error) {
	id, err := generateUUID()
	if err != nil {
		return nil, newErrInternal("could not generate id: %v", err)
	}
	class.ID = id

	err = m.validateThing(ctx, class)
	if err != nil {
		return nil, newErrInvalidUserInput("invalid thing: %v", err)
	}

	err = m.addNetworkDataTypesForThing(ctx, class)
	if err != nil {
		return nil, newErrInternal("could not update schema for network refs: %v", err)
	}

	m.repo.AddThing(ctx, class, class.ID)
	if err != nil {
		return nil, newErrInternal("could not store thing: %v", err)
	}

	return class, nil
}

func (m *Manager) validateThing(ctx context.Context, class *models.Thing) error {
	// Validate schema given in body with the weaviate schema
	databaseSchema := schema.HackFromDatabaseSchema(m.schemaManager.GetSchema())
	return validation.ValidateThingBody(
		ctx, class, databaseSchema, m.repo, m.network, m.config)
}

func (m *Manager) addNetworkDataTypesForThing(ctx context.Context, class *models.Thing) error {
	refSchemaUpdater := newReferenceSchemaUpdater(ctx, m.schemaManager, m.network, class.Class, kind.Thing)
	return refSchemaUpdater.addNetworkDataTypes(class.Schema)
}

func (m *Manager) addNetworkDataTypesForAction(ctx context.Context, class *models.Action) error {
	refSchemaUpdater := newReferenceSchemaUpdater(ctx, m.schemaManager, m.network, class.Class, kind.Action)
	return refSchemaUpdater.addNetworkDataTypes(class.Schema)
}
