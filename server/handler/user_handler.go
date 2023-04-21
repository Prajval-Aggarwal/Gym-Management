package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/user"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

//	@Description	Creates a new user record in database
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
// Param EmpDetails body request.CreateSubRequest true "User details"
//
//	@Tags			User
//	@Router			/createUser [post]
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

//	@Description	Gets a singler User
//	@Accept			json
//	@Produce		json
//
//	@Success		200		{object}	response.Success
//	@Failure		400		{object}	response.Error
//
//	@Param			UserId	body		request.UserRequest	true	"User type like tranier,cleaner"
//	@Tags			User
//	@Router			/getUserById [post]
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

//	@Description	Marks the User present for that day
//	@Accept			json
//	@Produce		json
//
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//
//	@Param			UserDetails	body		request.UserRequest	true	"Details of User whose attendence is to be marked"
//	@Tags			User
//	@Router			/userAttendence [post]
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

//	@Description	Gets the list of users
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//
//	@Tags			Employee
//	@Router			/getUsers [get]
func GetAllUsers(context *gin.Context) {
	utils.SetHeader(context)
	user.GetAllUserServices(context)
}
