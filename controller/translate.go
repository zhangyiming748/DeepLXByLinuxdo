package controller

import (
	"DeepLXByLinuxdo/core"
	"github.com/gin-gonic/gin"
)

type TranslateController struct{}

// 结构体必须大写 否则找不到
type TranslateReq struct {
	Src    string `json:"src"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type TranslateRep struct {
	Src    string `json:"src"`
	Dst    string `json:"dst"`
	Source string `json:"source"`
	Target string `json:"target"`
	From   string `json:"from"`
}

/*
curl --location --request POST 'http://127.0.0.1:8192/api/v1/translate' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "src":"hello"
	}'
*/
func (t TranslateController) TransWord(ctx *gin.Context) {
	req := new(TranslateReq)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("url = %s \nproxy = %s\n", req.URLs, req.Proxy)
	var rep TranslateRep
	result, err := core.QueryTranslationResult(req.Src, req.Source, req.Target)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	rep.Src = result.Src
	rep.Dst = result.Dst
	rep.Source = result.SourceLang
	rep.Target = result.TargetLang
	rep.From = result.From
	ctx.JSON(200, rep)
}
