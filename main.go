package main

import (
	"net/http"

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
	r.Static("/css", "./templates/css")
	r.LoadHTMLGlob(("templates/*.html"))

	apiRoutes := r.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, http.StatusOK)
			}

		})
	}

	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
