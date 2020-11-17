package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func inArray(needle interface{}, haystack interface{}) (exists bool) {
	exists = false

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				exists = true
				return
			}
		}
	}
	return
}

func myPanic(message string, theError error) {
	if BotConfig.DataStore != nil {
		BotConfig.DataStore.Close()
	}
	panic(fmt.Errorf(message, theError))
}

func baseURL(c *gin.Context) (url string) {
	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}
	url = scheme + c.Request.Host
	return
}

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func binaryFS(root string) *binaryFileSystem {
	return &binaryFileSystem{
		AssetFile(),
	}
}
