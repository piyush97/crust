package model

import "gorm.io/gorm"

/**
 * Product Model
 * @param id int
 * @param name string
 * @param description string
 * @param price float64
 * @param createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @param updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @param deletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @return Product
 * @return error
 */
type Product struct {
	gorm.Model
	Name        string  `json:"name"`        // The name of a product.
	Quantity    int     `json:"quantity"`    // The quantity of a product.
	Description string  `json:"description"` // The description of a product.
	Price       float32 `json:"price"`       // The price of a product.

}

func (Product) TableName() string {
	return "products"
}
