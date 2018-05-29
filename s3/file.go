package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	. "github.com/ctripcorp/nephele/storage"
)

type file struct {
	sess       *session.Session
	bucketname string
	key        string
	blob       []byte
	err        error
}

func (f *file) Key() string {
	return f.key
}

func (f *file) Exist() (bool, string, error) {
	svc := s3.New(f.sess)
	params := &s3.HeadObjectInput{
		Bucket: aws.String(f.bucketname),
		Key:    aws.String(f.key),
	}
	req, _ := svc.HeadObjectRequest(params)
	if req.Send() != nil {
		if strings.HasPrefix(req.Send().Error(), "NotFound") {
			return false, req.RequestID, nil
		}
		return false, req.RequestID, req.Send()
	}
	return true, req.RequestID, nil
}

func (f *file) Meta() (Fetcher, error) {
	return &fetcher{nil}, errors.New("Meta not supported")
}

func (f *file) GetFile(filename string) ([]byte, string, error) {
	svc := s3.New(f.sess)
	params := &s3.GetObjectInput{
		Bucket: aws.String(f.bucketname),
		Key:    aws.String(f.key),
	}
	req, resp := svc.GetObjectRequest(params)
	if req.Send() != nil {
		return nil, req.RequestID, req.Send()
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, req.RequestID, err
	}
	return b, req.RequestID, nil
}

func (f *file) PutFile(filename string, blob []byte, kvs ...KV) (string, error) {
	svc := s3.New(f.sess)
	params := &s3.PutObjectInput{
		Bucket: aws.String(f.bucketname),
		Key:    aws.String(filename),
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

func (f *file) Append(blob []byte, index int64, kvs ...KV) (int64, string, error) {
	if index == 0 {
		requestID, err := f.PutFile(f.key, blob, kvs...)
		if err != nil {
			return 0, requestID, err
		}
		return int64(len(blob)), requestID, nil
	}
	b, requestID, err := f.GetFile(f.key)
	if err != nil {
		return index, requestID, err
	}
	lenb := int64(len(b))
	if lenb == index {
		mix := append(b, blob...)
		requestID, err := f.PutFile(f.key, mix, kvs...)
		if err != nil {
			return lenb, requestID, err
		}
		return int64(len(mix)), requestID, nil
	}
	return lenb, requestID, err
}

func (f *file) Delete() (string, error) {
	svc := s3.New(f.sess)
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(f.bucketname),
		Key:    aws.String(f.key),
	}
	req, _ := svc.DeleteObjectRequest(params)
	if req.Send() != nil {
		return req.RequestID, req.Send()
	}
	return req.RequestID, nil
}

func (f *file) Bytes() ([]byte, string, error) {
	b, requestID, err := f.GetFile(f.key)
	if err != nil {
		return nil, requestID, err
	}
	return b, requestID, nil
}

func (f *file) SetMeta(kvs ...KV) error {
	return errors.New("SetMeta not supported")
}
