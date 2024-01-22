package error

import (
	"github.com/ocean386/common/error/rpc"
)

/**
常用通用固定错误
*/

type CodeError struct {
	rpc.Error
	// errCode uint32
	// errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() int32 {
	return e.ErrCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.ErrMsg
}

func NewErrCodeMsg(errCode int32, errMsg string) *CodeError {
	err := &CodeError{}
	err.ErrCode = errCode
	err.ErrMsg = errMsg
	return err
}
func NewErrCode(errCode int32) *CodeError {
	err := &CodeError{}
	err.ErrCode = errCode
	err.ErrMsg = MapErrMsg(ErrCode(errCode))
	return err
}

func NewErrMsg(errMsg string) *CodeError {
	err := &CodeError{}
	err.ErrCode = int32(StatusCodeError)
	err.ErrMsg = errMsg
	return err
}

func NewSystemErr() *CodeError {
	return NewErrCode(int32(StatusCodeError))
}

// 数据库-MySQL
func NewDBCreateErr() *CodeError {
	return NewErrCode(int32(StatusDBCreateError))
}

func NewDBRepeatCreateErr() *CodeError {
	return NewErrCode(int32(StatusDBRepeatCreateError))
}

func NewDBDeleteErr() *CodeError {
	return NewErrCode(int32(StatusDBDeleteError))
}

func NewDBUpdateErr() *CodeError {
	return NewErrCode(int32(StatusDBUpdateError))
}

func NewDBSelectErr() *CodeError {
	return NewErrCode(int32(StatusDBSelectError))
}

func NewDBCountErr() *CodeError {
	return NewErrCode(int32(StatusDBCountError))
}

/*
### RPC 使用方式如下
if err != nil && err != gorm.ErrRecordNotFound {
	st, _ := status.New(codes.Unavailable, err.Error()).WithDetails(xerr.NewSystemErr())
	return st.Err()
}

### API 使用方式如下
if err != nil {
	code := xerr.ParseErrCodeFromErr(err)
	resp.Code = int(code)
	resp.Msg = xerr.MapErrMsg(code)
	l.Logger.Errorf("GateWay  ErrCode[%d] ErrMsg[%s],error:%s", resp.Code, resp.Msg, err.Error())
	return resp, nil
}
*/
