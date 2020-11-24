package definition

type PlayerId int64

func (id PlayerId) Int64() int64 {
	return int64(id)
}

type PlayerType int32
const (
	PlayerTypeNone PlayerType = iota
	PlayerTypeNormal
)
