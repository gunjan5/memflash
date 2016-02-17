package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResp(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp := httptest.NewRecorder()
	Hello(resp, req)

	assert.Nil(t, err)
	assert.Equal(t, 200, resp.Code)

	if resp.Code != http.StatusOK {
		t.Fatalf("shit went south: expecting %v, got %v", http.StatusOK, resp.Code)
	}

}
