package integrate

const (
	// 公共头信息
	XSource     = "x-source"      // 请求来源。web,client(包括4个客户端)
	XClient     = "x-client"      // 客户端资源号。如：mga、mip
	XUid        = "x-uid"         // 用户id
	XCid        = "x-cid"         // 企业id
	XAid        = "x-aid"         // 个人账户id
	XSid        = "x-sid"         // 登录sessionId（client端使用），web端使用cookie
	XDid        = "x-did"         // 设备唯一id（客户端使用）
	XAppVersion = "x-app-version" // 客户端版本号， 如：3.4.0（客户端使用）
	XTs         = "x-ts"          // 请求时间戳，单位（毫秒）
	XSign       = "x-sign"        // 签名串。需要加签验证的接口必填
	XRequestId  = "x-request-id"  // 唯一请求id，跟踪日志使用
	XLocale     = "x-locale"      // 语言 简体中文:zh_CN; 繁体中文:zh_TW;英文:en_US
	XClientId   = "x-client-id"   // 开放平台请求所需的client id
	XCorpCode   = "x-corp-code"   // 开放平台请求所需的corp code

	Authorization = "Authorization" // Token authorization header.

	// 请求来源
	// Need to update IsKnownSource when adding a new resource
	ResourceWeb          = "web"           // web端请求
	ResourceClient       = "client"        // client端请求,包含windowns、mac、Android、ios4个客户端
	ResourceThird        = "third"         // Request from third party

	// 接到请求的时间
	RequestReceivedTime = "request-received-time"
	// http header:x-forwarded-for
	XForwardedFor = "x-forwarded-for"
)

// IsKnownResource returns if the given resource is known or not.
func IsKnownResource(s string) bool {
	return s == ResourceWeb ||
		s == ResourceClient ||
		s == ResourceThird
}
