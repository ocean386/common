package error

import (
	"github.com/ocean386/common/error/rpc"
	"google.golang.org/grpc/status"
)

var message map[ErrCode]string

// 定义错误码对应的msg
func init() {
	message = make(map[ErrCode]string)

	message[StatusCodeError] = "System error"
	message[StatusCodeSuccess] = "Success"

	//数据库-MySQL
	message[StatusDBCreateError] = "Database create error"
	message[StatusDBRepeatCreateError] = "Database repeat create"
	message[StatusDBDeleteError] = "Database delete error"
	message[StatusDBUpdateError] = "Database update error"
	message[StatusDBSelectError] = "Database select error"
	message[StatusDBCountError] = "Database count error"

}

// 根据错误码映射msg内容
func MapErrMsg(error ErrCode) string {
	if msg, ok := message[error]; ok {
		return msg
	} else {
		return "The server is missing, try again later"
	}
}

// 根据rpc Error转换int32 ErrCode
func ParseErrCodeFromErr(err error) ErrCode {
	if rpcErr, ok := status.FromError(err); ok {
		for _, v := range rpcErr.Details() {
			if detail, ok := v.(*rpc.Error); ok {
				return ErrCode(detail.ErrCode)
			}
		}
	}
	return StatusCodeError
}
