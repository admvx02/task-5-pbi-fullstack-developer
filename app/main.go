package main

import (
	"app/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
