package main

import "github.com/gin-gonic/gin"

func entry() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	r.Run(":3000")
}
