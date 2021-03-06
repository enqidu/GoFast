package routes

import (
	"github.com/Gogotchuri/GoFast/app/controllers/auth"
	"github.com/Gogotchuri/GoFast/app/middleware"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

/*InitializeRoutes given a fiber.App initializes routes on it*/
func InitializeRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080", "*"},
	}))
	initAPIRoutes(app)
	initStaticRoutes(app)
}

func initAPIRoutes(app *fiber.App) {
	apiGroup := app.Group("/api") //TODO:ratelimiters
	//TODO: initialize other api routes here (Decomposition recommended)
	initAuthRoutes(apiGroup)
	initUserRoutes(apiGroup)
	//Last api route maps all unmapped routes to one handler
	apiGroup.Get("/*", func(c *fiber.Ctx) {
		c.Send("First api route (index).")
	})
	apiGroup.Post("/*", func(c *fiber.Ctx) {
		c.Send("First api route (index).")
	})
}

func initAuthRoutes(rg *fiber.Group) {
	//On platform authorization routes
	rg.Post("/sign-in", auth.SignIn)
	rg.Post("/sign-up", auth.SignUp)
	rg.Post("/token/refresh", auth.RefreshJWTTokens)
	rg.Post("/forgot-password", auth.PasswordForgotten)
	rg.Post("/password-reset", auth.ResetPassword)
	//Social authorization routes
	rg.Get("/auth/:provider/callback", auth.CallbackHandler)
	rg.Post("/auth/:provider", auth.RedirectHandler)
}

func initUserRoutes(rg *fiber.Group) {
	userGroup := rg.Group("/user")
	//Only authorized users will be able to enter router under this group
	userGroup.Use(middleware.IsAuthorized())
	userGroup.Get("/details", auth.GetUserDetails)
	userGroup.Post("/resend-code", auth.ReSendVerificationMail)
	userGroup.Post("/verify-email", auth.VerifyEmail)
	userGroup.Post("/logout", auth.Logout)
}

func initStaticRoutes(app *fiber.App) {
	app.Static("/", "./public/")
	app.Get("/*", func(ctx *fiber.Ctx) {
		ctx.SendFile("./public/index.html")
	})
}
