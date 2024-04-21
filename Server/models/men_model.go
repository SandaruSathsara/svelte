package models

import (
	"gorm.io/gorm"
)

type MenSneaker struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Name  string  `json:"name"`
	Size  int     `json:"size"`
	Price float64 `json:"price"`
}

// CreateMenSneaker
func CreateMenSneaker(db *gorm.DB, menSneaker *MenSneaker) error {
	return db.Create(menSneaker).Error
}

// GetMenSneakers
func GetMenSneakers(db *gorm.DB) ([]MenSneaker, error) {
	var menSneakers []MenSneaker
	err := db.Find(&menSneakers).Error
	return menSneakers, err
}

// GetMenSneakerByID
func GetMenSneakerByID(db *gorm.DB, id uint) (*MenSneaker, error) {
	var menSneaker MenSneaker
	err := db.First(&menSneaker, id).Error
	return &menSneaker, err
}

// UpdateMenSneaker
func UpdateMenSneaker(db *gorm.DB, menSneaker *MenSneaker) error {
	return db.Save(menSneaker).Error
}

// DeleteMenSneaker
func DeleteMenSneaker(db *gorm.DB, id uint) error {
	var menSneaker MenSneaker
	if err := db.First(&menSneaker, id).Error; err != nil {
		return err
	}
	return db.Delete(&menSneaker).Error
}
