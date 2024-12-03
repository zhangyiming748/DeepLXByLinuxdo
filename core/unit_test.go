package core

import (
	"os"
	"testing"
)

func TestQueryTranslationResult(t *testing.T) {
	os.Setenv("TOKEN", "")
	result, err := QueryTranslationResult("hello", "", "")
	if err != nil {
		return
	}
	t.Logf("%+v\n", result)
}
