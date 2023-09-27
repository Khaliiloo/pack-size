package main

import (
	"fmt"

	"github.com/Khaliiloo/pack-size/api"
)

func main() {
	app := api.Api()
	port := 3000
	fmt.Printf("Server is running on port %d...\n", port)
	app.Listen(fmt.Sprintf(":%d", port))
}
