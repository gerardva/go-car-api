package models

type Car struct {
	ID    uint
	Make  string `gorm:"size:255;not null;" json:"make"`
	Model string `gorm:"size:255;not null;" json:"model"`
	Price uint `gorm:"not null;" json:"price"`
}