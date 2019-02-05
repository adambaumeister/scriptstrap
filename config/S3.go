package config

import (
	"bytes"
	"github.com/adamb/scriptdeliver/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/yaml.v2"
)

type S3Config struct {
	Location string
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

	buf := new(bytes.Buffer)
	buf.ReadFrom(configObject.Body)
	err = yaml.Unmarshal(buf.Bytes(), &config)
	errors.CheckError(err)

}
