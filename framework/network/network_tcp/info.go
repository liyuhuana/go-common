package network_tcp

type ServerInfo struct {
	host string
}

func (info ServerInfo) GetHost() string {
	return info.host
}
