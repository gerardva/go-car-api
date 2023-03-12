package repository

import (
	"github.com/gerardva/go-api/domain"
	"github.com/gerardva/go-api/domain/car"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
)

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return CarRepository{db}
}

func (r CarRepository) Create(input *car.Car) (err error) {
	if result := r.db.Create(&input); result.Error != nil {
		err = domain.NewErrorResponse(http.StatusInternalServerError, result.Error)
		return
	}

	return
}

func (r CarRepository) Update(id string, input *car.Car) (dbCar car.Car, err error) {

	if result := r.db.First(&dbCar, id); result.Error != nil {
		err = domain.NewErrorResponse(http.StatusNotFound, result.Error)
		return
	}

	dbCar.Make = input.Make
	dbCar.Model = input.Model
	dbCar.Price = input.Price

	r.db.Save(&dbCar)

	return
}

func (r CarRepository) Delete(id string) (err error) {
	var car car.Car
	if result := r.db.First(&car, id); result.Error != nil {
		err = domain.NewErrorResponse(http.StatusNotFound, result.Error)
		return
	}

	r.db.Delete(&car)
	return
}

func (r CarRepository) GetById(id string) (car car.Car, err error) {
	if result := r.db.First(&car, id); result.Error != nil {
		err = domain.NewErrorResponse(http.StatusNotFound, result.Error)
		return
	}

	return
}

func (r CarRepository) GetAll() (cars []car.Car, err error) {
	if result := r.db.Find(&cars); result.Error != nil {
		err = domain.NewErrorResponse(http.StatusInternalServerError, result.Error)
		return
	}

	return
}

// Mock

type MockCarRepository struct {
	mock.Mock
}

func (r MockCarRepository) Create(input *car.Car) (err error) {
	args := r.Called(input)
	return args.Error(0)
}

func (r MockCarRepository) Update(id string, input *car.Car) (dbCar car.Car, err error) {
	args := r.Called(id, input)
	return args.Get(0).(car.Car), args.Error(1)
}

func (r MockCarRepository) Delete(id string) (err error) {
	args := r.Called(id)
	return args.Error(0)
}

func (r MockCarRepository) GetById(id string) (dbCar car.Car, err error) {
	args := r.Called(id)
	return args.Get(0).(car.Car), args.Error(1)
}

func (r MockCarRepository) GetAll() (cars []car.Car, err error) {
	args := r.Called()
	return args.Get(0).([]car.Car), args.Error(1)
}
