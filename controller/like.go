package controller

import (
	"backend/service"

	"github.com/gin-gonic/gin"
)

// CreateLike godoc
// @Summary like create
// @Description like create
// @Tags Like API
// @ID /like
// @Accept  json
// @Produce  json
// @Param body body service.LikeService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /like [post]
func CreateLike(c *gin.Context) {
	var service service.LikeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteLike godoc
// @Summary like delete
// @Description like delete
// @Tags Like API
// @ID /like/:id
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /like/:id [delete]
func DeleteLike(c *gin.Context) {
	res := service.DeleteLike(c, c.Param("id"))
	c.JSON(200, res)
}
