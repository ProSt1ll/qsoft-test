package internal

import (
	"github.com/gin-gonic/gin"
)

func PingPong(c *gin.Context) {
	if c.Request.Header.Get("X-PING") == "ping" {
		c.Header("X-PONG", "pong")
	}
}
