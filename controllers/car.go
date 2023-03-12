package controllers

import (
	"github.com/gerardva/go-api/domain"
	"github.com/gerardva/go-api/domain/car"
	"net/http"

	"github.com/gerardva/go-api/handlers"
	"github.com/gin-gonic/gin"
)

type CarController struct {
	handler handlers.CarHandler
}

func NewCarController(handler handlers.CarHandler) CarController {
	return CarController{handler}
}

func (h CarController) CreateCar(c *gin.Context) {
	car := car.Car{}

	if err := c.BindJSON(&car); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.handler.CreateCar(&car); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, &car)
}

func (h CarController) UpdateCar(c *gin.Context) {
	id := c.Param("id")
	car := car.Car{}

	if err := c.BindJSON(&car); err != nil {
		handleError(c, err)
		return
	}

	updatedCar, err := h.handler.UpdateCar(id, &car)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &updatedCar)
}

func (h CarController) DeleteCar(c *gin.Context) {
	id := c.Param("id")
	if err := h.handler.DeleteCar(id); err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h CarController) GetCarById(c *gin.Context) {
	id := c.Param("id")

	car, err := h.handler.GetCarById(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &car)
}

func (h CarController) GetAllCars(c *gin.Context) {
	cars, err := h.handler.GetAllCars()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, &cars)
}

func handleError(c *gin.Context, err error) {
	if errResponse, ok := err.(*domain.ErrorResponse); ok {
		c.AbortWithError(errResponse.StatusCode, err)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
