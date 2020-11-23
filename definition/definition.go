package definition

type ByteLenType int

const (
	UInt8ByteLen  ByteLenType = 1
	UInt16ByteLen ByteLenType = 2
	UInt32ByteLen ByteLenType = 4
	UInt64ByteLen ByteLenType = 8

	Int8ByteLen  ByteLenType = 1
	Int16ByteLen ByteLenType = 2
	Int32ByteLen ByteLenType = 4
	Int64ByteLen ByteLenType = 8
)

func (t ByteLenType) Int() int {
	return int(t)
}
