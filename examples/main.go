package main

import (
	"context"
	"log/slog"
	"strings"

	"github.com/go-kit/log"
	"github.com/thanos-io/objstore/providers/cos"
)

func main() {
	logger := log.NewNopLogger()
	Bucket, err := cos.NewBucketWithConfig(
		logger,
		cos.Config{
			// Bucket:    "test-1253846545",
			// Region:    "ap-guangzhou",
			Bucket:    "liang-1318287928",
			Region:    "ap-shanghai",
			// Endpoint:  "https://liang-1318287928.cos.ap-shanghai.myqcloud.com",
		},
		"test",
		nil,
	)
	if err != nil {
		panic(err)
	}


	err = Bucket.Upload(context.Background(), "testr2.txt", strings.NewReader("test1"))
	if err != nil {
		panic(err)
	}

	slog.Info("uploaded file", "file", "test.txt")
}
