package config

import (
	"github.com/adamb/scriptdeliver/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
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

	for tn, tag := range config.Tags {
		tag.stateScripts = scriptsInDir(filepath.Join(c.Location, "tags", tn))
		tag.initScript = filepath.Join(c.Location, "tags", tn, "init.sh")
		if tag.SshConfig != nil {
			tag.SshConfig.SshKey, err = ioutil.ReadFile(tag.SshConfig.SshKeyFile)
			errors.CheckError(err)
		}
	}

	return config
}

func scriptsInDir(dir string) map[string]string {
	r := make(map[string]string)
	files, err := ioutil.ReadDir(dir)
	errors.CheckError(err)
	for _, file := range files {
		r[strings.Split(file.Name(), ".")[0]] = filepath.Join(dir, file.Name())
	}
	return r
}
