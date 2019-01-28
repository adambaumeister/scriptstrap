package main

import (
	"github.com/adamb/scriptdeliver/channels/sshchannel"
	"github.com/adamb/scriptdeliver/config"
	"github.com/adamb/scriptdeliver/events"
)

func main() {
	c := config.GetConfig()

	e := events.RouteEvent("192.168.1.18:22,test")
	tag := c.GetTagConfig(e.GetTag())
	s := sshchannel.Open(tag.SshConfig)
	s.RunScript(tag.GetInitScript(), "init.sh")
}
