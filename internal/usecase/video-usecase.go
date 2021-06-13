package usecase

import (
	"fmt"
	"github.com/PHChenGit/Go-blog/internal/entity"
)

type VideoUseCase interface {
	Save(video entity.Video) entity.Video
	All() []entity.Video
}

type videoUseCase struct {
	videos []entity.Video
}

func NewVideoService() VideoUseCase {
	return &videoUseCase{}
}

func (useCase *videoUseCase) Save(video entity.Video) entity.Video {
	fmt.Printf("%+v", video)
	useCase.videos = append(useCase.videos, video)
	return video
}

func (useCase *videoUseCase) All() []entity.Video {
	return useCase.videos
}
