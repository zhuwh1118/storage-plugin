package main

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/phyxdown/nephele-storage"

	"github.com/phyxdown/aliyun-oss-go-sdk/oss"
)

func main() {}

func Create(config map[string]string) Storage {
	endpoint := config["endpoint"]
	bucketname := config["bucketname"]
	accessKeyId := config["accessKeyId"]
	accessKeySecret := config["accessKeySecret"]

	var proxy oss.ClientOption = func(client *oss.Client) {}
	if os.Getenv("http_proxy") != "" {
		proxy = oss.Proxy(os.Getenv("http_proxy"))
	}

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret, proxy)
	if err != nil {
		return nil
	}

	bucket, err := client.Bucket(bucketname)
	if err != nil {
		return nil
	}

	return &storage{bucket}
}
