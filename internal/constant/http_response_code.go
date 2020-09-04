package constant

// HTTPResponseCode http response code
type HTTPResponseCode uint

const (
	// HTTPResponseCodeNotLogin 用户未登录
	HTTPResponseCodeNotLogin HTTPResponseCode = 1001

	// HTTPResponseCodeInvalidSession 无效的Session
	HTTPResponseCodeInvalidSession HTTPResponseCode = 1002

	// HTTPResponseCodeInvalidToken 无效的Token
	HTTPResponseCodeInvalidToken HTTPResponseCode = 1003

	// HTTPResponseCodeInvalidParams 请求参数错误
	HTTPResponseCodeInvalidParams HTTPResponseCode = 1004

	// HTTPResponseCodeServeBusy 服务繁忙
	HTTPResponseCodeServeBusy HTTPResponseCode = 1005

	// HTTPResponseCodeInvalidAuth 用户未授权
	HTTPResponseCodeInvalidAuth HTTPResponseCode = 1006
)

const (
	// HTTPResponseCodeSuccess Success
	HTTPResponseCodeSuccess HTTPResponseCode = 2000
)

const (
	// HTTPResponseCodeServeError 服务内部错误
	HTTPResponseCodeServeError HTTPResponseCode = 5000
)

var httpResponseCodeMap = make(map[HTTPResponseCode]string)

func init() {
	// 100*编码
	httpResponseCodeMap[HTTPResponseCodeNotLogin] = "用户未登陆"
	httpResponseCodeMap[HTTPResponseCodeInvalidSession] = "无效的Session"
	httpResponseCodeMap[HTTPResponseCodeInvalidToken] = "无效的Token"
	httpResponseCodeMap[HTTPResponseCodeInvalidParams] = "请求参数错误"
	httpResponseCodeMap[HTTPResponseCodeServeBusy] = "服务繁忙"
	httpResponseCodeMap[HTTPResponseCodeInvalidAuth] = "用户未授权"

	// 200*编码
	httpResponseCodeMap[HTTPResponseCodeSuccess] = "Success"

	// 500*编码
	httpResponseCodeMap[HTTPResponseCodeServeError] = "服务内部错误"
}

func (key HTTPResponseCode) String() string {

	if v, ok := httpResponseCodeMap[key]; ok {
		return v
	}
	return "unknown http response code"
}
