package controllers

import (
	"net/http"

	"github.com/gerardva/go-api/handlers"
	"github.com/gerardva/go-api/models"
	"github.com/gin-gonic/gin"
)

type CarController struct{

}

func (h CarController) CreateCar(c *gin.Context) {
	car := models.Car{}

	if err := c.BindJSON(&car); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := handlers.CreateCar(&car); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &car)
}

func (h CarController) UpdateCar(c *gin.Context) {
	id := c.Param("id")
	car := models.Car{}

	if err := c.BindJSON(&car); err != nil {
		handleError(c, err)
		return
	}
	
	updatedCar, err := handlers.UpdateCar(id, &car)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &updatedCar)
}

func (h CarController) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if err := handlers.DeleteCar(id); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h CarController) GetCarById(c *gin.Context) {
	id := c.Param("id")

	car, err := handlers.GetCarById(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &car)
}

func (h CarController) GetAllCars(c *gin.Context) {
	cars, err := handlers.GetAllCars()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &cars)
}

func handleError(c *gin.Context, err error) {
	if errResponse, ok  := err.(*models.ErrorResponse); ok {		
		c.AbortWithError(errResponse.StatusCode, err)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}