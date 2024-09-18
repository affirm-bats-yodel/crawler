package indexer

import "errors"

// IsErrInvalidObjectHash Check error is ErrInvalidObjectHash
func IsErrInvalidObjectHash(e error) bool {
	return errors.Is(e, ErrInvalidObjHash)
}

// ErrInvalidObjHash Error for when given object hash is invalid
var ErrInvalidObjHash = errors.New("error: invalid object hash")

// IsErrInvalidURI Check error is ErrInvalidURI
func IsErrInvalidURI(e error) bool {
	return errors.Is(e, ErrInvalidURI)
}

// ErrInvalidURI Error for when given uri is invalid
var ErrInvalidURI = errors.New("error: invalid uri")
