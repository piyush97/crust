package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/piyush97/crust/handler"
)

/**
 * Authorize JWT Token
 */
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) { //return gin.HandlerFunc
		const BearerSchema string = "Bearer "        //Bearer schema
		authHeader := ctx.GetHeader("Authorization") //get authorization header
		if authHeader == "" {                        //if authorization header is empty
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ //return unauthorized
				"error": "No Authorization header found"}) //return error

		}
		tokenString := authHeader[len(BearerSchema):] //get token string

		if token, err := handler.ValidateToken(tokenString); err != nil { //validate token

			fmt.Println("token", tokenString, err.Error())          //print token and error
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ //return unauthorized
				"error": "Not Valid Token"}) //return error

		} else {

			if claims, ok := token.Claims.(jwt.MapClaims); !ok { //if token is not valid
				ctx.AbortWithStatus(http.StatusUnauthorized) //return unauthorized

			} else {
				if token.Valid { //if token is valid
					ctx.Set("userID", claims["userID"])                   //set userID to context
					fmt.Println("during authorization", claims["userID"]) //print userID
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized) //return unauthorized
				}

			}
		}

	}

}
