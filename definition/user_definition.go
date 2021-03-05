package definition

import (
	"github.com/liyuhuana/go-common/utils/math_util"
	"strconv"
)

type Uid int32

const (
	EmptyUid Uid = -1
)

func ConvertStringToUid(str string) Uid {
	id, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return EmptyUid
	}
	return Uid(id)
}

func (id Uid) Int32() int32 {
	return int32(id)
}

func (id Uid) String() string {
	return math_util.Int32ToStr(id.Int32())
}

func (id Uid) IsEmpty() bool {
	return id == EmptyUid
}

type UserType int32
const (
	UserTypeNone UserType = iota
	UserTypeNormal
)
