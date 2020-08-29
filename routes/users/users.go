package users

import (
	usersController "backend/controllers/users"

	"github.com/labstack/echo"
)

// InitializeRoutes initializes all the required routes for the application
func InitializeRoutes(e *echo.Echo) {
	e.POST("/login", usersController.Login)
	e.POST("/register", usersController.Register)
	e.GET("/get-pending-students", usersController.GetPendingUsers)
}
