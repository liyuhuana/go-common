package definition

import (
	"github.com/liyuhuana/go-common/utils/math_util"
	"strconv"
)

type UserId int32

const (
	EmptyUserId UserId = -1
)

func ConvertStringToUserId(str string) UserId {
	id, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return EmptyUserId
	}
	return UserId(id)
}

func (id UserId) Int32() int32 {
	return int32(id)
}

func (id UserId) String() string {
	return math_util.Int32ToStr(id.Int32())
}

func (id UserId) IsEmpty() bool {
	return id == EmptyUserId
}

type UserType int32
const (
	UserTypeNone UserType = iota
	UserTypeNormal
)
