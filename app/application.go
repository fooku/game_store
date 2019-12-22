package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

// StartApplication เริ่มต้น application
func StartApplication() {
	fmt.Println("hello Gin")

	mapUrls()
	router.Run(":8080")
}
