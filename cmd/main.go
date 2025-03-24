package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/thanos-io/objstore/providers/oss"
	// "github.com/thanos-io/objstore/errutil"
)

func main() {
	
	config := oss.Config{
		Region:          "cn-beijing",
		Endpoint:        "https://oss-cn-beijing.aliyuncs.com",
		Bucket:          "zilliz-test-yellow-tmp",
	}

	bkt, err := oss.NewBucketWithConfig(log.NewNopLogger(), config, "test", nil)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	content := []byte("hello from RRSA + Go SDK1")

	err = bkt.Upload(context.Background(), "my/test1.txt", bytes.NewReader(content))
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("upload success")
}