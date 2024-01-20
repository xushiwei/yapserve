package main

import (
	"fmt"
	"github.com/goplus/yap"
	"context"
	"flag"
	_ "github.com/qiniu/x/http/fsx/cached"
	_ "github.com/qiniu/x/http/fsx/filter"
	_ "github.com/qiniu/x/http/fsx/local"
	_ "github.com/xushiwei/kodofs"
	_ "github.com/xushiwei/yapserve/env"
	_ "github.com/xushiwei/yapserve/fileblob"
	_ "github.com/xushiwei/yapserve/s3"
)

type yapserve struct {
	yap.App
}
//line yapserve_yap.gox:14
func (this *yapserve) MainEntry() {
//line yapserve_yap.gox:14:1
	addr := flag.String("host", ":8888", "server host")
//line yapserve_yap.gox:15:1
	flag.Parse()
//line yapserve_yap.gox:16:1
	if flag.NArg() < 1 {
//line yapserve_yap.gox:17:1
		fmt.Println("yapserve [flags] scheme")
//line yapserve_yap.gox:18:1
		flag.PrintDefaults()
//line yapserve_yap.gox:19:1
		return
	}
//line yapserve_yap.gox:21:1
	scheme := flag.Arg(0)
//line yapserve_yap.gox:23:1
	this.Static__1("/", context.Background(), scheme)
//line yapserve_yap.gox:25:1
	this.Run(*addr)
}
func main() {
	yap.Gopt_App_Main(new(yapserve))
}
