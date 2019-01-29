package main

import (
	"fmt"
	"github.com/adamb/scriptdeliver/channels/sshchannel"
	"github.com/adamb/scriptdeliver/config"
	"github.com/adamb/scriptdeliver/events"
	"github.com/adamb/scriptdeliver/events/api"
)

func main() {
	c := config.GetConfig()

	a := api.API{
		EventsOut: make(chan events.Event),
	}
	go a.StartApi()
	fmt.Printf("Started all listeners.\n")
	for {
		select {
		case e := <-a.EventsOut:
			t := c.GetTagConfig(e.GetTag())
			if t != nil {
				name, script := t.GetStateScript(e.GetState())

				channel := sshchannel.Open(e.GetHost(), t.SshConfig)
				channel.RunScript(script, name)
			}
		}
	}
}
