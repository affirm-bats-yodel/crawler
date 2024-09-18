package indexer

// IsObjHashValid Check objHash is valid
//
// It just checking objHash length is 64 Bytes
func IsObjHashValid(objHash string) bool {
	return len(objHash) == 64
}
