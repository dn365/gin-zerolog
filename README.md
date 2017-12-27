# Gin-Zerolog

[Gin](https://github.com/gin-gonic/gin) middleware for Logging with
[zerolog](https://github.com/rs/zerolog). It is meant as drop in
replacement for the default logger used in Gin.

## Project Context

Gin-Zerolog replaces the default logger of [Gin](https://github.com/gin-gonic/gin) to use
[zerolog](https://github.com/rs/zerolog).

## Requirements

Gin-Zerolog uses the following [Go](https://golang.org/) packages as
dependencies:

- github.com/gin-gonic/gin
- github.com/rs/zerolog

## Installation

Assuming you've installed Go and Gin, run this:

    go get github.com/dn365/gin-zerolog

## Usage
### Example

Start your webapp to log to STDERR:

```go
package main

import (
	"github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
)

func main() {
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
```
STDERR output of the example above. Lines with prefix "[Gin-Debug]"
are hardcoded output of Gin and will not appear in your logfile:

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:	export GIN_MODE=release
    - using code:	gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
    [GIN-debug] GET    /ping                     --> main.main.func2 (2 handlers)
    [GIN-debug] Listening and serving HTTP on :3000
    {"time":"2017-12-27T14:22:15+08:00","level":"info","ser_name":"gin","method":"GET","path":"/","resp_time":0.043148,"status":200,"client_ip":"::1","message":"Request"}
    {"time":"2017-12-27T14:22:24+08:00","level":"info","ser_name":"gin","method":"GET","path":"/ping","resp_time":0.012409,"status":200,"client_ip":"::1","message":"Request"}
    {"time":"2017-12-27T14:22:36+08:00","level":"warn","ser_name":"gin","method":"GET","path":"/hello","resp_time":0.00149,"status":404,"client_ip":"::1","message":"Request"}



```go
// isConsole true
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

```

STDERR output of the example above. Lines with prefix "[Gin-Debug]"
are hardcoded output of Gin and will not appear in your logfile:

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:	export GIN_MODE=release
    - using code:	gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /                         --> main.main.func1 (2 handlers)
    [GIN-debug] GET    /ping                     --> main.main.func2 (2 handlers)
    [GIN-debug] Listening and serving HTTP on :3000
    2017-12-27T14:18:40+08:00 |INFO| Request client_ip=::1 method=GET path=/ resp_time=0.063933 ser_name=gin status=200
    2017-12-27T14:18:45+08:00 |INFO| Request client_ip=::1 method=GET path=/ping resp_time=0.023901 ser_name=gin status=200
    2017-12-27T14:18:49+08:00 |WARN| Request client_ip=::1 method=GET path=/hello resp_time=0.001331 ser_name=gin status=404
