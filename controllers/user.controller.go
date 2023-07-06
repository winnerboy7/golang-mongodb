package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/winnerboy7/golang-mongodb/models"
	"github.com/winnerboy7/golang-mongodb/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*models.DBResponse)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": models.FilteredResponse(currentUser)}})
}
