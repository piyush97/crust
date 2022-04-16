package repository

import (
	"github.com/piyush97/crust/model"
	"gorm.io/gorm"
)

/**
 * User Repository Interface
 *
 * @return UserRepository
 */
type UserRepository interface {
	AddUser(model.User) (model.User, error)
	GetUser(int) (model.User, error)
	GetByEmail(string) (model.User, error)
	GetAllUser() ([]model.User, error)
	UpdateUser(model.User) (model.User, error)
	DeleteUser(model.User) (model.User, error)
	GetProductOrdered(int) ([]model.Order, error)
}

type userRepository struct {
	connection *gorm.DB
}

//NewUserRepository --> returns new user repository
func NewUserRepository() UserRepository {
	return &userRepository{
		connection: DB(),
	}
}

/**
 * Get User
 * @param id int
 * @return User
 * @return error
 */
func (db *userRepository) GetUser(id int) (user model.User, err error) {
	return user, db.connection.First(&user, id).Error
}

/**
 * Get User By Email
 * @param email string
 * @return User
 * @return error
 */
func (db *userRepository) GetByEmail(email string) (user model.User, err error) {
	return user, db.connection.First(&user, "email=?", email).Error
}

/**
 * Get All User
 * @return []User
 * @return error
 */
func (db *userRepository) GetAllUser() (users []model.User, err error) {
	return users, db.connection.Find(&users).Error
}

/**
 * Add User
 * @param user User
 * @return User
 * @return error
 */
func (db *userRepository) AddUser(user model.User) (model.User, error) {
	return user, db.connection.Create(&user).Error
}

/**
 * Update User
 * @param user User
 * @return User
 * @return error
 */
func (db *userRepository) UpdateUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Model(&user).Updates(&user).Error
}

/**
 * Delete User
 * @param user User
 * @return User
 * @return error
 */
func (db *userRepository) DeleteUser(user model.User) (model.User, error) {
	if err := db.connection.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, db.connection.Delete(&user).Error
}

/**
 * Get Product Ordered
 * @param id int
 * @return []Order
 * @return error
 */
func (db *userRepository) GetProductOrdered(userID int) (orders []model.Order, err error) {
	return orders, db.connection.Where("user_id = ?", userID).Set("gorm:auto_preload", true).Find(&orders).Error
}
