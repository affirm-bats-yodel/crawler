package indexer

import (
	"context"
	"net/url"
)

// IsObjHashValid Check objHash is valid
//
// It just checking objHash length is 64 Bytes
func IsObjHashValid(objHash string) bool {
	return len(objHash) == 64
}

// IsValidURI Check uri is valid
//
// - try to parse given uri
//
// - check parsed uri's scheme is starting
// with http:// or https://
func IsValidURI(uri string) bool {
	u, err := url.Parse(uri)
	if err != nil {
		return false
	}
	// check scheme is valid
	//
	// only accepts starting with
	// http:// or https://
	switch u.Scheme {
	case "http", "https":
		return true
	default:
		return false
	}
}

// IsContextCancelled Check given context is cancelled
func IsContextCancelled(ctx context.Context) bool {
	return ctx.Err() != nil
}
