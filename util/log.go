package util

import (
	col "github.com/fatih/color"

	"github.com/binchencoder/ease-gateway/util/glog"
)

func init() {
	col.NoColor = false
}

// janus-rest日志
var (
	// janus请求和响应详情日志. 支持格式包括:
	// 1. RequestRestFormat
	// 2. ResponseRestFormat
	RestLogger = glog.Context(nil, glog.FileName{Name: "janus-rest"})

	// request rest log format.
	// [request]traceId,serviceName,httpMethod,path,{reqInfo}
	// external {reqInfo}: source,client,appVersion,u:uid,c:cid,a:aid,d:did,remoteAddr,url
	// ldap {reqInfo}: remoteAddr,url
	RequestRestFormat = "[request]%s,%s,%s,%s,%s"

	// response rest log format.
	// [response]traceId,code,msg
	ResponseRestFormat = "[response]%s,%d,%s"
)

// janus-config日志
var (
	// janus配置操作相关日志.
	ConfigLogger = glog.Context(nil, glog.FileName{Name: "janus-config"})
)

// janus-stat日志
var (
	// janus访问统计日志. 支持格式包括:
	// 1. StatFormat
	StatLogger = glog.Context(nil, glog.FileName{Name: "janus-stat"})

	// stat log format.
	// traceId,serviceName,httpMethod,path,client,ok:isSuccess,code,ms:response-ms
	StatFormat = "%s,%s,%s,%s,%s,ok:%s,%d,ms:%g"
)

// janus-error日志
var (
	// janus错误日志.
	ErrorLogger = glog.Context(nil, glog.FileName{Name: "janus-error"})

	// error log format.
	ErrorFormat = "%s,%s"
)

// janus-limit日志
var (
	// janus限流日志.
	LimitLogger = glog.Context(nil, glog.FileName{Name: "janus-limit"})

	// traceId[limit type]service,httpMethod,path,limitInfo.
	LimitFormat = "%s[%s]%s,%s,%s,%s"
)

// janus-default日志
var (
	// janus默认日志文件, 排除上面的剩下的日志.
	DefaultLogger = glog.Context(nil, glog.FileName{Name: "janus-default"})

	// default log format.
	DefaultFormat = "%s,%s"
)

// Flush be used to flush logs immediately.
func Flush() {
	glog.Flush()
}
