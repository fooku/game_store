package app

import "gamestore/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
