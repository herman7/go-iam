package main

import (
	"log"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 业务处理之前
		t := time.Now()
		c.Set("example", "123123")

		c.Next()

		// 之后

		latency := time.Since(t)
		log.Print(latency)

		// 访问我们发送的状态
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()

	// 认证
	// r.Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"}))

	// RequestID
	r.Use(requestid.New(requestid.WithGenerator(func() string {
		return "test"
	})))

	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)
	})
	r.Run(":8080")
}
