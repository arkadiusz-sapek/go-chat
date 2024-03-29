package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	auth "github.com/ArkadiuszSa/go-chat/modules/auth/middleware"
	user "github.com/ArkadiuszSa/go-chat/modules/user/services"
)

//Routes - return user routes
func Routes(router *gin.Engine) {
	router.GET("/", auth.AuthorizeJWT(), welcome)
	router.GET("/users", user.GetAllUsers)
	// router.POST("/user", user.CreateUser)
	router.GET("/user/:userID", user.GetSingleUser)
	router.DELETE("/user/:userID", user.DeleteUser)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
