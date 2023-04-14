package handler

import (
	"gym/server/request"
	"gym/server/response"
	slot "gym/server/services/slots"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

func SlotUpdateHandler(context *gin.Context) {
	utils.SetHeader(context)
	var userId request.UpdateSlotRequest

	utils.RequestDecoding(context, &userId)

	err := validation.CheckValidation(&userId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	slot.SlotServices(context, userId)

}
