package main

import (
	"fmt"
	"flag"
	"net/url"
	"github.com/qiniu/x/token/protected"
	_ "github.com/xushiwei/yapserve/env"
	"github.com/qiniu/x/errors"
)
//line kodourl/kodourl.gop:9
func main() {
//line kodourl/kodourl.gop:9:1
	help := flag.Bool("h", false, "show this help information")
//line kodourl/kodourl.gop:10:1
	flag.Parse()
//line kodourl/kodourl.gop:12:1
	args := flag.Args()
//line kodourl/kodourl.gop:13:1
	if *help || len(args) != 3 {
//line kodourl/kodourl.gop:14:1
		fmt.Println("kodourl [flags] <bucket> <ak> <sk>\n")
//line kodourl/kodourl.gop:15:1
		flag.PrintDefaults()
//line kodourl/kodourl.gop:16:1
		return
	}
//line kodourl/kodourl.gop:19:1
	params := make(url.Values)
//line kodourl/kodourl.gop:20:1
	params.Set("ak", args[1])
//line kodourl/kodourl.gop:21:1
	params.Set("sk", args[2])
//line kodourl/kodourl.gop:22:1
	query := func() (_gop_ret string) {
//line kodourl/kodourl.gop:22:1
		var _gop_err error
//line kodourl/kodourl.gop:22:1
		_gop_ret, _gop_err = protected.Encode(params)
//line kodourl/kodourl.gop:22:1
		if _gop_err != nil {
//line kodourl/kodourl.gop:22:1
			_gop_err = errors.NewFrame(_gop_err, "protected.encode(params)", "kodourl/kodourl.gop", 22, "main.main")
//line kodourl/kodourl.gop:22:1
			panic(_gop_err)
		}
//line kodourl/kodourl.gop:22:1
		return
	}()
//line kodourl/kodourl.gop:24:1
	fmt.Println("kodo:" + args[0] + "?" + query)
}
