package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gerardva/go-api/models"
	"github.com/gerardva/go-api/database"
)

type CarController struct{}

func (h CarController) CreateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.BindJSON(&car); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := database.DB.Create(&car); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &car)
}

func (h CarController) UpdateCar(c *gin.Context) {
	id := c.Param("id")

	var dbCar models.Car
	if result := database.DB.First(&dbCar, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	updatedCar := models.Car{}

	// getting request's body
	if err := c.BindJSON(&updatedCar); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dbCar.Make = updatedCar.Make
	dbCar.Model = updatedCar.Model
	dbCar.Price = updatedCar.Price

	database.DB.Save(&dbCar)

	c.JSON(http.StatusOK, &dbCar)
}

func (h CarController) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	var car models.Car

	if result := database.DB.First(&car, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	database.DB.Delete(&car)
	c.Status(http.StatusOK)
}

func (h CarController) GetCarById(c *gin.Context) {
	id := c.Param("id")

	var car models.Car

	if result := database.DB.First(&car, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &car)
}

func (h CarController) GetAllCars(c *gin.Context) {
	var cars []models.Car

	if result := database.DB.Find(&cars); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &cars)
}