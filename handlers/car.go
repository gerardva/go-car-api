package handlers

import (
	"errors"
	"github.com/gerardva/go-api/domain"
	"github.com/gerardva/go-api/domain/car"
	"net/http"
)

type CarHandler struct {
	repository car.Repository
}

func NewCarHandler(repo car.Repository) CarHandler {
	return CarHandler{repo}
}

func (h CarHandler) CreateCar(input *car.Car) (err error) {
	if err = validateCarInput(input); err != nil {
		return
	}

	return h.repository.Create(input)
}

func (h CarHandler) UpdateCar(id string, input *car.Car) (dbCar car.Car, err error) {
	if err = validateCarInput(input); err != nil {
		return
	}

	return h.repository.Update(id, input)
}

func (h CarHandler) DeleteCar(id string) (err error) {
	// Validate, check permissions, business logic etc.
	return h.repository.Delete(id)
}

func (h CarHandler) GetCarById(id string) (car car.Car, err error) {
	// Validate, check permissions, business logic etc.
	return h.repository.GetById(id)
}

func (h CarHandler) GetAllCars() (cars []car.Car, err error) {
	// Validate, check permissions, business logic etc.
	return h.repository.GetAll()
}

func validateCarInput(car *car.Car) error {
	if car.Year < 1900 {
		return domain.NewErrorResponse(http.StatusBadRequest, errors.New("invalid year"))
	}

	return nil
}
