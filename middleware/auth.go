package middleware

import (
	"../structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {

	//var jwtStruct JWTStruct

	response := structs.JsonResponseToken{}

	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("decandTampan"), nil
	})

	if token != nil && err == nil {
		//fmt.Println("token verified")
	} else {
		response.ApiStatus = 0
		response.ApiMessage = "not authorized " + err.Error()
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
	}

}
