package indexer

import "errors"

// IsErrInvalidObjectHash Check error is ErrInvalidObjectHash
func IsErrInvalidObjectHash(e error) bool {
	return errors.Is(e, ErrInvalidObjHash)
}

// ErrInvalidObjHash Error for when given object hash is invalid
var ErrInvalidObjHash = errors.New("error: invalid object hash")
