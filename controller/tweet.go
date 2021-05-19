package controller

import (
	"backend/service"

	"github.com/gin-gonic/gin"
)

// CreateTweet godoc
// @Summary tweet create
// @Description tweet create
// @Tags Tweet API
// @ID /tweets
// @Accept  json
// @Produce  json
// @Param body body service.CreateTweetService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /tweets [post]
func CreateTweet(c *gin.Context) {
	var service service.CreateTweetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteTweet godoc
// @Summary tweet delete
// @Description tweet delete
// @Tags Tweet API
// @ID /tweet/:id
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /tweet/:id [delete]
func DeleteTweet(c *gin.Context) {
	service := service.DeleteTweetService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

// ListTweet godoc
// @Summary tweet list
// @Description tweet list
// @Tags Tweet API
// @ID /tweets
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /tweets [get]
func ListTweet(c *gin.Context) {
	service := service.ListTweetService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// TweetByID godoc
// @Summary tweet by id
// @Description tweet by id
// @Tags Tweet API
// @ID /tweet/:id
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /tweet/:id [get]
func TweetByID(c *gin.Context) {
	res := service.GetTweetByID(c, c.Param("id"))
	c.JSON(200, res)
}

// SearchTweetByText godoc
// @Summary tweet search by text
// @Description tweet search by text
// @Tags Tweet API
// @ID /searchTweetByText/:searchItem
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /searchTweetByText/:searchItem [get]
func SearchTweetByText(c *gin.Context) {
	res := service.SearchTweetByText(c, c.Param("searchItem"))
	c.JSON(200, res)
}

// SearchTweetByTags godoc
// @Summary tweet search by tags
// @Description tweet search by tags
// @Tags Tweet API
// @ID /searchTweetByTags/:searchItem
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /searchTweetByTags/:searchItem [get]
func SearchTweetByTags(c *gin.Context) {
	res := service.SearchTweetByTags(c, c.Param("searchItem"))
	c.JSON(200, res)
}
