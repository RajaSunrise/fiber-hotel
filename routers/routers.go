package routers

import (
	"github.com/RajaSunrise/hotel/controllers"
	"github.com/RajaSunrise/hotel/middleware"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

// register routes
func setRegisterRoutes(app *fiber.App){
	app.Post("/api/v1/signup", controllers.Register)
}


// login routes
func setLoginRoutes(app *fiber.App){
	app.Post("/api/v1/login", controllers.Login)
}


func setUserRoutes(app *fiber.App) {
	app.Get("/api/v1/users", controllers.GetAllUsers)
	app.Get("/api/v1/users/:id", controllers.GetUserByID)
	app.Post("/api/v1/users", controllers.CreateUser)
	app.Put("/api/v1/users/:id", controllers.UpdateUser)
	app.Delete("/api/v1/users/:id", controllers.DeleteUser)
}

func setReviewRoutes(app *fiber.App) {
	app.Get("/api/v1/reviews", controllers.GetAllReview)
	app.Get("/api/v1/reviews/:id", controllers.GetReviewByID)
	app.Post("/api/v1/reviews", controllers.CreateReview)
	app.Put("/api/v1/reviews/:id", controllers.UpdateReview)
	app.Delete("/api/v1/reviews/:id", controllers.DeleteReview)
}

func setRoomRoutes(app *fiber.App){
	app.Get("/api/v1/rooms", controllers.GetAllRoom)
	app.Get("/api/v1/rooms/:id", controllers.GetRoomByID)
	app.Post("/api/v1/rooms", controllers.CreateRoom)
	app.Put("/api/v1/rooms/:id", controllers.UpdateRoom)
	app.Delete("api/v1/rooms/:id", controllers.DeleteRoom)
}

func setFacilityRoutes(app *fiber.App){
	app.Get("/api/v1/facilitys", controllers.GetAllFacility)
	app.Get("/api/v1/facilitys/:id", controllers.GetFacilityByID)
	app.Post("/api/v1/facilitys", controllers.CreateFacility)
	app.Put("/api/v1/facilitys/:id", controllers.UpdateFacility)
	app.Delete("/api/v1/facilitys/:id", controllers.DeleteFacility)
}

func setBookingRoutes(app *fiber.App){
	app.Get("/api/v1/bookings", controllers.GetAllBooking)
	app.Get("/api/v1/bookings/:id", controllers.GetBookingByID)
	app.Post("/api/v1/bookings", controllers.CreateBooking)
	app.Put("/api/v1/bookings/:id", controllers.UpdateBooking)
	app.Delete("/api/v1/bookings/:id", controllers.DeleteBooking)
}

func setPromotionRoutes(app *fiber.App){
	app.Get("/api/v1/promotions", controllers.GetAllPromotion)
	app.Get("/api/v1/promotions/:id", controllers.GetpromotionByID)
	app.Post("/api/v1/promotions", controllers.CreatePromotion)
	app.Put("/api/v1/promotions/:id", controllers.UpdatePromotion)
	app.Delete("/api/v1/promotions/:id", controllers.DeletePromotion)
}

func setEmployeeRoutes(app *fiber.App){
	app.Get("/api/v1/employees", controllers.GetAllEmployee)
	app.Get("/api/v1/employees/:id", controllers.GetEmployeeByID)
	app.Post("/api/v1/employees", controllers.CreateEmployee)
	app.Put("/api/v1/employees/:id", controllers.UpdateEmployee)
	app.Delete("/api/v1/employees/:id", controllers.DeleteEmployee)
}

func setContentRoutes(app *fiber.App){
	app.Get("/api/v1/contents", controllers.GetAllContent)
	app.Get("/api/v1/contents/:id", controllers.GetContentByID)
	app.Post("/api/v1/contents", controllers.CreateContent)
	app.Put("/api/v1/contents/:id", controllers.UpdateContent)
	app.Delete("/api/v1/contents/:id", controllers.DeleteContent)
}

func setLocationRoutes(app *fiber.App){
	app.Get("/api/v1/locations", controllers.GetAllLocation)
	app.Get("/api/v1/locations/:id", controllers.GetLocationByID)
	app.Post("/api/v1/locations", controllers.CreateLocation)
	app.Put("/api/v1/locations/:id", controllers.UpdateLocation)
	app.Delete("/api/v1/locations/:id", controllers.DeleteLocation)
}

func setPaymentRoutes(app *fiber.App){
	app.Get("/api/v1/payments", controllers.GetAllPayment)
	app.Get("/api/v1/payments/:id", controllers.GetPaymentByID)
	app.Post("/api/v1/payments", controllers.CreatePayment)
	app.Put("/api/v1/payments/:id", controllers.Updatepayment)
	app.Delete("/api/v1/payments/:id", controllers.DeletePayment)
}



func middlewareRoutes(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{
            JWTAlg: jwtware.HS256,
            Key:    middleware.JwtSecret,
        },
    }))
}


func SetRouters(app *fiber.App) {
	setRegisterRoutes(app) // register tidaK membutuhkan jwt token
	setLoginRoutes(app)


	middlewareRoutes(app) // taruh diatas agar bisa membuat semua routes di middleware
	setUserRoutes(app)
	setReviewRoutes(app)
	setRoomRoutes(app)
	setFacilityRoutes(app)
	setBookingRoutes(app)
	setPromotionRoutes(app)
	setEmployeeRoutes(app)
	setContentRoutes(app)
	setLocationRoutes(app)
	setPaymentRoutes(app)
}
