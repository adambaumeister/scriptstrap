package api

import (
	"encoding/json"
	"github.com/adamb/scriptdeliver/errors"
	"github.com/adamb/scriptdeliver/events"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	EventsOut chan events.Event
}

func (a *API) StartApi() {
	http.HandleFunc("/", a.NewEvent)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (a *API) NewEvent(w http.ResponseWriter, r *http.Request) {
	e := EventJson{}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &e)
	errors.CheckError(err)
	ae := ApiEvent{}
	ae.Host = e.Host
	ae.Tag = e.Tags[0]

	response := JsonResponse{
		Status: "OK",
	}
	j, err := json.Marshal(response)
	errors.CheckError(err)
	w.Write(j)

	a.EventsOut <- ae
}
