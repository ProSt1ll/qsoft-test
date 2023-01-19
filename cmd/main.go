package main

import (
	"github.com/ProSt1ll/qsoft-test/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(internal.PingPong)

	router.GET("/when/:year", internal.HandlerWhen)

	router.Run("localhost:" + "8080")
}
