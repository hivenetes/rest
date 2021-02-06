package app

import (
	"fmt"
	"net/http"
	"rest/cmd/controllers"
)

// StartApp ...
func StartApp() {
	fmt.Println("Starting Application")
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
