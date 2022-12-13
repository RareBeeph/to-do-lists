package main

import (
	"backend/controllers"
)

func main() {
	router := controllers.CreateServer()
	router.Run()
}
