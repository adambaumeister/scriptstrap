package api

import (
	"bytes"
	"encoding/json"
	"github.com/adamb/scriptdeliver/errors"
	"net/http"
	"testing"
)

func TestEventApi(t *testing.T) {
	a := API{}
	go a.StartApi()

	e := EventJson{
		Host: "192.168.1.18",
		Tags: []string{"test"},
	}
	j, _ := json.Marshal(e)

	_, err := http.Post("http://127.0.0.1:8080/", "application/json", bytes.NewBuffer(j))
	errors.CheckError(err)
}
