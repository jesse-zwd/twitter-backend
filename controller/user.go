package controller

import (
	"backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister godoc
// @Summary user register
// @Description user register
// @Tags User API
// @ID /user/register
// @Accept  json
// @Produce  json
// @Param body body service.UserRegisterService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin godoc
// @Summary user login
// @Description user login
// @Tags User API
// @ID /user/login
// @Accept  json
// @Produce  json
// @Param body body service.UserLoginService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogout godoc
// @Summary user logout
// @Description user logout
// @Tags User API
// @ID /user/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /user/logout [delete]
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, service.Response{
		Code: 0,
		Msg:  "Logged out",
	})
}

// ListUser godoc
// @Summary user list
// @Description user list
// @Tags User API
// @ID /users
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{data=service.DataList} "success"
// @Router /users [get]
func ListUser(c *gin.Context) {
	service := service.ListUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserProfile godoc
// @Summary user profile
// @Description user profile
// @Tags User API
// @ID /user/:id
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /user/:id [get]
func UserProfile(c *gin.Context) {
	res := service.GetProfile(c, c.Param("id"))
	c.JSON(200, res)
}

// UpdateUser godoc
// @Summary user update
// @Description user update
// @Tags User API
// @ID /user
// @Accept  json
// @Produce  json
// @Param body body service.UpdateUserService true "body"
// @Success 200 {object} service.Response{} "success"
// @Router /user [put]
func UpdateUser(c *gin.Context) {
	service := service.UpdateUserService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// SearchUser godoc
// @Summary user search
// @Description user search
// @Tags User API
// @ID /searchUser/:searchItem
// @Accept  json
// @Produce  json
// @Success 200 {object} service.Response{} "success"
// @Router /searchUser/:searchItem [get]
func SearchUser(c *gin.Context) {
	res := service.SearchUser(c, c.Param("searchItem"))
	c.JSON(200, res)
}
