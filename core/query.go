package core

import (
	"DeepLXByLinuxdo/storage"
	"DeepLXByLinuxdo/util"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nalgeon/redka"
	"log"
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
	Src  string `json:"src"`
	Dst  string `json:"dst"`
	From string `json:"from"`
}

func QueryTranslationResult(src, source_lang, target_lang string) (a Answer, err error) {
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
	if cached, Err := storage.GetDatabase().Str().Get(src); Err != nil {
		if errors.Is(Err, redka.ErrNotFound) {
			log.Println("查询sqlite的错误:未找到缓存")
		}
	} else {
		log.Println("从缓存中获取")
		return Answer{
			Src:  src,
			Dst:  cached.String(),
			From: "cache",
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
	err = storage.GetDatabase().Str().Set(src, result.Data)
	if err != nil {
		return Answer{}, err
	} else {
		log.Printf("写入缓存")
	}
	return Answer{
		Src:  src,
		Dst:  result.Data,
		From: "search",
	}, nil
}
