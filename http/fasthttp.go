package http

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
	"net"
	"net/url"
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

	dialer := &net.Dialer{
		Timeout:   defaultTimeout,
		KeepAlive: defaultTimeout,
	}

	fc := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return dialer.Dial("tcp", addr)
		},
		MaxIdleConnDuration: 90 * time.Second,
	}

	if bProxy {
		proxyAddr := "http://127.0.0.1:10808"
		proxyURL, err := url.Parse(proxyAddr)
		if err != nil {
			return nil, 0, errors.WithStack(err)
		}
		fc.Dial = func(addr string) (net.Conn, error) {
			proxyConn, err := dialer.Dial("tcp", proxyURL.Host)
			if err != nil {
				return nil, err
			}
			_, err = proxyConn.Write([]byte("CONNECT " + addr + " HTTP/1.1\r\nHost: " + addr + "\r\n\r\n"))
			if err != nil {
				proxyConn.Close()
				return nil, err
			}
			buf := make([]byte, 1024)
			proxyConn.SetReadDeadline(time.Now().Add(10 * time.Second))
			n, err := proxyConn.Read(buf)
			if err != nil {
				proxyConn.Close()
				return nil, err
			}
			if string(buf[:n])[:12] != "HTTP/1.1 200" {
				proxyConn.Close()
				return nil, errors.New("proxy connection failed")
			}
			return proxyConn, nil
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
