package routes

import (
	"pengi-med-saas/core/envelope"
	"pengi-med-saas/core/logger"
	user_handlers "pengi-med-saas/features/users/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userHandler := user_handlers.NewUserHandler(db, logger.Log)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("", envelope.Handle(userHandler.GetUsers))
	}

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", envelope.Handle(userHandler.SignUp))
		authRoutes.POST("/login", envelope.Handle(userHandler.Login))
		authRoutes.POST("/refresh", envelope.Handle(userHandler.RefreshAuthToken))
		authRoutes.POST("/extend", envelope.Handle(userHandler.ExtendSession))
		authRoutes.POST("/validate", envelope.Handle(userHandler.ValidateBearerToken))
	}

}
