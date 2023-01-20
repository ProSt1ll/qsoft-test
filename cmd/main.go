package main

import (
	"flag"
	"github.com/ProSt1ll/qsoft-test/internal"
	"github.com/gin-gonic/gin"
	"log"
)

var port string

func init() {
	//flags
	flag.StringVar(&port, "Port", "8080", "server port")
	flag.Parse()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//middlewares settings
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	router.Use(internal.PingPong)

	//handlers settings
	router.Any("/when/:year", internal.HandlerWhen)

	//start server
	log.Fatal(router.Run("localhost:" + port))
}
