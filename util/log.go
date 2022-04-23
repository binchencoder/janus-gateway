package util

import (
	col "github.com/fatih/color"

	"github.com/binchencoder/janus-gateway/util/glog"
)

func init() {
	col.NoColor = false
}

// gateway-rest日志
var (
	// janus-gateway 请求和响应详情日志. 支持格式包括:
	// 1. RequestRestFormat
	// 2. ResponseRestFormat
	RestLogger = glog.Context(nil, glog.FileName{Name: "gateway-rest"})

	// request rest log format.
	// [request]traceId,serviceName,httpMethod,path,{reqInfo}
	// external {reqInfo}: source,client,appVersion,u:uid,c:cid,a:aid,d:did,remoteAddr,url
	// ldap {reqInfo}: remoteAddr,url
	RequestRestFormat = "[request]%s,%s,%s,%s,%s"

	// response rest log format.
	// [response]traceId,code,msg
	ResponseRestFormat = "[response]%s,%d,%s"
)

// gateway-config日志
var (
	// janus-gateway 配置操作相关日志.
	ConfigLogger = glog.Context(nil, glog.FileName{Name: "gateway-config"})
)

// gateway-stat日志
var (
	// janus-gateway 访问统计日志. 支持格式包括:
	// 1. StatFormat
	StatLogger = glog.Context(nil, glog.FileName{Name: "gateway-stat"})

	// stat log format.
	// traceId,serviceName,httpMethod,path,client,ok:isSuccess,code,ms:response-ms
	StatFormat = "%s,%s,%s,%s,%s,ok:%s,%d,ms:%g"
)

// gateway-error日志
var (
	// gateway错误日志.
	ErrorLogger = glog.Context(nil, glog.FileName{Name: "gateway-error"})

	// error log format.
	ErrorFormat = "%s,%s"
)

// gateway-limit日志
var (
	// gateway限流日志.
	LimitLogger = glog.Context(nil, glog.FileName{Name: "gateway-limit"})

	// traceId[limit type]service,httpMethod,path,limitInfo.
	LimitFormat = "%s[%s]%s,%s,%s,%s"
)

// gateway-default日志
var (
	// gateway默认日志文件, 排除上面的剩下的日志.
	DefaultLogger = glog.Context(nil, glog.FileName{Name: "gateway-default"})

	// default log format.
	DefaultFormat = "%s,%s"
)

// Logf is used to record format logs for RestLogger/ConfigLogger/
// StatLogger/DefaultLogger.
func Logf(logger *glog.Logger, format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Logef is used to record format error logs for ErrorLogger.
func Logef(logger *glog.Logger, format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// LogColorInfof logs in colored format for RestLogger/ConfigLogger/StatLogger/
// DefaultLogger.
func LogColorInfof(logger *glog.Logger, format string, args ...interface{}) {
	logger.Infof(col.GreenString(format, args...))
}

// LogColorErrorf logs in colored format for ErrorLogger.
func LogColorErrorf(logger *glog.Logger, format string, args ...interface{}) {
	logger.Errorf(col.RedString(format, args...))
}

// VLogf is used to record format logs for RestLogger/ConfigLogger/
// StatLogger/DefaultLogger guarded by the value of level.
func VLogf(logger *glog.Logger, level int32, format string, args ...interface{}) {
	if glog.V(glog.Level(level)) {
		logger.Infof(format, args...)
	}
}

// VLogef is used to record format error logs for ErrorLogger guarded by
// the value of level.
func VLogef(logger *glog.Logger, level int32, format string, args ...interface{}) {
	if glog.V(glog.Level(level)) {
		logger.Errorf(format, args...)
	}
}

// Flush be used to flush logs immediately.
func Flush() {
	glog.Flush()
}
