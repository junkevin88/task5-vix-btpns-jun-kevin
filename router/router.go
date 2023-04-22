package router

import (
	"btpn-backend-go/controller"
	"btpn-backend-go/middleware"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World"})
	})

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.UserRegister)
		userRouter.POST("/login", controller.UserLogin)
		userRouter.PUT("/:userId", controller.UserUpdate)
		userRouter.DELETE("/:userId", controller.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())

		photoRouter.POST("/", controller.CreatePhoto)
		photoRouter.GET("/", controller.ListPhoto)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}

	return r
}
