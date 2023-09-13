package main

import "tomdog/tdnet"

func main() {
	server := tdnet.NewServer("testServer")
	
	server.AddRouter(&tdnet.PingRouter{})
	
	server.Serve()
}
