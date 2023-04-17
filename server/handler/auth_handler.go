package handler

import (
	"fmt"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/authentication"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

func AdminRegisterHandler(context *gin.Context) {
	utils.SetHeader(context)
	var adminRequest request.RegisterRequest
	utils.RequestDecoding(context, &adminRequest)
	err := validation.CheckValidation(&adminRequest)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	authentication.AdminRegisterService(context, adminRequest)
}

func SendOtpHandler(context *gin.Context) {
	utils.SetHeader(context)
	var phoneNumber request.SendOtpRequest
	utils.RequestDecoding(context, &phoneNumber)
	fmt.Println("phoneNumber is", phoneNumber)
	authentication.SendOtpService(context, phoneNumber)

}

func VerifyOtpHandler(context *gin.Context) {
	utils.SetHeader(context)
	var verifyRequest request.VerifyOtpRequest
	utils.RequestDecoding(context, &verifyRequest)

	authentication.VerifyOtpService(context, verifyRequest)

}

func UserRegisterHandler(context *gin.Context) {
	utils.SetHeader(context)
	var userRequest request.RegisterRequest
	utils.RequestDecoding(context, &userRequest)
	err := validation.CheckValidation(&userRequest)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	authentication.UserRegisterService(context, userRequest)
}

func LogoutHandler(context *gin.Context) {
	cookie, err := context.Request.Cookie("token")
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	tokenString := cookie.Value
	fmt.Println("token in the cookies is", tokenString)

	//create its service blacklist the token and expire the cookie

}
