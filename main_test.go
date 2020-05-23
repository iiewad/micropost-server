package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

type mainRes struct {
	Message string `json:"message"`
}

func TestMain(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		log.Fatal(err)
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var got, want mainRes
	want.Message = "pong"
	json.Unmarshal(resBody, &got)
	if got != want {
		t.Errorf("main() = %q, want %q", got, want)
	}

}
