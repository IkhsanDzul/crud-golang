package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
}