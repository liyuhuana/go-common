package network_tcp

type IHandler interface {
	OnOpen(*Session)
	OnClose(*Session, bool)
	OnRequest(*Session, int32, []byte) (int32, []byte)
	OnPush(*Session, int32, []byte)
}
