package definition

type PlayerId int64

const (
	EmptyPlayerId PlayerId = -1
)

func (id PlayerId) Int64() int64 {
	return int64(id)
}

func (id PlayerId) IsEmpty() bool {
	return id == EmptyPlayerId
}

type PlayerType int32
const (
	PlayerTypeNone PlayerType = iota
	PlayerTypeNormal
)
