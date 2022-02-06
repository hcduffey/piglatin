package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hcduffey/piglatin/translator"
	"net/http"
)

type Input struct {
	Message string
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/translate", func(context *gin.Context) {
		var json Input

		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			context.JSON(200, gin.H{
				"message": translator.Translate(json.Message),
			})
		}
	})
	r.Run()
}
