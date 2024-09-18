package badgerindexer_test

import (
	"testing"

	"github.com/affirm-bats-yodel/crawler/pkg/indexer/badgerindexer"
	"github.com/stretchr/testify/assert"
)

func TestKeyObjectToURI(t *testing.T) {
	for _, tc := range []struct {
		name     string
		objHash  string
		optIdx   []uint64
		expected []byte
	}{
		{
			name:     "KeyForReferenceCounts",
			objHash:  "xxx",
			optIdx:   nil,
			expected: []byte("/xxx/references"),
		},
		{
			name:     "KeyForReferenceIndex",
			objHash:  "xxx",
			optIdx:   []uint64{1},
			expected: []byte("/xxx/references/1"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := badgerindexer.KeyObjectToURI(tc.objHash, tc.optIdx...)
			assert.Equal(t, tc.expected, act)
		})
	}
}
