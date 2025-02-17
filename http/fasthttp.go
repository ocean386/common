package http

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"net"
	"time"
)

var (
	defaultTimeout = 30 * time.Second
)

func HttpDoTimeout(bProxy bool, requestBody []byte, method string, requestURI string, headers map[string]string, timeout time.Duration) ([]byte, int, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(requestURI)
	req.Header.SetMethod(method)

	switch method {
	case "POST":
		req.SetBody(requestBody)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	fc := &fasthttp.Client{}
	if bProxy == true { // 设置代理
		proxyAddr := "127.0.0.1:10808" // 代理地址
		fc = &fasthttp.Client{
			Dial: func(addr string) (net.Conn, error) {
				// 通过代理服务器建立连接
				return fasthttp.Dial(proxyAddr)
			},
		}
	}

	err := fc.DoTimeout(req, resp, timeout)

	var respBytes []byte
	statusCode := 0
	if err == nil {
		respBytes = append(respBytes, resp.Body()...)
		statusCode = resp.StatusCode()
	}

	return respBytes, statusCode, errors.WithStack(err)
}

func HttpGet(bProxy bool, url string, headers map[string]string) ([]byte, int, error) {

	return HttpDoTimeout(bProxy, nil, fasthttp.MethodGet, url, headers, defaultTimeout)

}

func HttpPost(bProxy bool, url string, reqBody []byte, headers map[string]string) ([]byte, int, error) {

	return HttpDoTimeout(bProxy, reqBody, fasthttp.MethodPost, url, headers, defaultTimeout)
}
