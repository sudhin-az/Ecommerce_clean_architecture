package routes

import (
	"ecommerce_clean_architecture/pkg/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handlers.UserHandler) {

	router.POST("/usersignup", userHandler.UserSignUp)
	router.POST("/verify-otp/:email", userHandler.VerifyOTP)
	router.POST("/resend-otp/", userHandler.ResendOTP)
	router.POST("/userlogin", userHandler.UserLogin)
}
func AuthRoutes(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	router.GET("/google/login", authHandler.GoogleLogin)
	router.GET("/google/callback", authHandler.GoogleCallback)
}
