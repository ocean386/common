package http

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type JsonResult struct {
	Status string      `json:"message"`
	Data   interface{} `json:"data"`
}

type Response struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data []interface{} `json:"data"`
}

func PrintResp(ctx *fasthttp.RequestCtx, state string, data interface{}) {

	ctx.SetStatusCode(200)
	ctx.SetContentType("application/json")

	res := JsonResult{
		Status: state,
		Data:   data,
	}

	bytes, err := jsoniter.Marshal(res)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}

	ctx.SetBody(bytes)
}
