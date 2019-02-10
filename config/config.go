package config

import (
	"fmt"
	"github.com/adamb/go_osegp/bgp/errors"
	"github.com/adamb/scriptdeliver/channels/sshchannel"
	"io/ioutil"
	"os"
	"strings"
)

type Config interface {
	Read() ServerConfig
}

type ServerConfig struct {
	Tags map[string]*Tag
}

type Tag struct {
	SshConfig *sshchannel.Opts

	initScript   string
	stateScripts map[string]Script
}

func (t *Tag) GetInitScript() []byte {
	f, err := ioutil.ReadFile(t.initScript)
	errors.CheckError(err)

	return f
}

type Script struct {
	Data     []byte
	Filename string
}

func (t *Tag) GetStateScript(s string) (string, []byte) {
	return t.stateScripts[s].Filename, t.stateScripts[s].Data
}

func (sc *ServerConfig) GetTagConfig(t string) *Tag {
	if tag, ok := sc.Tags[t]; ok {
		return tag
	} else {
		fmt.Printf("Requested tag not found: %v\n", t)
		return nil
	}
}

func GetFromEnv(key string) (string, bool) {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key), true
	} else {
		return "", false
	}
}

func GetConfig() ServerConfig {
	s, result := GetFromEnv("SCRIPTSTRAP_CONFIG")

	if result == false {
		errors.RaiseError("No SCRIPTSTRAP_CONFIG specified.")
	}

	splt := strings.Split(s, ":")
	t := splt[0]
	loc := strings.Join(splt[1:], ":")
	fmt.Printf("%v\n", t)

	var c Config
	var sc ServerConfig
	if t == "local" {
		c = LocalConfig{
			Location: loc,
		}
	} else if t == "s3" {
		c = S3Config{
			Location: loc,
		}
	}
	sc = c.Read()

	return sc
}
