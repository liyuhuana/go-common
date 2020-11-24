package net_error

type NetError struct {
	// 错误码
	ErrCode int32
	// 错误信息
	ErrMsg string
}

func NewError(errCode int32, errMsg string) *NetError {
	return &NetError{
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}