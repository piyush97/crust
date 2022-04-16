package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piyush97/crust/model"
	"github.com/piyush97/crust/repository"
)

//ProductHandler --> interface to Product handler
type ProductHandler interface {
	GetProduct(*gin.Context)    //GetProduct --> returns product by id
	GetAllProduct(*gin.Context) //GetAllProduct --> returns all products
	AddProduct(*gin.Context)    //AddProduct --> adds new product
	UpdateProduct(*gin.Context) //UpdateProduct --> updates product
	DeleteProduct(*gin.Context) //DeleteProduct --> deletes product
}

type productHandler struct {
	repo repository.ProductRepository //repo for product
}

//NewProductHandler --> returns new handler for product entity
func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

/**
 * Get All Products
 * @return []Product
 * @return error
 */
func (h *productHandler) GetAllProduct(ctx *gin.Context) {
	product, err := h.repo.GetAllproduct() //GetAllproduct --> returns all products
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //InternalServerError --> returns error
		return

	}
	ctx.JSON(http.StatusOK, product) //OK --> returns product

}

/**
 * Get Product by id
 * @param id int
 * @return Product
 * @return error
 */
func (h *productHandler) GetProduct(ctx *gin.Context) {
	prodStr := ctx.Param("product")      //GetParam --> returns param
	prodID, err := strconv.Atoi(prodStr) //Atoi --> converts string to int
	if err != nil {                      //if error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //BadRequest --> returns error
		return
	}
	product, err := h.repo.Getproduct(prodID) //Getproduct --> returns product by id
	if err != nil {                           //if error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //InternalServerError --> returns error
		return

	}
	ctx.JSON(http.StatusOK, product) //OK --> returns product

}

/**
 *  Add Product
 * @param name string
 * @param quantity int
 * @param description string
 */
func (h *productHandler) AddProduct(ctx *gin.Context) {
	var product model.Product                            //Product --> model
	if err := ctx.ShouldBindJSON(&product); err != nil { //ShouldBindJSON --> binds json to model
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //BadRequest --> returns error
		return
	}
	product, err := h.repo.AddProduct(product) //AddProduct --> adds new product
	if err != nil {                            //if error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //InternalServerError --> returns error
		return

	}
	ctx.JSON(http.StatusOK, product) //OK --> returns product

}

/**
 * Update Product
 * @param id int
 * @param name string
 * @param quantity int
 * @param description string
 * @return Product
 */
func (h *productHandler) UpdateProduct(ctx *gin.Context) {

	var product model.Product                            //Product --> model
	if err := ctx.ShouldBindJSON(&product); err != nil { //ShouldBindJSON --> binds json to model
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //BadRequest --> returns error
		return
	}
	prodStr := ctx.Param("product")      //GetParam --> returns param
	prodID, err := strconv.Atoi(prodStr) //Atoi --> converts string to int
	if err != nil {                      //if error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //BadRequest --> returns error
		return
	}
	product.ID = uint(prodID)                    //ID --> model
	product, err = h.repo.UpdateProduct(product) //UpdateProduct --> updates product
	if err != nil {                              //if error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //InternalServerError --> returns error
		return

	}
	ctx.JSON(http.StatusOK, product) //OK --> returns product

}

/**
 * Delete Product
 * @param id int
 * @return Product
 * @return error
 */
func (h *productHandler) DeleteProduct(ctx *gin.Context) { //DeleteProduct --> deletes product

	var product model.Product                     //Product --> model
	prodStr := ctx.Param("product")               //GetParam --> returns param
	prodID, _ := strconv.Atoi(prodStr)            //Atoi --> converts string to int
	product.ID = uint(prodID)                     //ID --> model
	product, err := h.repo.DeleteProduct(product) //DeleteProduct --> deletes product
	if err != nil {                               //if error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //InternalServerError --> returns error
		return

	}
	ctx.JSON(http.StatusOK, product) //OK --> returns product

}
