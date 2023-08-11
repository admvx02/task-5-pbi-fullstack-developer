package router

import (
	"app/controllers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	users := v1.Group("/users")
	users.POST("/register", controllers.RegisterUser)
	users.GET("/login", controllers.LoginUser)
	users.PUT("/:userId", middlewares.AuthMiddleware(), controllers.UpdateUser)
	users.DELETE("/:userId", middlewares.AuthMiddleware(), controllers.DeleteUser)

	photos := v1.Group("/photos")
	photos.Use(middlewares.AuthMiddleware())
	photos.POST("/", controllers.CreatePhoto)
	photos.GET("/", controllers.GetPhotos)
	photos.PUT("/:photoId", controllers.UpdatePhoto)
	photos.DELETE("/:photoId", controllers.DeletePhoto)

	return r
}
