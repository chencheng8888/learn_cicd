package learncicd

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()

	r.GET("/hello",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Hello World",
		})
	})

	if err:=r.Run(":8080");err!=nil {
		panic(err)
	}
}