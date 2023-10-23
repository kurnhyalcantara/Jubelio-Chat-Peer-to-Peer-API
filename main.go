package main

import (
	"github.com/kurnhyalcantara/Jubelio-Chat-Peer-to-Peer-API/app/config"
	"github.com/kurnhyalcantara/Jubelio-Chat-Peer-to-Peer-API/cmd/server"
)

func main() {
	config.LoadConfigs()
	server.Serve()
}