package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest/cmd/services"
	"rest/cmd/utils"
	"strconv"
)

// GetUser ...
func GetUser(c *gin.Context) {

	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v not found", userID),
			StatusCode: http.StatusNotFound,
		}
		utils.RespondErr(c, apiErr)
		return
	}
	user, apiErr := services.UserService.GetUser(userID)
	if apiErr != nil {
		utils.RespondErr(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
