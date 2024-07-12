package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Тестируем count > totalCount
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=6&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, list, totalCount)
}
func TestMainHandlerWhenCityIsNotSupported(t *testing.T) {
	// Тестируем город в параметре `city` не поддерживается
	req := httptest.NewRequest("GET", "/cafe?count=2&city=london", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

func TestMainHandlerWhenOkAndBodyNotEmpty(t *testing.T) {
	// Тестируем код ответа ОК, тело ответа не пустое
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil) // здесь нужно создать запрос к сервису
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body.String())

}
