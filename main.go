package main

import (
	"github.com/adamb/scriptdeliver/channels/sshchannel"
	"github.com/adamb/scriptdeliver/config"
)

func main() {
	c := config.GetConfig()
	tag := c.GetTagConfig("test")
	s := sshchannel.Open(tag.SshConfig)
	s.RunScript(tag.GetInitScript(), "init.sh")
}
