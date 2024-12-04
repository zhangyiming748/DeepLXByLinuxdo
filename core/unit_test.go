package core

import (
	"DeepLXByLinuxdo/storage"
	"testing"
)

// go test -v -run TestQueryTranslationResult
func TestQueryTranslationResult(t *testing.T) {
	storage.SetDatabase()
	result, err := QueryTranslationResult("hello", "", "")
	if err != nil {
		return
	}
	t.Logf("%+v\n", result)
}
