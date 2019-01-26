package sshchannel

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
)

type SshChannel struct {
}

func Open() SshChannel {

	s := SshChannel{}

	key, err := ioutil.ReadFile("C:\\Users\\adam\\.ssh\\testrsa")
	if err != nil {
		log.Fatalf("Unable to read key. %v", err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Unable to read key.")
	}

	c := ssh.ClientConfig{
		User: "adam",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: s.checkHostKey,
	}
	s.openSession("buildserver.home:22", c)
	return s
}

func (s *SshChannel) openSession(h string, c ssh.ClientConfig) {
	client, err := ssh.Dial("tcp", h, &c)
	if err != nil {
		log.Fatalf("Unable to connect to host :%v", err)
	}
	client.Close()
}

func (s *SshChannel) checkHostKey(h string, r net.Addr, key ssh.PublicKey) error {
	return nil
}
