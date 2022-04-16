package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

/**
 * Generate JWT Token
 * @param userID int
 * @return string
 * @return error
 */
func GenerateToken(userid uint) string {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 3).Unix(), // 3 hours
		"iat":    time.Now().Unix(),                    // time of token creation
		"userID": userid,                               // user id
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)  // create new token
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET"))) // sign token and return string
	return t                                                    // return token string

}

/**
 * Validate JWT Token
 * @param token string
 *
 */
func ValidateToken(token string) (*jwt.Token, error) { // return token object and error

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { // parse token and return token object and error
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // if token is not signed with HMAC
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // return error
		}
		return []byte(os.Getenv("JWT_SECRET")), nil // return secret key
	})
}
