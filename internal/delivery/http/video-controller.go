package http

import (
	"fmt"
	"github.com/PHChenGit/Go-blog/internal/entity"
	"github.com/PHChenGit/Go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	All() []entity.Video
}

type controller struct {
	service usecase.VideoUseCase
}

func New(service usecase.VideoUseCase) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	fmt.Printf("%+v", video)
	c.service.Save(video)
	return video
}

func (c *controller) All() []entity.Video {
	return c.service.All()
}