package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/parag08/go-errors/pkg/api/sandwich"
	"github.com/parag08/go-errors/pkg/app"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

func run() error {
	// create router dependecy
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(cors.Default())

	server := app.NewServer(router, sandwich.NewSandWichService())

	// start the server
	err := server.Run()

	if err != nil {
		return err
	}

	return nil
}
