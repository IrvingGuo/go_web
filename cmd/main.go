package main

import (
	"go_web/router"
)

func main() {
	router := router.InitRouter()
	router.Run("localhost:8080")
}
