package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	sc := GetConfig()
	tag := sc.GetTagConfig("test")
	name, _ := tag.GetStateScript("init")
	fmt.Printf("%v\n", name)
}

func S3Test(t *testing.T) {
	c := S3Config{}
	sc := c.Read()
	tag := sc.GetTagConfig("test")
	name, _ := tag.GetStateScript("init")
	fmt.Printf("SCript name: %v\n", name)
}
