package main

import "github.com/liyuhuana/go-common/network/network_tcp"

func main() {
	s := network_tcp.NewServer("", nil)
	s.Start()
}
