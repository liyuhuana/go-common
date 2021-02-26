package definition

import (
	"github.com/liyuhuana/go-common/utils/math_util"
	"strconv"
)

type PlayerId int32

const (
	EmptyPlayerId PlayerId = -1
)

func ConvertStringToPlayerId(str string) PlayerId {
	id, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return EmptyPlayerId
	}
	return PlayerId(id)
}

func (id PlayerId) Int32() int32 {
	return int32(id)
}

func (id PlayerId) String() string {
	return math_util.Int32ToStr(id.Int32())
}

func (id PlayerId) IsEmpty() bool {
	return id == EmptyPlayerId
}

type PlayerType int32
const (
	PlayerTypeNone PlayerType = iota
	PlayerTypeNormal
)
