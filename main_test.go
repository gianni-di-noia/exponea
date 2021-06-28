package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"exponea.com/core"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	e := newRouter()
	req := httptest.NewRequest("GET", "/api/all?timeout=1000", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	var data []core.Work
	log.Printf("bytes %s", res.Body.Bytes())
	err := json.Unmarshal(res.Body.Bytes(), &data)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, 2, len(data))

}

func TestFirst(t *testing.T) {
	e := newRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/first?timeout=1000", nil)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

}
