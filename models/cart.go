package models

import "gorm.io/gorm"

type CartBook struct {
	gorm.Model
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	Book Book `gorm:"foreignKey:BookID"`
	Qty int `json:"qty"`
}