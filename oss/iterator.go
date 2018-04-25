package main

import (
	"time"

	. "github.com/phyxdown/nephele-storage"

	"github.com/phyxdown/aliyun-oss-go-sdk/oss"
)

type iterator struct {
	bucket          *oss.Bucket
	prefix, lastKey string
	files           chan File
}

func (iter *iterator) syncing() {
	for {
		r, err := iter.bucket.ListObjects(oss.Marker(iter.lastKey), oss.Prefix(iter.prefix))
		if err != nil {
			time.Sleep(time.Minute)
			continue
		}
		if len(r.Objects) == 0 {
			iter.files <- nil
			time.Sleep(time.Minute)
			continue
		}
		for _, object := range r.Objects {
			iter.files <- &file{
				bucket: iter.bucket,
				key:    object.Key,
			}
			iter.lastKey = object.Key
		}
	}
}

func (iter *iterator) Next() File {
	return <-iter.files
}

func (iter *iterator) LastKey() string {
	return iter.lastKey
}
