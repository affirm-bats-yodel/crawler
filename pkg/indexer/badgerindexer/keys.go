package badgerindexer

import (
	"path/filepath"
	"strconv"
)

// KeyObjectToURI Generate a Key that references object hash to uri
func KeyObjectToURI(objHash string, optIdx ...uint64) []byte {
	var optIdxVal string
	if optIdx != nil {
		optIdxVal = strconv.FormatUint(optIdx[0], 10)
	}
	return []byte(filepath.Join("/", objHash, "/references", optIdxVal))
}
