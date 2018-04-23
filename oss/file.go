package main

import (
	"bytes"
	. "github.com/phyxdown/nephele-storage"
	"io/ioutil"

	"github.com/phyxdown/aliyun-oss-go-sdk/oss"
)

type file struct {
	bucket   *oss.Bucket
	filename string
	blob     []byte
}

func (f *file) Exist() (bool, string, error) {
	return f.bucket.IsObjectExist(f.filename)
}

func (f *file) Meta() (Fetcher, error) {
	h, err := f.bucket.GetObjectDetailedMeta(f.filename)
	if err != nil {
		return nil, err
	}
	return &fetcher{h}, nil
}

func (f *file) Append(blob []byte, index int64, KVs ...KV) (int64, string, error) {
	options := make([]oss.Option, 0)
	for _, KV := range KVs {
		options = append(options, oss.Meta(KV[0], KV[1]))
	}
	return f.bucket.AppendObject(f.filename, bytes.NewReader(blob), index, options...)
}

func (f *file) Delete() (string, error) {
	return f.bucket.DeleteObject(f.filename)
}

func (f *file) Bytes() ([]byte, string, error) {
	r, rid, err := f.bucket.GetObject(f.filename)
	if err != nil {
		return nil, "", err
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, "", err
	}
	return b, rid, nil
}

func (f *file) SetMeta(KVs ...KV) error {
	options := make([]oss.Option, 0)
	for _, KV := range KVs {
		options = append(options, oss.Meta(KV[0], KV[1]))
	}
	return f.bucket.SetObjectMeta(f.filename, options...)
}
