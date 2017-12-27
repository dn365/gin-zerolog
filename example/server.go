package main

import (
	"os"

	"github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// isConsole true
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	r := gin.New()
	r.Use(ginzerolog.Logger("gin"))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, gin-zerolog example")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":3000")
}
