package main

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"io"
	"log"
	"net/http"
	"testing"
)

func Test_Post(t *testing.T) {
	responseBody := bytes.NewBuffer([]byte(`{"key":"value"}`))

	resp, err := http.Post("http://0.0.0.0:8080/api/v1/images:upload", "application/json", responseBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, string(body), "upload image")
}
