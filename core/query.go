package core

import (
	"DeepLXByLinuxdo/model"
	"DeepLXByLinuxdo/util"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type DeeplxRep struct {
	Alternatives []string `json:"alternatives"`
	Code         int      `json:"code"`
	Data         string   `json:"data"`
	Id           int64    `json:"id"`
	Method       string   `json:"method"`
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
}
type Answer struct {
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
	Src        string `json:"src"`
	Dst        string `json:"dst"`
	From       string `json:"from"`
}

func QueryTranslationResult(src string, source_lang, target_lang string) (a Answer, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	if source_lang == "" {
		source_lang = "auto"
	}
	if target_lang == "" {
		target_lang = "zh"
	}
	data := map[string]string{
		"text":        src,
		"source_lang": source_lang,
		"target_lang": target_lang,
	}
	cached := model.TranslateCache{
		Src:         src,
		Source_lang: source_lang,
		Target_lang: target_lang,
	}

	has, _ := cached.FindCache()
	if has {
		return Answer{
			SourceLang: cached.Source_lang,
			TargetLang: cached.Target_lang,
			Src:        cached.Src,
			Dst:        cached.Dst,
			From:       "缓存",
		}, nil
	}
	token := os.Getenv("TOKEN")
	if token == "" {
		notfound := errors.New("没有找到deeplx的apikey环境变量$TOKEN")
		return Answer{}, notfound
	}
	uri := strings.Join([]string{"https://api.deeplx.org", token, "translate"}, "/")
	j, err := util.HttpPostJson(headers, data, uri)
	if err != nil {
		return Answer{}, err
	}
	fmt.Println(string(j))
	var result DeeplxRep
	json.Unmarshal(j, &result)
	a = Answer{
		SourceLang: source_lang,
		TargetLang: target_lang,
		Src:        src,
		Dst:        result.Data,
		From:       "查询",
	}
	return a, nil
}
