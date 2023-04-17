package provider

import (
	"fmt"
	"gym/server/response"
	"gym/server/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

//Generate JWT Token
func GenerateToken(id string, role string, context *gin.Context) string {
	claims := Claims{
		Id:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(utils.TOKENEXPIRATIONDURATION),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))

	if err != nil {
		response.ErrorResponse(context, 401, "Error signing token")
	}
	return tokenString
}


//Decode Token function
func DecodeToken(tokenString string) (Claims, error) {
	claims := &Claims{}

	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil || !parsedToken.Valid {
		return *claims, fmt.Errorf("invalid or expired token")
	}

	return *claims, nil
}

//Set cookie handler
func SetCookie(context *gin.Context, tokenString string) {

	context.SetCookie(
		"cookie",
		tokenString,
		7200,
		"",
		"",
		true,
		true,
	)
	
	response.ShowResponse(
		"Success",
		200,
		"Cookies saved successfully",
		"",
		context,
	)
}

//Delete cookie handler
func DeleteCookie(context *gin.Context) {
	context.SetCookie(
        "cookie",
        "",
        -1, 
        "",
        "",
        false,
        true,
    )

    response.ShowResponse(
        "Success",
        200,
        "Cookie deleted successfully",
        "",
        context,
    )
}
