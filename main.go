package main

import "github.com/gin-gonic/gin"


func main() {
	r := setup()
	if err:=r.Run(":8080");err!=nil {
		panic(err)
	}
}

func setup() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", HelloWorld)
	return r
}

func HelloWorld(ctx *gin.Context) {
	ctx.JSON(200,gin.H{
		"message":"Hello World",
	})
}
