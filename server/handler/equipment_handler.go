package handler

import (
	"gym/server/model"
	"gym/server/response"
	"gym/server/services/equipment"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

func CreateEquipmentHandler(context *gin.Context) {

	utils.SetHeader(context)

	var createEquipment model.Equipment

	utils.RequestDecoding(context, &createEquipment)

	err := validation.CheckValidation(&createEquipment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	equipment.CreateEquipmentService(context, createEquipment)

}
func GetEquipmentHandler(context *gin.Context) {
	utils.SetHeader(context)
	equipment.GetEquipmentService(context)
}
