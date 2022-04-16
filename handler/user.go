package handler

import (
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/piyush97/crust/model"
	"github.com/piyush97/crust/repository"
	"golang.org/x/crypto/bcrypt"
)

//UserHandler -> interface to User entity
type UserHandler interface {
	AddUser(*gin.Context)           //AddUser
	GetUser(*gin.Context)           //GetUser
	GetAllUser(*gin.Context)        //GetAllUser
	SignInUser(*gin.Context)        //SignInUser
	UpdateUser(*gin.Context)        //UpdateUser
	DeleteUser(*gin.Context)        //DeleteUser
	GetProductOrdered(*gin.Context) //GetProductOrdered
}

type userHandler struct {
	repo repository.UserRepository //repository
}

//NewUserHandler --> returns new user handler
func NewUserHandler() UserHandler {

	return &userHandler{
		repo: repository.NewUserRepository(), //repository
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)                                             //convert string to byte
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost) //hash password
	*pass = string(hPass)                                                 //convert byte to string
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil //compare password
}

func (h *userHandler) GetAllUser(ctx *gin.Context) {
	fmt.Println(ctx.Get("userID"))   //get userID from context
	user, err := h.repo.GetAllUser() //get all user
	if err != nil {                  //if error
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //return error
		return                                                                //return

	}
	ctx.JSON(http.StatusOK, user) //return user

}

func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")          //get id from url
	intID, err := strconv.Atoi(id) //convert string to int
	if err != nil {                //if error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //return error
		return
	}
	user, err := h.repo.GetUser(intID) //get user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //return error
		return

	}
	ctx.JSON(http.StatusOK, user) //return user

}

func (h *userHandler) SignInUser(ctx *gin.Context) {
	var user model.User                               //user
	if err := ctx.ShouldBindJSON(&user); err != nil { //bind json
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //return error
	}

	dbUser, err := h.repo.GetByEmail(user.Email) //get user by email
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No Such User Found"}) //return error
		return

	}
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)                                        //get userID
		token := GenerateToken(dbUser.ID)                                            //generate token
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignIN", "token": token}) //return token
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"}) //return error
	return

}

func (h *userHandler) AddUser(ctx *gin.Context) {
	var user model.User                               //user
	if err := ctx.ShouldBindJSON(&user); err != nil { //bind json
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //return error
		return
	}
	hashPassword(&user.Password)      //hash password
	user, err := h.repo.AddUser(user) //add user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //return error
		return

	}
	user.Password = ""            //remove password
	ctx.JSON(http.StatusOK, user) //return user

}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user model.User                               //user
	if err := ctx.ShouldBindJSON(&user); err != nil { //bind json
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //return error
		return
	}
	id := ctx.Param("user")        //get id from url
	intID, err := strconv.Atoi(id) //convert string to int
	if err != nil {                //if error
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //return error
	}
	user.ID = uint(intID)               //set id
	user, err = h.repo.UpdateUser(user) //update user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //return error
		return

	}
	ctx.JSON(http.StatusOK, user) //return user

}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user model.User                  //user
	id := ctx.Param("user")              //get id from url
	intID, _ := strconv.Atoi(id)         //convert string to int
	user.ID = uint(intID)                //set id
	user, err := h.repo.DeleteUser(user) //delete user
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //return error
		return

	}
	ctx.JSON(http.StatusOK, user) //return user

}

func (h *userHandler) GetProductOrdered(ctx *gin.Context) {

	userStr := ctx.Param("user")                                       //get user from url
	userID, _ := strconv.Atoi(userStr)                                 //convert string to int
	if products, err := h.repo.GetProductOrdered(userID); err != nil { //get product ordered
		ctx.JSON(http.StatusBadRequest, gin.H{ //return error
			"error": err.Error(), //error
		})
	} else {
		ctx.JSON(http.StatusOK, products) //return product
	}
}
