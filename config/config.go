package config

import (
	"fmt"
	"github.com/adamb/go_osegp/bgp/errors"
	"github.com/adamb/scriptdeliver/channels/sshchannel"
	"os"
	"strings"
)

type Config interface {
	Read() ServerConfig
}

type ServerConfig struct {
	Tags map[string]Tag
}

type Tag struct {
	SshConfig sshchannel.Opts
}

func GetFromEnv(key string) (string, bool) {
	if len(os.Getenv(key)) > 0 {
		return os.Getenv(key), true
	} else {
		return "", false
	}
}

func GetConfig() Config {
	s, result := GetFromEnv("SCRIPTSTRAP_CONFIG")

	if result == false {
		errors.RaiseError("No SCRIPTSTRAP_CONFIG specified.")
	}

	splt := strings.Split(s, ":")
	t := splt[0]
	loc := strings.Join(splt[1:], ":")
	fmt.Printf("%v\n", t)

	var c Config
	if t == "local" {
		c = LocalConfig{
			Location: loc,
		}
		sc := c.Read()
		fmt.Printf("%v\n", sc.Tags["test"].SshConfig.SshKeyFile)
	}

	return c
}
