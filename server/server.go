package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gerardva/go-api/controllers"
)

func Init() {
	r := gin.Default()
	health := new(controllers.HealthController)
	r.GET("/health", health.Check)

	carGroup := r.Group("car")
	{
		car := new(controllers.CarController)
		carGroup.GET("/", car.GetAllCars)
		carGroup.POST("/", car.CreateCar)
		carGroup.GET("/:id", car.GetCarById)
		carGroup.DELETE("/:id", car.DeleteCar)
		carGroup.PATCH("/:id", car.UpdateCar)
	}

	r.Run()
}