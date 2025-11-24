package common

import "errors"

// Common error variables
var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrCodeTimeout         = errors.New("verification code timeout")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrPhoneNumberNotFound = errors.New("phone number not found for customer") // Added
	ErrPhoneNumberExpired  = errors.New("phone number assignment expired")     // Added
	ErrInvalidTimeRange    = errors.New("invalid time range")                // Added
)

// ErrorCode 统一错误码类型
type ErrorCode int

// 通用错误码定义
const (
	// 成功响应
	CodeSuccess ErrorCode = 200

	// 客户端错误 4xxxx
	CodeBadRequest   ErrorCode = 40001 // 请求参数错误
	CodeUnauthorized ErrorCode = 40101 // 未授权，token错误或过期
	CodeForbidden    ErrorCode = 40301 // 禁止访问，没有权限
	CodeNotFound     ErrorCode = 40401 // 资源不存在

	// 服务端错误 5xxxx
	CodeInternalError ErrorCode = 50001 // 服务器内部错误

	// 业务错误 6xxxx
	CodeInsufficientBalance ErrorCode = 60001 // 客户余额不足
	CodeInvalidSecret       ErrorCode = 60101 // 客户密钥错误
	CodeInvalidPhone        ErrorCode = 60201 // 客户手机号错误
	CodeInvalidBusinessType ErrorCode = 60301 // 业务类型错误
	CodePhoneNotFound       ErrorCode = 60401 // 客户手机号不存在
	CodeCodeNotFound        ErrorCode = 60501 // 验证码不存在
	CodeThirdPartyError     ErrorCode = 60601 // 第三方服务错误
	CodePhoneExists         ErrorCode = 60701 // 客户手机号已存在
	CodeGetCodeFailed       ErrorCode = 60801 // 获取验证码失败
	CodeCodeTimeout         ErrorCode = 60802 // 验证码获取超时
	CodeUserExists          ErrorCode = 60901 // 用户已存在
	CodeInvalidCredentials  ErrorCode = 61001 // 用户名或密码错误
	CodePhoneNumberExpired  ErrorCode = 61101 // 手机号分配已过期 (Added)
	CodeProviderNotFound    ErrorCode = 61201 // 服务商未找到或未启用
	CodeProviderBusinessNotFound ErrorCode = 61202 // 服务商业务类型未配置
	CodeNoProviderAvailable ErrorCode = 61203 // 没有可用的服务商
)

// Error 实现 error 接口
func (e ErrorCode) Error() string {
	return e.Message()
}

// Message 返回错误码对应的中文消息
func (e ErrorCode) Message() string {
	messages := map[ErrorCode]string{
		CodeSuccess:             "请求成功",
		CodeBadRequest:          "请求参数错误",
		CodeUnauthorized:        "未授权，token错误或过期",
		CodeForbidden:           "禁止访问，没有权限",
		CodeNotFound:            "资源不存在",
		CodeInternalError:       "服务器内部错误",
		CodeInsufficientBalance: "客户余额不足",
		CodeInvalidSecret:       "客户密钥错误",
		CodeInvalidPhone:        "客户手机号错误",
		CodeInvalidBusinessType: "业务类型错误",
		CodePhoneNotFound:       "客户手机号不存在",
		CodeCodeNotFound:        "验证码不存在",
		CodeThirdPartyError:     "第三方服务错误",
		CodePhoneExists:         "客户手机号已存在",
		CodeGetCodeFailed:       "获取验证码失败",
		CodeCodeTimeout:         "验证码获取超时",
		CodeUserExists:          "用户已存在",
		CodeInvalidCredentials:  "用户名或密码错误",
		CodePhoneNumberExpired:  "手机号分配已过期", // Added
		CodeProviderNotFound:    "服务商未找到或未启用",
		CodeProviderBusinessNotFound: "服务商业务类型未配置",
		CodeNoProviderAvailable: "没有可用的服务商",
	}

	if msg, ok := messages[e]; ok {
		return msg
	}
	return "未知错误"
}

// HTTPStatus 返回错误码对应的HTTP状态码
func (e ErrorCode) HTTPStatus() int {
	switch {
	case e >= 200 && e < 300:
		return 200
	case e >= 40100 && e < 40200:
		return 401
	case e >= 40300 && e < 40400:
		return 403
	case e >= 40400 && e < 40500:
		return 404
	case e >= 40000 && e < 41000:
		return 400
	case e >= 50000 && e < 60000:
		return 500
	default:
		return 400 // 业务错误默认返回400
	}
}
