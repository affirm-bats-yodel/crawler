package indexer

import "context"

// Indexer Mapping Hash and URI
type Indexer interface {
	// Create Map Object Hash to URI
	//
	// - objHash: SHA 256 Based Object Hash
	//
	// - uri: URI
	Create(ctx context.Context, objHash string, uri string) error
	// GetByHash Get URIs by Object Hash
	//
	// - objHash: SHA 256 Based Object Hash
	//
	// - []string: List of URI Mapped by hash of object
	GetByHash(ctx context.Context, objHash string) ([]string, error)
	// GetByURI Get Object Hashes by URI
	//
	// - uri: URI to List Object Hashes
	//
	// - []string: List of Object Hashes
	GetByURI(ctx context.Context, uri string) ([]string, error)
}
