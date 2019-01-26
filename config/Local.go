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
func (c LocalConfig) Read() ServerConfig {
	config := ServerConfig{}

	// Get server configuration
	p := c.Location + "server_config.yml"
	data, err := ioutil.ReadFile(p)
	errors.CheckError(err)

	err = yaml.Unmarshal(data, &config)
	errors.CheckError(err)

	return config
}
