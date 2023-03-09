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

	if resp := handlers.CreateCar(&car); resp.Error != nil {
		c.AbortWithError(resp.StatusCode, resp.Error)
		return
	}

	c.JSON(http.StatusCreated, &car)
}

func (h CarController) UpdateCar(c *gin.Context) {
	id := c.Param("id")
	car := models.Car{}

	if err := c.BindJSON(&car); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	updatedCar, resp := handlers.UpdateCar(id, &car)
	if resp.Error != nil {
		c.AbortWithError(resp.StatusCode, resp.Error)
		return
	}

	c.JSON(http.StatusOK, &updatedCar)
}

func (h CarController) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if resp := handlers.DeleteCar(id); resp.Error != nil {
		c.AbortWithError(resp.StatusCode, resp.Error)
		return
	}

	c.Status(http.StatusOK)
}

func (h CarController) GetCarById(c *gin.Context) {
	id := c.Param("id")

	car, resp := handlers.GetCarById(id)
	if resp.Error != nil {
		c.AbortWithError(resp.StatusCode, resp.Error)
		return
	}

	c.JSON(http.StatusOK, &car)
}

func (h CarController) GetAllCars(c *gin.Context) {
	cars, resp := handlers.GetAllCars()
	if resp.Error != nil {
		c.AbortWithError(resp.StatusCode, resp.Error)
		return
	}

	c.JSON(http.StatusOK, &cars)
}