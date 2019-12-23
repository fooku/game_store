package users

import (
	"fmt"
	"gamestore/domain/users"
	"gamestore/services"
	"gamestore/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetUser รับผู้ใช้
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not Implement me!")
}

//CreateUser รับผู้ใช้
func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//Handle json error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//Handle json error
	// 	return
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user, result)
	c.JSON(http.StatusCreated, result)
}

// FindUser ค้นหาผู้ใช้
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not Implement me!")
}
