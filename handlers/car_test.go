package handlers

import (
	"github.com/gerardva/go-api/database/repository"
	"github.com/gerardva/go-api/domain/car"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateValid(t *testing.T) {
	assert := assert.New(t)

	validCar := car.NewCar(0, "BMW", "M3", 80000, 2023)
	repo := new(repository.MockCarRepository)
	repo.On("Create", &validCar).Return(nil)

	handler := NewCarHandler(repo)
	err := handler.CreateCar(&validCar)

	assert.Nil(err)
}

func TestCreateInvalidYear(t *testing.T) {
	assert := assert.New(t)

	invalidCar := car.NewCar(0, "BMW", "M3", 80000, 1800)
	repo := new(repository.MockCarRepository)
	handler := NewCarHandler(repo)
	err := handler.CreateCar(&invalidCar)

	assert.NotNil(err)
}
