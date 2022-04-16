package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/piyush97/crust/handler"
	"github.com/piyush97/crust/middleware"
)

func RunAPI(address string) error {

	userHandler := handler.NewUserHandler()       //create new user handler
	productHandler := handler.NewProductHandler() //create new product handler

	r := gin.Default() //create new gin router

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Crust API") //return welcome message
	})

	apiRoutes := r.Group("/api")           //create new api group
	userRoutes := apiRoutes.Group("/user") //create new user group

	{
		userRoutes.POST("/register", userHandler.AddUser)  //add user
		userRoutes.POST("/signin", userHandler.SignInUser) //sign in user
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT()) //create new user protected group
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)                      //get all user
		userProtectedRoutes.GET("/:user", userHandler.GetUser)                    //get user
		userProtectedRoutes.GET("/:user/products", userHandler.GetProductOrdered) //get product ordered
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)                 //update user
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)              //delete user
	}

	productProtectedRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT()) //create new product protected group
	{
		productProtectedRoutes.GET("/", productHandler.GetAllProduct)            //get all product
		productProtectedRoutes.GET("/:product", productHandler.GetProduct)       //get product
		productProtectedRoutes.POST("/", productHandler.AddProduct)              //add product
		productProtectedRoutes.PUT("/:product", productHandler.UpdateProduct)    //update product
		productProtectedRoutes.DELETE("/:product", productHandler.DeleteProduct) //delete product
	}

	return r.Run(address) //run server

}
