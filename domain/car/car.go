package car

type Car struct {
	ID    uint
	Make  string `gorm:"size:255;not null;" json:"make"`
	Model string `gorm:"size:255;not null;" json:"model"`
	Price uint   `gorm:"not null;" json:"price"`
	Year  uint   `gorm:"not null;default:2023" json:"year"`
}

func NewCar(id uint, make string, model string, price uint, year uint) Car {
	return Car{id, make, model, price, year}
}

type Repository interface {
	Create(input *Car) (err error)
	Update(id string, input *Car) (dbCar Car, err error)
	Delete(id string) (err error)
	GetById(id string) (car Car, err error)
	GetAll() (cars []Car, err error)
}

type Handler interface {
	Create(input *Car) (err error)
	Update(id string, input *Car) (dbCar Car, err error)
	Delete(id string) (err error)
	GetById(id string) (car Car, err error)
	GetAll() (cars []Car, err error)
}
