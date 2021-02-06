package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest/cmd/services"
	"rest/cmd/utils"
	"strconv"
)

// GetUser ...
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {

		apiErr := &utils.ApplicationError{
			Message: fmt.Sprintf("user %v not found", userID),
			StatusCode: http.StatusNotFound,
		}
		jsonValue, _  := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _  := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	payload, _ := json.Marshal(user)
	resp.Write(payload)
}
