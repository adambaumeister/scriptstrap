package config

import (
	"bytes"
	"fmt"
	"github.com/adamb/scriptdeliver/errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/yaml.v2"
	"strings"
)

type S3Config struct {
	Location string

	AwsSvc *s3.S3
}

func (c S3Config) Read() ServerConfig {
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

	c.ScriptsFromTags(&config)
	return config
}

func (c *S3Config) ScriptsFromTags(sc *ServerConfig) {
	for tn, tag := range sc.Tags {
		r := make(map[string]Script)
		scripts := c.ListObjects(tn)
		for _, sn := range scripts {
			s := Script{}
			s.Filename = NameFromKey(sn)
			s.Data = c.GetFileBytes(sn)
			stateName := strings.Split(NameFromKey(sn), ".")[0]
			r[stateName] = s
		}
		tag.stateScripts = r
	}
}

func (c *S3Config) GetFileBytes(k string) []byte {
	svc := c.AwsSvc
	configObject, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("testscriptstrap"),
		Key:    aws.String(k),
	})
	errors.CheckAwsError(err)
	c.AwsSvc = svc

	buf := new(bytes.Buffer)
	buf.ReadFrom(configObject.Body)
	return buf.Bytes()
}

func (c *S3Config) ListObjects(k string) []string {
	var r []string

	svc := c.AwsSvc
	objs, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String("testscriptstrap"),
		Prefix: aws.String(fmt.Sprintf("%v/%v/", "tags", k)),
	})
	errors.CheckAwsError(err)
	c.AwsSvc = svc

	for _, obj := range objs.Contents {
		r = append(r, *obj.Key)
	}
	return r
}

func NameFromKey(s string) string {
	l := strings.Split(s, "/")
	length := len(l)
	str := l[length-1]
	return str
}
