package main

import (
	controller "github.com/PHChenGit/Go-blog/internal/delivery/http"
	"github.com/PHChenGit/Go-blog/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	videoService    usecase.VideoUseCase       = usecase.NewVideoService()
	videoController controller.VideoController = controller.New(videoService)
)

func main()  {
	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.All())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":9888")
}

