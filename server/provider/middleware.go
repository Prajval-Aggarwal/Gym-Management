package provider

import (
	"fmt"
	"gym/server/response"

	"github.com/gin-gonic/gin"
)

func UserAuthorization(c *gin.Context)  {
	
		fmt.Println("inside middleware")
		token, err := c.Request.Cookie("cookie")
		if err != nil {
			response.ErrorResponse(c, 400, err.Error())
			return
		}
		
		claims , err := DecodeToken(token.Value)
		if err!=nil{
			response.ErrorResponse(c , 401 , err.Error())
		}
		err = claims.Valid()
		if err!=nil{
			response.ErrorResponse(c , 401 , err.Error())
		}
		if claims.Role == "user"{
			c.Next()
		}else{
			response.ErrorResponse(c , 403 , "Access Denied")
			c.Abort()
			return
		}
		

}

func AdminAuthorization(c *gin.Context)  {
	
	fmt.Println("inside middleware")
	token, err := c.Request.Cookie("cookie")
	if err != nil {
		response.ErrorResponse(c, 400, err.Error())
		return
	}
	
	claims , err := DecodeToken(token.Value)
	if err!=nil{
		response.ErrorResponse(c , 401 , err.Error())
	}
	err = claims.Valid()
	if err!=nil{
		response.ErrorResponse(c , 401 , err.Error())
	}
	if claims.Role == "admin"{
		c.Next()
	}else{
		response.ErrorResponse(c , 403 , "Access Denied")
		c.Abort()
		return
	}
	

}

