package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/admin"
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
	admin.AdminRegisterService(context, adminRequest)
}
