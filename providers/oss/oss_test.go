package oss

import (
	"bytes"
	"context"
	"testing"

	"github.com/efficientgo/core/testutil"
	"github.com/go-kit/log"
	"github.com/thanos-io/objstore/errutil"
)

func TestNewBucketWithErrorRoundTripper(t *testing.T) {
	config := Config{
		Endpoint:        "http://test.com/",
		AccessKeyID:     "123",
		AccessKeySecret: "123",
		Bucket:          "test",
	}

	bkt, err := NewBucketWithConfig(log.NewNopLogger(), config, "test", errutil.WrapWithErrRoundtripper)
	// We expect an error from the RoundTripper
	testutil.Ok(t, err)
	_, err = bkt.Get(context.Background(), "test")
	testutil.NotOk(t, err)
	testutil.Assert(t, errutil.IsMockedError(err), "Expected RoundTripper error, got: %v", err)
}


func TestNewBucketWithOidcCredential(t *testing.T) {
	config := Config{
		Region:          "cn-beijing",
		Endpoint:        "https://oss-cn-beijing.aliyuncs.com",
		Bucket:          "zilliz-test-yellow-tmp",
	}

	bkt, err := NewBucketWithConfig(log.NewNopLogger(), config, "test", nil)
	testutil.Ok(t, err)

	content := []byte("hello from OIDC credential")

	err = bkt.Upload(context.Background(), "test", bytes.NewReader(content))
	testutil.Ok(t, err)
}

