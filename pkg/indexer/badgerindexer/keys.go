package badgerindexer

import (
	"encoding/base64"
	"path/filepath"
	"strconv"
)

// KeyObjectToURI Generate a Key that references object hash to uri
func KeyObjectToURI(objHash string, optIdx ...uint64) []byte {
	var optIdxVal string
	if optIdx != nil {
		optIdxVal = strconv.FormatUint(optIdx[0], 10)
	}
	return []byte(filepath.Join("/objects/", objHash, "/references", optIdxVal))
}

// KeyURIToObject Generate a Key that references uri to objects
func KeyURIToObject(uri string, isTotal bool) []byte {
	var totalPath string
	if isTotal {
		totalPath = "/total"
	}
	return []byte(filepath.Join("/uris/", base64.StdEncoding.EncodeToString([]byte(uri)), totalPath))
}
