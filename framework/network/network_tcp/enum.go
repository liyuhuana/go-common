package network_tcp

type Pattern byte

const (
	Push Pattern = iota
	Request
	Response
	Ping
	Pong
	Sub
	Unsub
	Pub
)