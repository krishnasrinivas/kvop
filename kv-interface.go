package main

import (
	"errors"
	"path"
	"strings"
)

type KVAPI interface {
	Put(key string, value []byte) error
	Get(keyStr string, value []byte) ([]byte, error)
	Delete(key string) error
	List(prefix string) ([]string, error)
}

var errDiskNotFound = errors.New("disk not found")
var errValueTooLong = errors.New("value too long")
var errFileNotFound = errors.New("file not found")

const kvDataDir = "data/"
const kvMetaDir = "meta/"

func pathJoin(elem ...string) string {
	trailingSlash := ""
	if len(elem) > 0 {
		if strings.HasSuffix(elem[len(elem)-1], "/") {
			trailingSlash = "/"
		}
	}
	return path.Join(elem...) + trailingSlash
}
