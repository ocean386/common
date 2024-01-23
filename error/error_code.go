package error

type ErrCode int32

// 成功返回
const OK ErrCode = 200

// 默认状态码
const (
	StatusCodeError   ErrCode = -1
	StatusCodeSuccess ErrCode = 200
)

// 数据库-MySQL相关状态码
const (
	StatusDBCreateError ErrCode = iota + 1000
	StatusDBRepeatCreateError
	StatusDBDeleteError
	StatusDBUpdateError
	StatusDBSelectError
	StatusDBCountError
)

// Redis相关状态码
const (
	StatusRedisError ErrCode = 1010
)

// 交易所平台相关状态码
const (
	StatusPlatformError ErrCode = 1100
)
