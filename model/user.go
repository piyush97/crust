package model

import "gorm.io/gorm"

/**
 * User Model
 * @param id int
 * @param name string
 * @param email string
 * @param password string
 * @param role string
 * @param createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @param updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @param deletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
 * @return User
 * @return error
 */
type User struct {
	gorm.Model

	Name     string  `json:"name"`     // The name of a user.
	Email    *string `json:"email"`    // The email of a user.
	Password string  `json:"password"` // The password of a user.
	Role     string  `json:"role"`     // The role of a user.

}

func (User) TableName() string {
	return "users"
}
