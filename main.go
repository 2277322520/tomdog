package main

import "tomdog/tdnet"

func main() {
	server := tdnet.NewServer()
	
	server.AddRouter(&tdnet.PingRouter{})
	
	server.Serve()
}
