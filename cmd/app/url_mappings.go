package app

import (
	"rest/cmd/controllers"
)

// MapUrls ...
func MapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
