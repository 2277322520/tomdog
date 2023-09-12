package main

import "tomdog/tdnet"

func main() {
	server := tdnet.NewServer("testServer")

	server.Serve()
}
