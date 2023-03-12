package server

import (
	"github.com/gerardva/go-api/controllers"
	"github.com/gerardva/go-api/database"
	"github.com/gerardva/go-api/database/repository"
	"github.com/gerardva/go-api/handlers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	health := new(controllers.HealthController)
	r.GET("/health", health.Check)

	carRepo := repository.NewCarRepository(database.GetDatabase())
	carHandler := handlers.NewCarHandler(carRepo)

	carGroup := r.Group("car")
	{
		car := controllers.NewCarController(carHandler)
		carGroup.GET("/", car.GetAllCars)
		carGroup.POST("/", car.CreateCar)
		carGroup.GET("/:id", car.GetCarById)
		carGroup.DELETE("/:id", car.DeleteCar)
		carGroup.PATCH("/:id", car.UpdateCar)
	}

	r.Run()
}
