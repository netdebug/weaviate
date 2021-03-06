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
package meta

import (
	"testing"

	"github.com/creativesoftwarefdn/weaviate/usecases/kinds"
)

func Test_QueryBuilder_NumberProps(t *testing.T) {
	tests := testCases{
		testCase{
			name: "with only an int, with only count",
			inputProps: []kinds.MetaProperty{
				kinds.MetaProperty{
					Name:                "area",
					StatisticalAnalyses: []kinds.StatisticalAnalysis{kinds.Count},
				},
			},
			expectedQuery: `
				.union(
				  values("area").union(
					  count().project("count").project("area")
					)
				)
				.group().by(select(keys).unfold()).by(
					select(values).unfold().group()
					.by( select(keys).unfold())
					.by( select(values).unfold())
				)
			`,
		},

		testCase{
			name: "with only an int, with all props",
			inputProps: []kinds.MetaProperty{
				kinds.MetaProperty{
					Name: "area",
					StatisticalAnalyses: []kinds.StatisticalAnalysis{
						kinds.Mean, kinds.Type, kinds.Sum, kinds.Maximum, kinds.Minimum, kinds.Count,
					},
				},
			},
			expectedQuery: `
				.union(
				  values("area").union(
					  mean().project("mean").project("area"),
					  sum().project("sum").project("area"),
					  max().project("maximum").project("area"),
					  min().project("minimum").project("area"),
					  count().project("count").project("area")
					)
				)
				.group().by(select(keys).unfold()).by(
					select(values).unfold().group()
					.by( select(keys).unfold())
					.by( select(values).unfold())
				)
			`,
		},
	}

	tests.AssertQuery(t, nil)
}
