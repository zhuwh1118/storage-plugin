package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	. "github.com/ctripcorp/nephele/storage"
)

func main() {}

func New(config map[string]string) Storage {
	region := config["region"]
	bucketname := config["bucketname"]
	accessKeyId := config["accessKeyId"]
	accessKeySecret := config["accessKeySecret"]

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyId, accessKeySecret, ""),
		Region:      &region,
	}))

	if sess == nil {
		return nil
	}
	return &storage{sess, bucketname}
}
