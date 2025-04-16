package controllers

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/services"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (r *UserController) Register(c *gin.Context) {
	var payload models.RegisterRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	err := r.userService.Register(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Status: true,
		Data:   nil,
		Error:  "",
	})
	return
}

func (r *UserController) Login(c *gin.Context) {
	var payload models.LoginRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
	}
	user, err := r.userService.Login(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, utils.Response{
		Status: true,
		Data:   user,
		Error:  "",
	})
}
