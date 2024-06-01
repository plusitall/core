// request/request.go
package request

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RequestParams struct {
	Method string
	URL    *url.URL
	Header http.Header
	Auth   struct {
		Username string
		Password string
	}
	Body string
}

type ErrorResponse struct {
	StatusCode int
	Message    string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.StatusCode, e.Message)
}

func ParseParams(method, urlStr, user, header, data string) (*RequestParams, error) {
	// 验证 method 和 URL 是否为必填项
	if method == "" {
		return nil, &ErrorResponse{StatusCode: http.StatusBadRequest, Message: "未指定 HTTP 方法"}
	}
	if urlStr == "" {
		return nil, &ErrorResponse{StatusCode: http.StatusBadRequest, Message: "未指定请求 URL"}
	}

	params := &RequestParams{
		Method: method,
		URL:    parseURL(urlStr),
		Header: parseHeaders(header),
		Auth:   parseAuth(user),
		Body:   data,
	}

	if params.URL == nil {
		return nil, &ErrorResponse{StatusCode: http.StatusBadRequest, Message: "解析 URL 失败"}
	}

	return params, nil
}

func parseURL(urlStr string) *url.URL {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}
	return u
}

func parseHeaders(headerStr string) http.Header {
	headers := http.Header{}
	if headerStr != "" {
		headerPairs := strings.Split(headerStr, ",")
		for _, pair := range headerPairs {
			kv := strings.Split(pair, ":")
			if len(kv) == 2 {
				headers.Set(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
			}
		}
	}
	return headers
}

func parseAuth(authStr string) struct {
	Username string
	Password string
} {
	if authStr == "" {
		return struct {
			Username string
			Password string
		}{}
	}
	userPass := strings.Split(authStr, ":")
	if len(userPass) != 2 {
		return struct {
			Username string
			Password string
		}{}
	}
	return struct {
		Username string
		Password string
	}{
		Username: userPass[0],
		Password: userPass[1],
	}
}
