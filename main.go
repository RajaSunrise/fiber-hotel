package main


import (
	"github.com/RajaSunrise/hotel/database/migrations"
	"github.com/RajaSunrise/hotel/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    
    migrations.Migrate()
    
    routers.SetRouters(app)

    app.Listen(":8080")
}
