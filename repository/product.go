package repository

import (
	"github.com/piyush97/crust/model"
	"gorm.io/gorm"
)

//ProductRepository --> Interface to ProductRepository
type ProductRepository interface {
	Getproduct(int) (model.Product, error)              //GetProduct --> returns product by id
	GetAllproduct() ([]model.Product, error)            //GetAllProduct --> returns all products
	AddProduct(model.Product) (model.Product, error)    //AddProduct --> adds new product
	UpdateProduct(model.Product) (model.Product, error) //UpdateProduct --> updates product
	DeleteProduct(model.Product) (model.Product, error) //DeleteProduct --> deletes product
}

type productRepository struct {
	connection *gorm.DB //connection to database
}

//NewProductRepository --> returns new product repository
func NewProductRepository() ProductRepository {
	return &productRepository{ //connection to database
		connection: DB(),
	}
}

/**
 * Get Product by id
 * @param id int
 * @return Product
 */
func (db *productRepository) Getproduct(id int) (product model.Product, err error) {
	return product, db.connection.First(&product, id).Error //return product by id
}

/**
 * Get All Products
 * @return []Product
 * @return error
 */
func (db *productRepository) GetAllproduct() (products []model.Product, err error) {
	return products, db.connection.Find(&products).Error //return all products
}

/**
 * Add Product
 * @param product Product
 * @return Product
 */
func (db *productRepository) AddProduct(product model.Product) (model.Product, error) {
	return product, db.connection.Create(&product).Error //return product
}

/**
 * Update Product
 * @param product Product
 * @return Product
 */
func (db *productRepository) UpdateProduct(product model.Product) (model.Product, error) {
	if err := db.connection.First(&product, product.ID).Error; err != nil {
		return product, err //return product
	}
	return product, db.connection.Model(&product).Updates(&product).Error //return product
}

/**
 * Delete Product
 * @param product Product
 * @return Product
 */
func (db *productRepository) DeleteProduct(product model.Product) (model.Product, error) {
	if err := db.connection.First(&product, product.ID).Error; err != nil { //return product
		return product, err //return product
	}
	return product, db.connection.Delete(&product).Error //return product
}
