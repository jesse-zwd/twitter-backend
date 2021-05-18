package controller

import (
	"backend/service"

	"github.com/gin-gonic/gin"
)

// CreateFollow godoc
// @Summary follow create
// @Description follow create
// @Tags Follow API
// @ID /follow
// @Accept  json
// @Produce  json
// @Param body body service.FollowService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /follow [post]
func CreateFollow(c *gin.Context) {
	var service service.FollowService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteFollow godoc
// @Summary follow delete
// @Description follow delete
// @Tags Follow API
// @ID /follow
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /follow [delete]
func DeleteFollow(c *gin.Context) {
	res := service.DeleteFollow(c, c.Param("id"))
	c.JSON(200, res)
}
