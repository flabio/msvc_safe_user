package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/safe_msvc_user/insfractruture/routers"
)

func main() {
	app := fiber.New()
	// Custom CORS configuration
	// Enable CORS with specific settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080", // Specify the origin that is allowed to make requests
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	routers.NewAuthRouter(app)
	routers.NewUserRouter(app)
	app.Listen(":3005")

}
