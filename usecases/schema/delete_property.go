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
 */package schema

import (
	"context"
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/entities/schema"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/kind"
)

// DeleteActionProperty to an existing Action
func (m *Manager) DeleteActionProperty(ctx context.Context, class string, property string) error {
	return m.deleteClassProperty(ctx, class, property, kind.Action)
}

// DeleteThingProperty to an existing Thing
func (m *Manager) DeleteThingProperty(ctx context.Context, class string, property string) error {
	return m.deleteClassProperty(ctx, class, property, kind.Thing)
}

func (m *Manager) deleteClassProperty(ctx context.Context, className string, propName string, k kind.Kind) error {
	unlock, err := m.locks.LockSchema()
	if err != nil {
		return err
	}
	defer unlock()

	err = m.migrator.DropProperty(ctx, k, className, propName)
	if err != nil {
		return fmt.Errorf("could not migrate database schema: %v", err)
	}

	semanticSchema := m.state.SchemaFor(k)
	class, err := schema.GetClassByName(semanticSchema, className)
	if err != nil {
		return err
	}

	var propIdx = -1
	for idx, prop := range class.Properties {
		if prop.Name == propName {
			propIdx = idx
			break
		}
	}

	if propIdx == -1 {
		return fmt.Errorf("could not find property '%s' - it might have already been deleted?", propName)
	}

	class.Properties[propIdx] = class.Properties[len(class.Properties)-1]
	class.Properties[len(class.Properties)-1] = nil // to prevent leaking this pointer.
	class.Properties = class.Properties[:len(class.Properties)-1]

	err = m.saveSchema(ctx)
	if err != nil {
		return fmt.Errorf("could not persists schema change in configuration: %v", err)
	}

	return nil
}
