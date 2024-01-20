package fileblob

import (
	"context"
	"net/http"
	"strings"

	"github.com/qiniu/x/http/fsx"
	"gocloud.dev/blob/fileblob"
)

const (
	Scheme = "fileblob"
)

func init() {
	fsx.Register(Scheme, Open)
}

func Open(ctx context.Context, url string) (_ http.FileSystem, _ fsx.Closer, err error) {
	url = strings.TrimPrefix(url, Scheme+":")
	bucket, err := fileblob.OpenBucket(url, nil)
	if err != nil {
		return
	}
	return http.FS(bucket), bucket.Close, nil
}
