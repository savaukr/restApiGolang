package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServerer_HandleMessage (t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/message", nil)
	s.handleMessage().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Message1")

}