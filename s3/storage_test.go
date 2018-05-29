package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func TestStorage(t *testing.T) {
	println("start test stroage")
	region := ""
	bucketname := ""
	accessKeyId := ""
	accessKeySecret := ""

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyId, accessKeySecret, ""),
		Region:      &region,
	}))

	var stor storage
	stor.bucketname = bucketname
	stor.sess = sess
	str, err := stor.StoreFile("just for test storage", []byte("test storage"))
	if err != nil {
		t.Error("err storage")
	}
	println("requestID is ", str)
}
