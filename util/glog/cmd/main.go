package main

import (
	"flag"

	"github.com/binchencoder/ease-gateway/util/glog"
)

func main() {
	flag.Parse()
	p := glog.Context(nil, nil)
	defer p.Flush()
	p.Warning("aaaa")

	var a interface{}
	// If -roll_type=date filename will be used
	a = glog.FileName{Name: "janus-common"}
	f := glog.Context(nil, a)
	f.Info("bbbbbbbbbb")
	defer f.Flush()

	var b interface{}
	// If -roll_type=date filename will be used
	b = glog.FileName{Name: "janus-error"}
	z := glog.Context(nil, b)
	z.Error("zzzzzzzzzzz")
	defer z.Flush()
}
