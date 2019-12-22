package users

import "github.com/gin-gonic/gin"

import "net/http"

//GetUser รับผู้ใช้
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not Implement me!")
}

//CreateUser รับผู้ใช้
func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not Implement me!")
}

// FindUser ค้นหาผู้ใช้
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not Implement me!")
}
