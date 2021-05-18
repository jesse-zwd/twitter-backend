package controller

import (
	"backend/service"

	"github.com/gin-gonic/gin"
)

// CreateComment godoc
// @Summary comment create
// @Description comment create
// @Tags Comment API
// @ID /comments
// @Accept  json
// @Produce  json
// @Param body body service.CreateCommentService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	var service service.CreateCommentService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteComment godoc
// @Summary comment delete
// @Description comment delete
// @Tags Comment API
// @ID /comment/:id
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /comment/:id [delete]
func DeleteComment(c *gin.Context) {
	service := service.DeleteCommentService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
