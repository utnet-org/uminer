package errors

type CodeType int32

const (
	Success            CodeType = 0  // 成功
	InvalidArgument    CodeType = 1  // 请求参数非法
	OutOfRange         CodeType = 2  // 操作超过有效范围
	FailedPrecondition CodeType = 3  // 拒绝访问
	Unauthorized       CodeType = 4  // 未授权,请求没有操作的有效身份凭证
	PermissionDenied   CodeType = 5  // 没有访问权限
	NotFound           CodeType = 6  // 请求资源未找到
	AlreadyExists      CodeType = 7  // 资源已存在
	Aborted            CodeType = 8  // 操作被中止
	ResourceExhausted  CodeType = 9  // 资源已耗尽或已超出限制
	Unknown            CodeType = 10 // 不知名错误
	Internal           CodeType = 11 // 内部系统错误
	DataLoss           CodeType = 12 // 不可恢复的数据丢失或损坏
	Unimplemented      CodeType = 13 // 操作暂不支持
	Unavailable        CodeType = 14 // 服务不可用或过载拒绝
	DeadlineExceeded   CodeType = 15 // 请求超时
)

var codeTypeMsg = map[CodeType]string{
	Success:            "成功",
	InvalidArgument:    "请求参数非法",
	OutOfRange:         "操作超过有效范围",
	FailedPrecondition: "拒绝访问",
	Unauthorized:       "未授权,请求没有操作的有效身份凭证",
	PermissionDenied:   "没有访问权限",
	NotFound:           "请求资源未找到",
	AlreadyExists:      "资源已存在",
	Aborted:            "操作被中止",
	ResourceExhausted:  "资源已耗尽或已超出限制",
	Unknown:            "不知名错误",
	Internal:           "内部系统错误",
	DataLoss:           "不可恢复的数据丢失或损坏",
	Unimplemented:      "操作暂不支持",
	Unavailable:        "服务不可用或过载拒绝",
	DeadlineExceeded:   "请求超时",
}
