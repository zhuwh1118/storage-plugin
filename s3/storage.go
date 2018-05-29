package main

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	. "github.com/ctripcorp/nephele/storage"
)

type storage struct {
	sess       *session.Session
	bucketname string
}

func (s *storage) File(key string) File {
	return &file{
		sess:       s.sess,
		bucketname: s.bucketname,
		key:        key,
	}
}

func (s *storage) Iterator(prefix string, lastKey string) Iterator {
	return nil
}

func (s *storage) StoreFile(key string, blob []byte, kvs ...KV) (string, error) {
	svc := s3.New(s.sess)
	params := &s3.PutObjectInput{
		Bucket: aws.String(s.bucketname),
		Key:    aws.String(key),
		Body:   bytes.NewReader(blob),
	}
	if len(kvs) > 0 {
		metadata := make(map[string]*string)
		for i := range kvs {
			metadata[kvs[i][0]] = &kvs[i][1]
		}
		params.SetMetadata(metadata)
	}
	req, _ := svc.PutObjectRequest(params)
	if req.Send() != nil {
		return req.RequestID, req.Send()
	}
	return req.RequestID, nil
}
