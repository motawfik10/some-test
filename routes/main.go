package routes

import (
	authControllers "backend/controllers/auth"
	announcementRouter "backend/routes/announcements"
	quizRouter "backend/routes/quizzes"
	userRouter "backend/routes/users"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// InitializeRoutes initializes all the required routes for the application
func InitializeRoutes(e *echo.Echo) {
	// to enable sending requests from the frontend application
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "https://motawfik10.github.io"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	middleware.ErrJWTMissing.Message = "Please login"
	adminRouter := e.Group("/admin",
		middleware.JWT([]byte("secret")),
		authControllers.CheckAdmin)

	userRouter.InitializeRoutes(e, adminRouter)
	announcementRouter.InitializeRoutes(e, adminRouter)
	quizRouter.InitializeRoutes(e, adminRouter)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "From APache")
	})

}
