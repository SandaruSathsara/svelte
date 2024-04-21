package models

import (
	"gorm.io/gorm"
)

type Sneaker struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Size        int     `json:"size"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
}

// CreateSneaker
func CreateWomenSneaker(db *gorm.DB, sneaker *Sneaker) error {
	return db.Create(sneaker).Error
}

// GetSneakers
func GetWomenSneakers(db *gorm.DB) ([]Sneaker, error) {
	var sneakers []Sneaker
	err := db.Find(&sneakers).Error
	return sneakers, err
}

// GetSneakerByID
func GetWomenSneakerByID(db *gorm.DB, id uint) (*Sneaker, error) {
	var sneaker Sneaker
	err := db.First(&sneaker, id).Error
	return &sneaker, err
}

// UpdateSneaker
func UpdateWomenSneaker(db *gorm.DB, sneaker *Sneaker) error {
	return db.Save(sneaker).Error
}

// DeleteSneaker
func DeleteWomenSneaker(db *gorm.DB, id uint) error {
	var sneaker Sneaker
	if err := db.First(&sneaker, id).Error; err != nil {
		return err
	}
	return db.Delete(&sneaker).Error
}
