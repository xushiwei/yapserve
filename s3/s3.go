package s3

import (
	"context"
	"net/http"

	"github.com/qiniu/x/http/fsx"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
)

const (
	Scheme = s3blob.Scheme
)

func init() {
	fsx.Register(Scheme, Open)
}

func Open(ctx context.Context, url string) (_ http.FileSystem, _ fsx.Closer, err error) {
	bucket, err := blob.OpenBucket(ctx, url)
	if err != nil {
		return
	}
	return http.FS(bucket), bucket.Close, nil
}
