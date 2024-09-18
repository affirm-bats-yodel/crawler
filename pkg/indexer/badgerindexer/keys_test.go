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
			expected: []byte("/objects/xxx/references"),
		},
		{
			name:     "KeyForReferenceIndex",
			objHash:  "xxx",
			optIdx:   []uint64{1},
			expected: []byte("/objects/xxx/references/1"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := badgerindexer.KeyObjectToURI(tc.objHash, tc.optIdx...)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func TestURIToObject(t *testing.T) {
	for _, tc := range []struct {
		name     string
		uri      string
		isTotal  bool
		expected []byte
	}{
		{
			name:     "KeyForTotalObjectCounts",
			uri:      "https://example.com",
			isTotal:  true,
			expected: []byte("/uris/aHR0cHM6Ly9leGFtcGxlLmNvbQ==/total"),
		},
		{
			name:     "KeyForObjectReferences",
			uri:      "https://example.com",
			isTotal:  false,
			expected: []byte("/uris/aHR0cHM6Ly9leGFtcGxlLmNvbQ=="),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			act := badgerindexer.KeyURIToObject(tc.uri, tc.isTotal)
			assert.Equal(t, tc.expected, act)
		})
	}
}
