package api

import (
	"bytes"
	"encoding/json"
	"github.com/adamb/scriptdeliver/errors"
	"net/http"
	"testing"
)

func TestEventApi(t *testing.T) {
	go StartApi()

	e := EventJson{
		Host: "192.168.1.18",
		tags: []string{"test"},
	}
	j, _ := json.Marshal(e)

	_, err := http.Post("http://127.0.0.1:8000/", "application/json", bytes.NewBuffer(j))
	errors.CheckError(err)
}
