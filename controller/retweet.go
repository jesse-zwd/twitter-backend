package controller

import (
	"backend/service"

	"github.com/gin-gonic/gin"
)

// CreateRetweet godoc
// @Summary retweet create
// @Description retweet create
// @Tags Retweet API
// @ID /retweet
// @Accept  json
// @Produce  json
// @Param body body service.RetweetService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /retweet [post]
func CreateRetweet(c *gin.Context) {
	var service service.RetweetService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteRetweet godoc
// @Summary retweet delete
// @Description retweet delete
// @Tags Retweet API
// @ID /retweet/:id
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /retweet/:id [delete]
func DeleteRetweet(c *gin.Context) {
	res := service.DeleteRetweet(c, c.Param("id"))
	c.JSON(200, res)
}
