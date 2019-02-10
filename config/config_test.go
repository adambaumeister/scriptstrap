package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	sc := GetConfig()
	name, _ := sc.Tags["test"].GetStateScript("init")
	fmt.Printf("%v\n", name)
}

func TestS3(t *testing.T) {
	c := S3Config{}
	sc := c.Read()
	fmt.Printf("%v\n", sc.Tags["test"].stateScripts)
}
