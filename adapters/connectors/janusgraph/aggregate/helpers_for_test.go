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

package aggregate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/adapters/connectors/janusgraph/state"
	"github.com/creativesoftwarefdn/weaviate/entities/filters"
	"github.com/creativesoftwarefdn/weaviate/entities/models"
	"github.com/creativesoftwarefdn/weaviate/entities/schema"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/usecases/kinds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeNameSource struct{}

func (f *fakeNameSource) MustGetMappedPropertyName(className schema.ClassName,
	propName schema.PropertyName) state.MappedPropertyName {
	switch propName {
	case schema.PropertyName("name"):
		return "prop_1"
	case schema.PropertyName("population"):
		return "prop_2"
	case schema.PropertyName("inCountry"):
		return "prop_3"
	case schema.PropertyName("isCapital"):
		return "prop_4"
	case schema.PropertyName("inContinent"):
		return "prop_5"
	}
	return state.MappedPropertyName("prop_20")
}

func (f *fakeNameSource) MustGetMappedClassName(className schema.ClassName) state.MappedClassName {
	switch className {
	case schema.ClassName("Country"):
		return state.MappedClassName("class_18")
	case schema.ClassName("Continent"):
		return state.MappedClassName("class_19")
	default:
		panic(fmt.Sprintf("fake name source does not contain a fake for '%s'", className))
	}
}

type fakeTypeSource struct{}

func (f *fakeTypeSource) GetProperty(kind kind.Kind, className schema.ClassName,
	propName schema.PropertyName) (error, *models.SemanticSchemaClassProperty) {

	switch propName {
	case "isCapital":
		return nil, &models.SemanticSchemaClassProperty{DataType: []string{"bool"}}
	case "population":
		return nil, &models.SemanticSchemaClassProperty{DataType: []string{"int"}}
	case "area":
		return nil, &models.SemanticSchemaClassProperty{DataType: []string{"number"}}
	case "name":
		return nil, &models.SemanticSchemaClassProperty{DataType: []string{"string"}}
	case "dateOfFirstApperance":
		return nil, &models.SemanticSchemaClassProperty{DataType: []string{"date"}}
	case "inCountry":
		return nil, &models.SemanticSchemaClassProperty{DataType: []string{"Country"}}
	}

	return fmt.Errorf("fake type source does not have an implementation for prop '%s'", propName), nil
}

func (f *fakeTypeSource) FindPropertyDataType(dataType []string) (schema.PropertyDataType, error) {
	switch dataType[0] {
	case "bool":
		return &fakeDataType{dataType: schema.DataTypeBoolean}, nil
	case "int":
		return &fakeDataType{dataType: schema.DataTypeInt}, nil
	case "number":
		return &fakeDataType{dataType: schema.DataTypeNumber}, nil
	case "string":
		return &fakeDataType{dataType: schema.DataTypeString}, nil
	case "date":
		return &fakeDataType{dataType: schema.DataTypeDate}, nil
	case "Country":
		return &fakeDataType{dataType: schema.DataTypeCRef}, nil
	}

	return nil, fmt.Errorf("fake type source does not have an implementation for dataType '%v'", dataType)
}

type fakeDataType struct {
	dataType schema.DataType
}

func (p *fakeDataType) Kind() schema.PropertyKind {
	panic("not implemented")
}

func (p *fakeDataType) IsPrimitive() bool {
	return p.dataType != schema.DataTypeCRef
}

func (p *fakeDataType) AsPrimitive() schema.DataType {
	return p.dataType
}

func (p *fakeDataType) IsReference() bool {
	return p.dataType == schema.DataTypeCRef
}

func (p *fakeDataType) Classes() []schema.ClassName {
	return []schema.ClassName{"Country", "WeaviateB/Country"}
}

func (p *fakeDataType) ContainsClass(needle schema.ClassName) bool {
	panic("not implemented")
}

type fakeFilterSource struct {
	queryToReturn string
}

func (s *fakeFilterSource) String() (string, error) {
	return s.queryToReturn, nil
}

type testCase struct {
	name          string
	inputProps    []kinds.AggregateProperty
	inputGroupBy  *filters.Path
	expectedQuery string
}

type testCases []testCase

func (tests testCases) AssertQuery(t *testing.T, nameSource nameSource) {
	filter := &fakeFilterSource{}
	tests.AssertQueryWithFilterSource(t, nameSource, filter)
}

func (tests testCases) AssertQueryWithFilterSource(t *testing.T, nameSource nameSource,
	filterSource filterSource) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := &kinds.AggregateParams{
				Properties: test.inputProps,
				GroupBy:    test.inputGroupBy,
			}
			query, err := NewQuery(params, nameSource, &fakeTypeSource{}, filterSource).String()
			require.Nil(t, err, "should not error")
			assert.Equal(t, breakOnDot(stripAll(test.expectedQuery)), breakOnDot(stripAll(query)), "should match the query")
		})
	}
}

func stripAll(input string) string {
	input = strings.Replace(input, " ", "", -1)
	input = strings.Replace(input, "\t", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func breakOnDot(input string) string {
	input = strings.Replace(input, ".", "\n.", -1)
	return input
}
