package definition

import (
	"github.com/liyuhuana/go-common/utils/math_util"
	"strconv"
)

type PlayerId int64

const (
	EmptyPlayerId PlayerId = -1
)

func ConvertStringToPlayerId(str string) PlayerId {
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return EmptyPlayerId
	}
	return PlayerId(id)
}

func (id PlayerId) Int64() int64 {
	return int64(id)
}

func (id PlayerId) String() string {
	return math_util.Int64ToStr(id.Int64())
}

func (id PlayerId) IsEmpty() bool {
	return id == EmptyPlayerId
}

type PlayerType int32
const (
	PlayerTypeNone PlayerType = iota
	PlayerTypeNormal
)
