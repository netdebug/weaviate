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

package kinds

import (
	"testing"

	"github.com/creativesoftwarefdn/weaviate/entities/filters"
	"github.com/creativesoftwarefdn/weaviate/entities/schema"
	"github.com/creativesoftwarefdn/weaviate/entities/schema/kind"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetMetaParamsHashing(t *testing.T) {
	params := func() GetMetaParams {
		return GetMetaParams{
			Analytics:        filters.AnalyticsProps{UseAnaltyicsEngine: true},
			IncludeMetaCount: true,
			ClassName:        schema.ClassName("MyBestClass"),
			Filters:          nil,
			Kind:             kind.Thing,
			Properties: []MetaProperty{
				MetaProperty{
					Name:                schema.PropertyName("bestprop"),
					StatisticalAnalyses: []StatisticalAnalysis{Count},
				},
			},
		}
	}
	hash := func() string { return "18c6e361737a82cb99c5223fb2a61ba7" }

	t.Run("it generates a hash", func(t *testing.T) {
		p := params()
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.Equal(t, h, hash())
	})

	t.Run("it generates the same hash if analytical meta props are changed", func(t *testing.T) {
		p := params()
		p.Analytics.ForceRecalculate = true
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.Equal(t, hash(), h)
	})

	t.Run("it generates a different hash if a meta prop is changed", func(t *testing.T) {
		p := params()
		p.Properties[0].StatisticalAnalyses[0] = Maximum
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.NotEqual(t, hash(), h)
	})

	t.Run("it generates a different hash if where filter is added", func(t *testing.T) {
		p := params()
		p.Filters = &filters.LocalFilter{Root: &filters.Clause{Value: &filters.Value{Value: "foo"}}}
		h, err := p.AnalyticsHash()
		require.Nil(t, err)
		assert.NotEqual(t, hash(), h)
	})
}
