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

// @Description	Registers a admin
// @Accept			json
// @Produce		json
// @Success		200				{object}	response.Success
// @Failure		400				{object}	response.Error
// @Param			AdminDetails	body		request.RegisterRequest	true	"Registers a admin"
// @Tags			Authentication
// @Router			/adminRegister [post]
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

// @Description	Sends a otp to the number entered
// @Accept			json
// @Produce		json
// @Success		200			{object}	response.Success
// @Failure		400			{object}	response.Error
// @Param			phoneNumber	body		request.SendOtpRequest	true	"Phone Number of registered user"
// @Tags			Authentication
// @Router			/sendOTP [post]
func SendOtpHandler(context *gin.Context) {
	utils.SetHeader(context)
	var phoneNumber request.SendOtpRequest
	utils.RequestDecoding(context, &phoneNumber)
	fmt.Println("phoneNumber is", phoneNumber)
	authentication.SendOtpService(context, phoneNumber)
}

// @Description	Verifies the OTP sent to the user
// @Accept			json
// @Produce		json
// @Success		200		{object}	response.Success
// @Failure		401		string		"Verification Failed"
// @Param			details	body		request.VerifyOtpRequest	true	"Phone Number of registered user and the otp sent to it"
// @Tags			Authentication
// @Router			/verifyOTP [post]
func VerifyOtpHandler(context *gin.Context) {
	utils.SetHeader(context)
	var verifyRequest request.VerifyOtpRequest
	utils.RequestDecoding(context, &verifyRequest)

	authentication.VerifyOtpService(context, verifyRequest)

}

// @Description	Logs out a user
// @Accept			json
// @Produce		json
// @Success		200	{object}	response.Success
// @Failure		400	{object}	response.Error
// @Tags			Authentication
// @Router			/logOut [get]
func LogoutHandler(context *gin.Context) {
	cookie, err := context.Request.Cookie("cookie")
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	authentication.LogoutService(context, cookie.Value)

}
