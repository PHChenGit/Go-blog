package service

import (
	"fmt"
	"github.com/PHChenGit/Go-blog/internal/entity"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	All() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	fmt.Printf("%+v", video)
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) All() []entity.Video {
	return service.videos
}
