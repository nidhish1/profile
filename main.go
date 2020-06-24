package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nidhish1/profile/controller"
	"github.com/nidhish1/profile/middlewares"
	"github.com/nidhish1/profile/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	r := gin.Default()
	r.Use(middlewares.BasicAuth())

	r.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	r.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
