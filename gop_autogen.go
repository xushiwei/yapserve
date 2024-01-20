package main

import (
	"fmt"
	"github.com/goplus/yap"
	"context"
	"flag"
	_ "github.com/qiniu/x/http/fsx/cached"
	_ "github.com/qiniu/x/http/fsx/filter"
	_ "github.com/qiniu/x/http/fsx/local"
)

type yapserve struct {
	yap.App
}
//line yapserve_yap.gox:10
func (this *yapserve) MainEntry() {
//line yapserve_yap.gox:10:1
	addr := flag.String("host", ":8888", "server host")
//line yapserve_yap.gox:11:1
	flag.Parse()
//line yapserve_yap.gox:12:1
	if flag.NArg() < 1 {
//line yapserve_yap.gox:13:1
		fmt.Println("yapserve [flags] scheme")
//line yapserve_yap.gox:14:1
		flag.PrintDefaults()
//line yapserve_yap.gox:15:1
		return
	}
//line yapserve_yap.gox:17:1
	scheme := flag.Arg(0)
//line yapserve_yap.gox:19:1
	this.Static__1("/", context.Background(), scheme)
//line yapserve_yap.gox:21:1
	this.Run(*addr)
}
func main() {
	yap.Gopt_App_Main(new(yapserve))
}
