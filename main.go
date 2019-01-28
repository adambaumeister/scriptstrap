package main

import (
	"fmt"
	"github.com/adamb/scriptdeliver/events"
	"github.com/adamb/scriptdeliver/events/api"
)

func main() {
	a := api.API{
		EventsOut: make(chan events.Event),
	}
	go a.StartApi()
	fmt.Printf("Started all listeners.\n")
	for {
		select {
		case e := <-a.EventsOut:
			fmt.Printf("Host: %v\n", e.GetHost())
		}
	}
}
