package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sstut/controller"
	"sstut/middleware"
	"sstut/service"
)

var(
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func main() {
	router := gin.New()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		//middleware.Logger(),
		middleware.BasicAuth(),
		// Adds extra details about the request
		//gindump.Dump(),
		)

	router.GET("/videos", func(c *gin.Context) {
		c.JSON(http.StatusOK, videoController.FindAll())
	})
	router.POST("/videos", func(c *gin.Context) {
		err := videoController.Save(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}else {
			c.JSON(http.StatusCreated, gin.H{
				"message": "Video saved!",
			})
		}
	})

	router.Run(":6969")
}
