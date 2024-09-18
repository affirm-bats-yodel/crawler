package badgerindexer_test

import (
	"testing"

	"github.com/affirm-bats-yodel/crawler/pkg/indexer/badgerindexer"
	"github.com/stretchr/testify/assert"
)

func TestKeyObjectToURI(t *testing.T) {
	for _, tc := range []struct {
		objHash  string
		optIdx   []uint64
		expected string
	}{
		{
			objHash:  "xxx",
			optIdx:   nil,
			expected: "/xxx/references",
		},
		{
			objHash:  "xxx",
			optIdx:   []uint64{1},
			expected: "/xxx/references/1",
		},
	} {
		t.Run(tc.expected, func(t *testing.T) {
			act := badgerindexer.KeyObjectToURI(tc.objHash, tc.optIdx...)
			assert.Equal(t, tc.expected, act)
		})
	}
}
