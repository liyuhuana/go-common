package network_tcp

type IHandler interface {
	OnOpen(*Session)
	OnClose(*Session, bool)
	OnRequest(*Session, uint32, []byte)
	OnPush(*Session, uint32, []byte)
}
