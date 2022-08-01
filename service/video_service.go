package service

import "sstut/model"

// VideoService - Create an interface of related items that all have the methods "Save" and "FindAll"
type VideoService interface {
	Save(video model.Video) model.Video
	FindAll() []model.Video
}

// A group of videos
type videoService struct {
	videos []model.Video
}

// Save a new video
func (v *videoService) Save(video model.Video) model.Video {
	v.videos = append(v.videos, video)
	return video
}

// FindAll - Retrieve all videos
func (v *videoService) FindAll() []model.Video {
	return v.videos
}

func New() VideoService{
	return &videoService{}
}
