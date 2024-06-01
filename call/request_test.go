package request

import (
	"net/http"
	"net/url"
	"testing"
)

func TestParseParams(t *testing.T) {
	testCases := []struct {
		name           string
		method         string
		url            string
		user           string
		header         string
		data           string
		expectedParams *RequestParams
		expectedError  error
	}{
		{
			name:   "Valid request",
			method: http.MethodGet,
			url:    "api/v1/users",
			user:   "username:password",
			header: "Content-Type:application/json,X-Custom-Header:value",
			data:   `{"name":"John Doe"}`,
			expectedParams: &RequestParams{
				Method: http.MethodGet,
				URL:    &url.URL{Path: "api/v1/users"},
				Header: http.Header{
					"Content-Type":    {"application/json"},
					"X-Custom-Header": {"value"},
				},
				Auth: struct {
					Username string
					Password string
				}{
					Username: "username",
					Password: "password",
				},
				Body: `{"name":"John Doe"}`,
			},
			expectedError: nil,
		},
		{
			name:           "Missing method",
			method:         "",
			url:            "api/v1/users",
			expectedParams: nil,
			expectedError: &ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "未指定 HTTP 方法",
			},
		},
		{
			name:           "Missing URL",
			method:         http.MethodGet,
			url:            "",
			expectedParams: nil,
			expectedError: &ErrorResponse{
				StatusCode: http.StatusBadRequest,
				Message:    "未指定请求 URL",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params, err := ParseParams(tc.method, tc.url, tc.user, tc.header, tc.data)
			if tc.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error, but got nil")
					return
				}
				if _, ok := err.(*ErrorResponse); !ok {
					t.Errorf("Expected ErrorResponse, but got %T", err)
					return
				}
				if err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
					return
				}
				if !requestParamsEqual(params, tc.expectedParams) {
					t.Errorf("Expected params %+v, but got %+v", tc.expectedParams, params)
				}
			}
		})
	}
}

func requestParamsEqual(p1, p2 *RequestParams) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}
	if p1.Method != p2.Method {
		return false
	}
	if !urlEqual(p1.URL, p2.URL) {
		return false
	}
	if !headerEqual(p1.Header, p2.Header) {
		return false
	}
	if p1.Auth.Username != p2.Auth.Username || p1.Auth.Password != p2.Auth.Password {
		return false
	}
	if p1.Body != p2.Body {
		return false
	}
	return true
}

func urlEqual(u1, u2 *url.URL) bool {
	if u1 == nil && u2 == nil {
		return true
	}
	if u1 == nil || u2 == nil {
		return false
	}
	return u1.String() == u2.String()
}

func headerEqual(h1, h2 http.Header) bool {
	if len(h1) != len(h2) {
		return false
	}
	for k, v1 := range h1 {
		v2, ok := h2[k]
		if !ok || !stringSliceEqual(v1, v2) {
			return false
		}
	}
	return true
}

func stringSliceEqual(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
