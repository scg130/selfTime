package main

import (
	"fmt"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/gin-gonic/gin"
)

type TInt interface {
	int64 | int32 | int16 | int8 | int | uint64 | uint32 | uint16 | uint8 | uintptr | float32 | float64
}

func sum[T TInt](a T, b T) T {
	return a + b
}

func main() {
	var patches *gomonkey.Patches
	r := gin.Default() // 带 Logger & Recovery
	r.GET("/set", func(c *gin.Context) {
		tStr := c.DefaultQuery("t", "2025-01-01 12:00:00")
		fakeTime, _ := time.Parse("2006-01-02 15:04:05", tStr)

		patches = gomonkey.ApplyFunc(time.Now, func() time.Time {
			return fakeTime
		})
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	})
	r.GET("/reset", func(c *gin.Context) {
		if patches != nil {
			patches.Reset()
		}
	})
	r.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	r.Run(":8080")
}

