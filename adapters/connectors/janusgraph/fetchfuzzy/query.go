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
 */package fetchfuzzy

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/adapters/connectors/janusgraph/gremlin"
	"github.com/creativesoftwarefdn/weaviate/adapters/connectors/janusgraph/state"
	"github.com/creativesoftwarefdn/weaviate/entities/schema"
)

// Query prepares a Local->Fetch Query. Can be built with String(). Create with
// NewQuery() to be sure that all required properties are set
type Query struct {
	params     []string
	nameSource nameSource
	typeSource typeSource
}

// NewQuery is the preferred way to create a query
func NewQuery(p []string, ns nameSource, ts typeSource) *Query {
	return &Query{
		params:     p,
		nameSource: ns,
		typeSource: ts,
	}
}

type nameSource interface {
	GetMappedPropertyNames(rawProps []schema.ClassAndProperty) ([]state.MappedPropertyName, error)
	GetClassNameFromMapped(className state.MappedClassName) schema.ClassName
}

type typeSource interface {
	GetPropsOfType(string) []schema.ClassAndProperty
}

// String builds the query and returns it as a string
func (b *Query) String() (string, error) {

	predicates, err := b.predicates()
	if err != nil {
		return "", err
	}

	q := gremlin.New().
		Raw("g.V()").
		Or(predicates...).
		Limit(20).
		Raw(`.valueMap("uuid", "kind", "classId")`)

	return q.String(), nil
}

func (b *Query) predicates() ([]*gremlin.Query, error) {
	var result []*gremlin.Query
	props := b.typeSource.GetPropsOfType("string")

	mappedProps, err := b.nameSource.GetMappedPropertyNames(props)
	if err != nil {
		return nil, fmt.Errorf("could not get mapped names: %v", err)
	}

	// janusgraph does not allow safely allow more than aroudn 120 arguments, so
	// we must abort once we hit too many
	argsCounter := 0
	limit := 120

outer:
	for _, searchterm := range b.params {
		// Note that we are first iterating over the term, then over the prop. This
		// is because the terms are already an ordered list where the match with
		// the highest certainty is the first item. Janus also imposes a limit, so
		// we need to abort if we get too many predicates. Assuming we have 20 props
		// and 20 matches, by iterating over the matches first, this way, we'll be
		// able to cover about the 6 best matches on all props. If we split the
		// order of iteration, we would cover all matches, but completely ignore 14
		// out of 20 props.

		for _, prop := range mappedProps {
			if argsCounter >= limit {
				break outer
			}

			result = append(result,
				gremlin.New().Has(string(prop), gremlin.New().TextContainsFuzzy(searchterm)),
			)

			argsCounter++
		}
	}

	return result, nil
}
