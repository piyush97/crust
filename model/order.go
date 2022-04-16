package model

import "gorm.io/gorm"

/**
 * Order Model
 * @param id int
 * @param productId int
 * @param userId int
 * @param quantity int
 * @param createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @param updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @param deletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @return Order
 * @return error
 */
type Order struct {
	gorm.Model

	User      User    `gorm:"foreignkey:UserID"`    // The user who placed the order.
	Product   Product `gorm:"foreignkey:ProductID"` // The product which is ordered.
	UserID    uint    `json:"user_id"`              // The user who placed the order.
	ProductID uint    `json:"product_id"`           // The product which is ordered.
	Quantity  int     `json:"quantity"`             // The quantity of the product.
}

/**
 * Get the table name of the model.
 * @return string
 */
func (Order) TableName() string {
	return "orders"
}
