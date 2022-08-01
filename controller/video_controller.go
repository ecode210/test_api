package controller

import (
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
	"sstut/model"
	"sstut/service"
	"sstut/validator"
)

// VideoController - Create an interface of related items that all have the methods "Save" and "FindAll"
type VideoController interface {
	FindAll() []model.Video
	Save(c *gin.Context) error
}

type controller struct {
	service service.VideoService
}

// FindAll - Retrieve all videos
func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

// Save - Gets JSON request details and save new video from details
func (c *controller) Save(ctx *gin.Context) error{
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	// Validates the given struct based on the tags
	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return  nil
}

var validate *validator2.Validate

func New(service service.VideoService) VideoController{
	// Creates a new validator and registers custom validations
	validate = validator2.New()
	err := validate.RegisterValidation("is-cool", validator.CoolValidation)
	if err != nil {
		return nil
	}

	return &controller{service: service}
}
