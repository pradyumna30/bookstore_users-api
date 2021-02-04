package app

import (
	"github.com/pradyumna30/bookstore_users-api/controllers/ping_controller"
	"github.com/pradyumna30/bookstore_users-api/controllers/user_controller"
)

func mapUrl() {

	router.GET("/ping", ping_controller.Ping)

	router.GET("/user/:id", user_controller.GetUser)
	router.POST("/users", user_controller.CreateUser)
	//router.GET("/user/search", controllers.SearchUser)
}
