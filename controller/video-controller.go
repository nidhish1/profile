package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nidhish1/profile/entity"
	"github.com/nidhish1/profile/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return c.service.Save(video)
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
