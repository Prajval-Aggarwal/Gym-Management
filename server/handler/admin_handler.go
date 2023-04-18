package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/admin"
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
	admin.AdminRegisterService(context, adminRequest)
}
