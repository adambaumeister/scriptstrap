package sshchannel

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/adamb/scriptdeliver/errors"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"sync"
)

type SshChannel struct {
	scriptDir string
	scriptOwn string

	stdout   *bytes.Buffer
	client   *ssh.Client
	sessions []*ssh.Session
}

func Open(config *Opts) SshChannel {

	s := SshChannel{
		scriptDir: "/tmp/",
		scriptOwn: "0775",
		stdout:    &bytes.Buffer{},
	}

	key := config.SshKey
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Unable to read key.")
	}

	c := ssh.ClientConfig{
		User: config.SshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: s.checkHostKey,
	}
	s.openConnection("192.168.1.18:22", c)
	return s
}

//Runscript requires
//	script bytes
//	script name
func (s *SshChannel) RunScript(f []byte, n string) {
	s.WriteFile(f, n, s.scriptDir, s.scriptOwn)

	session, err := s.client.NewSession()
	session.Stdout = s.stdout
	s.sessions = append(s.sessions, session)

	err = session.Run(s.scriptDir + n)
	errors.CheckError(err)
	fmt.Printf("Output: %v\n", string(s.stdout.Bytes()))
}

func (s *SshChannel) openConnection(h string, c ssh.ClientConfig) {
	client, err := ssh.Dial("tcp", h, &c)
	if err != nil {
		log.Fatalf("Unable to connect to host :%v", err)
	}
	s.client = client
}

func (s *SshChannel) checkHostKey(h string, r net.Addr, key ssh.PublicKey) error {
	return nil
}

// Write a file to the remote endpoint using SCP and a pipe
// Requires
// 	[]byte object of file data
//	string filename
//	string containing end directory
// 	chown string (format: 0775)
func (s *SshChannel) WriteFile(f []byte, fn string, d string, p string) {
	// Start the session
	session, err := s.client.NewSession()
	session.Stdout = s.stdout
	s.sessions = append(s.sessions, session)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		hostIn, _ := session.StdinPipe()
		defer hostIn.Close()
		fmt.Fprintf(hostIn, "C%v %d %s\n", p, len(f), fn)
		binary.Write(hostIn, binary.BigEndian, f)
		fmt.Fprint(hostIn, "\x00")
		wg.Done()
	}()
	err = session.Run("/usr/bin/scp -t " + d)
	errors.CheckError(err)
	wg.Wait()
}
