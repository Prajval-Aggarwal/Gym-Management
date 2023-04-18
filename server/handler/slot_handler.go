package handler

import (
	"gym/server/request"
	"gym/server/response"
	slot "gym/server/services/slots"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

//	@Description	updates sthe slot for the user
//	@Accept			json
//	@Produce		json
//
//	@Success		200		{object}	response.Success
//	@Failure		400		{object}	response.Error
//
//	@Param			UserId	body		request.UpdateSlotRequest	true	"Slot number"
//	@Tags			Slot
//	@Router			/slotUpdate [put]
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
