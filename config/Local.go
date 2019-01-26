package config

import (
	"github.com/adamb/scriptdeliver/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LocalConfig struct {
	Location string
}

// REad a yaml configuration
func (c LocalConfig) Read() Config {
	// Setup defaults
	p := c.Location
	data, err := ioutil.ReadFile(p)
	errors.CheckError(err)

	err = yaml.Unmarshal(data, &c)
	errors.CheckError(err)

	return c
}
