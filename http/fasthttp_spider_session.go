package http

import (
	"bytes"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

// SpiderSession 爬虫会话，基于fasthttp实现
// 能力：长连接复用、会话Cookie自动维护、预访问获取初始Cookie、代理支持
// 注意：同一个Session实例仅用于同一个目标站点；多站点请创建多个SpiderSession实例
// 不同网站建议部署独立服务，彻底隔离会话状态
type SpiderSession struct {
	client    *fasthttp.Client
	cookieJar map[string]string // key:cookieName, value:cookieValue
	useProxy  bool
	proxyURL  *url.URL
}

const (
	spiderDefaultTimeout = 30 * time.Second
	spiderDialTimeout    = 30 * time.Second
	spiderMaxIdleDur     = 90 * time.Second
	defaultProxyAddr     = "http://127.0.0.1:10808"
)

// NewSpiderSession 创建爬虫会话
// bProxy 是否开启本地代理
func NewSpiderSession(bProxy bool) (*SpiderSession, error) {
	s := &SpiderSession{
		cookieJar: make(map[string]string),
		useProxy:  bProxy,
	}

	dialer := &net.Dialer{
		Timeout:   spiderDialTimeout,
		KeepAlive: spiderMaxIdleDur,
	}

	fc := &fasthttp.Client{
		MaxIdleConnDuration: spiderMaxIdleDur,
	}

	if bProxy {
		proxyU, err := url.Parse(defaultProxyAddr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		s.proxyURL = proxyU
		fc.Dial = func(addr string) (net.Conn, error) {
			proxyConn, err := dialer.Dial("tcp", proxyU.Host)
			if err != nil {
				return nil, err
			}
			// CONNECT隧道代理
			_, err = proxyConn.Write([]byte("CONNECT " + addr + " HTTP/1.1\r\nHost: " + addr + "\r\n\r\n"))
			if err != nil {
				_ = proxyConn.Close()
				return nil, err
			}
			buf := make([]byte, 1024)
			proxyConn.SetReadDeadline(time.Now().Add(10 * time.Second))
			n, err := proxyConn.Read(buf)
			if err != nil {
				_ = proxyConn.Close()
				return nil, err
			}
			if !strings.HasPrefix(string(buf[:n]), "HTTP/1.1 200") {
				_ = proxyConn.Close()
				return nil, errors.New("proxy connect failed, response not 200")
			}
			return proxyConn, nil
		}
	} else {
		fc.Dial = func(addr string) (net.Conn, error) {
			return dialer.Dial("tcp", addr)
		}
	}

	s.client = fc
	return s, nil
}

// do 内部请求逻辑
func (s *SpiderSession) do(method string, reqURI string, body []byte, headers map[string]string, timeout time.Duration) ([]byte, int, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(reqURI)
	req.Header.SetMethod(method)

	if method == fasthttp.MethodPost {
		req.SetBody(body)
	}

	// 设置业务传入headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 注入cookie到请求头
	s.injectCookies(req)

	// 执行请求
	err := s.client.DoTimeout(req, resp, timeout)
	if err != nil {
		return nil, 0, errors.WithStack(err)
	}

	// 解析响应Set‑Cookie，更新cookieJar
	s.parseSetCookie(resp)

	respBody := append([]byte(nil), resp.Body()...)
	statusCode := resp.StatusCode()
	return respBody, statusCode, nil
}

// injectCookies 将jar内cookie组装为Cookie请求头
func (s *SpiderSession) injectCookies(req *fasthttp.Request) {
	if len(s.cookieJar) == 0 {
		return
	}
	var buf bytes.Buffer
	for k, v := range s.cookieJar {
		if buf.Len() > 0 {
			buf.WriteByte(';')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(v)
	}
	req.Header.Set("Cookie", buf.String())
}

// parseSetCookie 解析响应全部Set‑Cookie头，更新cookieJar
func (s *SpiderSession) parseSetCookie(resp *fasthttp.Response) {
	// 获取全部 Set‑Cookie
	setCookieHeaders := resp.Header.PeekAll("Set‑Cookie")
	for _, raw := range setCookieHeaders {
		// raw示例：sessionId=abc123; Path=/; HttpOnly
		part := string(raw)
		firstSeg, _, _ := strings.Cut(part, ";")
		k, v, ok := strings.Cut(firstSeg, "=")
		if !ok {
			continue
		}
		s.cookieJar[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
}

// PreVisit 预访问页面，用于获取网站初始Set‑Cookie
// 模拟浏览器行为：先访问页面，再调用接口
func (s *SpiderSession) PreVisit(preUrl string, headers map[string]string) error {
	_, _, err := s.do(fasthttp.MethodGet, preUrl, nil, headers, spiderDefaultTimeout)
	return err
}

// HttpGet 会话GET请求
func (s *SpiderSession) HttpGet(url string, headers map[string]string) ([]byte, int, error) {
	return s.do(fasthttp.MethodGet, url, nil, headers, spiderDefaultTimeout)
}

// HttpPost 会话POST请求
func (s *SpiderSession) HttpPost(url string, reqBody []byte, headers map[string]string) ([]byte, int, error) {
	return s.do(fasthttp.MethodPost, url, reqBody, headers, spiderDefaultTimeout)
}

// GetCookieCount 获取当前会话存储cookie数量，用于调试
func (s *SpiderSession) GetCookieCount() int {
	return len(s.cookieJar)
}

// Close 清空cookie，任务结束调用
func (s *SpiderSession) Close() {
	s.cookieJar = make(map[string]string)
}
