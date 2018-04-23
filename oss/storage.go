package main

import (
	. "github.com/phyxdown/nephele-storage"

	"github.com/phyxdown/aliyun-oss-go-sdk/oss"
)

type storage struct {
	bucket *oss.Bucket
}

func (s *storage) File(filename string) File {
	return &file{
		bucket:   s.bucket,
		filename: filename,
	}
}

func (s *storage) Iterator(prefix string) Iterator {
	return nil
}

func (s *storage) StoreFile(filename string, blob []byte, KVs ...KV) (string, error) {
	return "", nil
}
