package core

import (
	"DeepLXByLinuxdo/model"
	"DeepLXByLinuxdo/storage"
	"fmt"
	"os"
	"testing"
)

func init() {
	storage.SetMysql()
	err := storage.GetMysql().Sync2(model.TranslateCache{})
	if err != nil {
		panic(err)
	}
}
func TestQueryTranslationResult(t *testing.T) {
	os.Setenv("TOKEN", "")
	result, err := QueryTranslationResult("hello", "", "")
	if err != nil {
		return
	}
	t.Logf("%+v\n", result)
}
func TestInsertOne(t *testing.T) {
	os.Setenv("TOKEN", "")
	result, err := QueryTranslationResult("hello", "", "")
	if err != nil {
		return
	}
	t.Logf("%+v\n", result)
	tc := new(model.TranslateCache)
	tc.Src = result.Src
	tc.Dst = result.Dst
	tc.Source_lang = result.SourceLang
	tc.Target_lang = result.TargetLang

	err = tc.CreateOne()
	if err != nil {
		panic(err)
	}
}
func TestFindOne(t *testing.T) {
	cached := model.TranslateCache{
		Src:         "hello",
		Source_lang: "auto",
		Target_lang: "zh",
	}
	cached.FindCache()
	fmt.Printf("%+v\n", cached)
}
