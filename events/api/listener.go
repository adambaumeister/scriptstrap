package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
}

func StartApi() {
	a := API{}
	http.HandleFunc("/", a.NewEvent)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func (a *API) NewEvent(w http.ResponseWriter, r *http.Request) {
	e := EventJson{}

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &e)
	fmt.Printf("Host: %v\n", e.Host)
}
