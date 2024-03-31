package route

import (
	"fastfooducate/module/feature/auth"
	users "fastfooducate/module/feature/user"
	user "fastfooducate/module/feature/user/domain"
	"fastfooducate/utils/token"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface, userService user.UserServiceInterface) {
	users.InitializeUser(db)
	users.SetupRoutesUser(app)
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	// schedule.InitializeSchedule(db)
	// schedule.SetupRoutesSchedule(app, jwt, userService)
}
