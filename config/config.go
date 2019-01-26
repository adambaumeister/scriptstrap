package config

import (
	"fmt"
	"github.com/adamb/go_osegp/bgp/errors"
	"os"
	"strings"
)

type Config interface {
	Read() Config
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
	loc := splt[1]
	fmt.Printf("%v\n", t)

	var c Config
	if t == "local" {
		c = LocalConfig{
			Location: loc,
		}
	}

	return c
}
