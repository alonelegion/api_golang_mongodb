package routes

import (
	"github.com/alonelegion/api_golang_mongodb/controllers"
	"github.com/alonelegion/api_golang_mongodb/middleware"
	"github.com/alonelegion/api_golang_mongodb/services"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (ac *AuthRouteController) AuthRoute(rg *gin.RouterGroup, userService services.UserService) {
	router := rg.Group("/auth")

	router.POST("/register", ac.authController.SignUpUser)
	router.POST("/login", ac.authController.SignInUser)
	router.GET("/refresh", ac.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(userService), ac.authController.LogoutUser)
	router.GET("/verifyemail/:verificationCode", ac.authController.VerifyEmail)
}
