package main

import "github.com/liyuhuana/go-common/framework/network/network_tcp"

func main() {
	s := network_tcp.NewServer("", nil)
	s.Start()
}
