package main

import "tomdog/tdnet"

func main() {
	server := tdnet.NewServer()

	server.AddRouter(&tdnet.PingRouter{})
	server.AddRouter(&tdnet.TestRouter{})
	server.AddRouter(&tdnet.MultiClientRouter{})

	server.Serve()
}
