package http

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	defaultTimeout = 10 * time.Second
)

func HttpDoTimeout(requestBody []byte, method string, requestURI string, headers map[string]string, timeout time.Duration) ([]byte, int, error) {

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

	// time.Second * 20
	fc := &fasthttp.Client{}
	err := fc.DoTimeout(req, resp, timeout)

	var respBytes []byte
	statusCode := 0
	if err == nil {
		respBytes = append(respBytes, resp.Body()...)
		statusCode = resp.StatusCode()
	}

	return respBytes, statusCode, errors.WithStack(err)
}

func HttpGet(url string, headers map[string]string) ([]byte, int, error) {

	return HttpDoTimeout(nil, fasthttp.MethodGet, url, headers, defaultTimeout)

}

func HttpPost(url string, reqBody []byte, headers map[string]string) ([]byte, int, error) {

	return HttpDoTimeout(reqBody, fasthttp.MethodPost, url, headers, defaultTimeout)
}
