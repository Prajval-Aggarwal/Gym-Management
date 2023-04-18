package admin

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"

	"github.com/gin-gonic/gin"
)

func AdminRegisterService(context *gin.Context, adminRequest request.RegisterRequest) {

	var admin model.Admin
	var existRecord model.Admin

	if !db.RecordExist("admins" , "contact" , adminRequest.Contact){
		response.ErrorResponse(context, 400, "Admin already exists")
		return
	}

	admin.Name = adminRequest.Username
	admin.Contact = adminRequest.Contact
	admin.Role = "admin"

	err := db.FindById(&existRecord, adminRequest.Contact, "contact")
	if err == nil {
		response.ErrorResponse(context, 400, "Account already exists")
		return
	}
	err = db.CreateRecord(&admin)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.Response(context, 200, admin)
}
