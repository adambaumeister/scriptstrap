package config

import (
	"bytes"
	"fmt"
	"github.com/adamb/scriptdeliver/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/yaml.v2"
)

type S3Config struct {
	Location string

	AwsSvc *s3.S3
}

func (c *S3Config) Read() {
	config := ServerConfig{}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	_ = s3.ListObjectsInput{
		Bucket: aws.String("testscriptstrap"),
	}
	svc := s3.New(sess)

	configObject, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("testscriptstrap"),
		Key:    aws.String("server_config.yml"),
	})
	errors.CheckAwsError(err)
	c.AwsSvc = svc

	buf := new(bytes.Buffer)
	buf.ReadFrom(configObject.Body)
	err = yaml.Unmarshal(buf.Bytes(), &config)
	errors.CheckError(err)

	for tn, _ := range config.Tags {
		fmt.Printf("%v\n", tn)
	}

}

func (c *S3Config) GetStateScript(s string) (string, []byte) {
	svc := c.AwsSvc
	configObject, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("testscriptstrap"),
		Key:    aws.String(s),
	})
	errors.CheckAwsError(err)
	c.AwsSvc = svc

	buf := new(bytes.Buffer)
	buf.ReadFrom(configObject.Body)
	return s, buf.Bytes()
}
