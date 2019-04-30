# Glog

## Glog Description

- Glog is extends base on github.com/golang/glog and jingoal.com/letsgo/log,
 modify the part of the source code.

## Extend Features

- -roll\_type flag add.if -roll\_type=size, log file will division deps file
 size;
if -roll\_type=date,log file will division deps date. default is size. size is
 glog support.
- Can set log file name by glog's api. which can help program record specific
log into specific file.
- Traceid and log file name not required.
- Flags which be used in github.com/golang/glog be added prefix g_.

## Matters Needing Attention

- If -roll\_type=size,log file name will contains pid、time and so on.
- Log file naming rules make minor readjustments, could refer to glog_file.go.
- Traceid features is still support,could refer to jingoal.com/letsgo/log/log.go.
- If log file name not set,default generate INFO、WARN and so on glog supported
file,and high level log will also write into low level log file, high and
 low level log file all generated, this rule could refer to glog
- If log file name is set, log msg will exactly write into one log file naming
by the name set

## Instructions

- glog.Flush() and flag.Parse() will must be called. But if it is a long-term
running of the program, flush will timed automatic execution.

```
import (
    "flag"
    "github.com/binchencoder/ease-gateway/util/glog"
)
func main() {
    flag.Parse()
    p := glog.Context(nil,nil)
    defer p.Flush()
    p.Warning("aaaa")
}
 ```

- Log file will like Build main.go and rungo.tanpeng.TANPENG\_tanpeng.log.INFO
.20161124-171613.33776(roll\_type=size) or Build main.go and rungo.tanpeng.
TANPENG\_tanpeng.log.WARNING.20161124-171613.33776(roll\_type=size) or Build
main.go and rungo.tanpeng.TANPENG_tanpeng.log.INFO.20161124(roll\_type=date);
rule like program.host.username.log.INFO(log level).time.pid or program.host.
username.log.INFO(log level).date

```
import (
    "flag"
    "github.com/binchencoder/ease-gateway/util/glog"
)
func main() {
    flag.Parse()
    var a interface{}
    a = glog.FileName{"mobile-common"}
    p := glog.Context(nil,a)
    defer p.Flush()
    p.Warning("aaaa")
}
```

- gl.FileName is a struct extend in glog.go,it decide log file name;
log file will like Build main.go and rungo.tanpeng.TANPENG\_tanpeng.
log.mobile-common.20161124-172846.32712(roll\_type=size) or Build
main.go and rungo.tanpeng.TANPENG\_tanpeng.log.mobile-common.20161124
(roll\_type=date); rule like program.host.username.log.filename.time.pid
or program.host.username.log.filename.date

- How to add traceid

```
ctx := trace.NewTraceId(context.Background())
p := log.Context(ctx,nil)
```

- Don't forget add import golang.org/x/net/context and jingoal.com/letsgo/trace