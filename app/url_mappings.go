package app

import (
	"gamestore/controllers/ping"
	"gamestore/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.CreateUser)
	// router.GET("/users/search", controllers.FindUser)
	router.POST("/users", users.CreateUser)

}
