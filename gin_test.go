package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	r := setup()
	// 创建一个 HTTP 请求（GET /hello）
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 验证响应码
	assert.Equal(t, 200, w.Code)

	// 验证响应内容
	expected := `{"message":"Hello World"}`
	assert.JSONEq(t, expected, w.Body.String())
}