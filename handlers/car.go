package handlers

import (
	"net/http"

	"github.com/gerardva/go-api/database"
	"github.com/gerardva/go-api/models"
)

func CreateCar(input *models.Car) (err models.ErrorResponse) {
	db := database.GetDatabase()

	if result := db.Create(&input); result.Error != nil {
		err = models.NewErrorResponse(http.StatusInternalServerError, result.Error)
		return
	}

	return
}

func UpdateCar(id string, input *models.Car) (dbCar models.Car, err models.ErrorResponse) {

	if result := database.GetDatabase().First(&dbCar, id); result.Error != nil {
		err = models.NewErrorResponse(http.StatusNotFound, result.Error)
		return
	}

	dbCar.Make = input.Make
	dbCar.Model = input.Model
	dbCar.Price = input.Price

	database.GetDatabase().Save(&dbCar)

	return
}

func DeleteCar(id string) (err models.ErrorResponse) {
	var car models.Car
	if result := database.GetDatabase().First(&car, id); result.Error != nil {
		err = models.NewErrorResponse(http.StatusNotFound, result.Error)
		return
	}

	database.GetDatabase().Delete(&car)
	return
}

func GetCarById(id string) (car models.Car, err models.ErrorResponse) {
	if result := database.GetDatabase().First(&car, id); result.Error != nil {
		err = models.NewErrorResponse(http.StatusNotFound, result.Error)
		return
	}

	return
}

func GetAllCars() (cars []models.Car, err models.ErrorResponse) {
	if result := database.GetDatabase().Find(&cars); result.Error != nil {
		err = models.NewErrorResponse(http.StatusInternalServerError, result.Error)
		return
	}

	return
}