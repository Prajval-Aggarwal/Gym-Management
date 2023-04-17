package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/user"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)



func CreateUserHandler(context *gin.Context) {

	utils.SetHeader(context)

	var createUserRequest request.CreateUserRequest

	utils.RequestDecoding(context, &createUserRequest)

	err := validation.CheckValidation(&createUserRequest)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	user.CreateUserService(context, createUserRequest)
}

func GetUserByIdHandler(context *gin.Context) {

	utils.SetHeader(context)

	var userId request.UserRequest

	utils.RequestDecoding(context, &userId)

	err := validation.CheckValidation(&userId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	user.GetUserByIdService(context, userId)

}

func UserAttendenceHandler(context *gin.Context) {

	utils.SetHeader(context)

	var userId request.UserRequest

	utils.RequestDecoding(context, &userId)

	err := validation.CheckValidation(&userId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	user.UserAttendenceService(context, userId)
}

func GetAllUsers(context *gin.Context){
	utils.SetHeader(context)
	user.GetAllUserServices(context)
}
