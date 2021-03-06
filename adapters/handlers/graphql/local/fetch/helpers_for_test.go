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
package fetch

import (
	"context"

	testhelper "github.com/creativesoftwarefdn/weaviate/adapters/handlers/graphql/test/helper"
	"github.com/creativesoftwarefdn/weaviate/usecases/kinds"
	"github.com/stretchr/testify/mock"
)

type mockRequestsLog struct{}

func (m *mockRequestsLog) Register(first string, second string) {

}

type mockResolver struct {
	testhelper.MockResolver
}

func newMockResolver() *mockResolver {
	field := Build()
	mocker := &mockResolver{}
	mockLog := &mockRequestsLog{}
	mocker.RootFieldName = "Fetch"
	mocker.RootField = field
	mocker.RootObject = map[string]interface{}{
		"Resolver":    Resolver(mocker),
		"RequestsLog": mockLog,
	}
	return mocker
}

func (m *mockResolver) LocalFetchKindClass(ctx context.Context, params *kinds.FetchSearch) (interface{}, error) {
	args := m.Called(params)
	return args.Get(0), args.Error(1)
}

func (m *mockResolver) LocalFetchFuzzy(ctx context.Context, params kinds.FetchFuzzySearch) (interface{}, error) {
	args := m.Called(params)
	return args.Get(0), args.Error(1)
}

func newMockContextionary() *mockContextionary {
	return &mockContextionary{}
}

type mockContextionary struct {
	mock.Mock
}
