package main

import (
	"backend/server"
)

func main() {
	router := server.CreateServer()
	router.Run()
}
