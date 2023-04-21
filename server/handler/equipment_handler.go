package handler

import (
	"gym/server/model"
	"gym/server/response"
	"gym/server/services/equipment"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

//	@Description	Adds a new equipment to the gym
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
// Param EquipDetails body model.Equipment true "Equipment details"
//
//	@Tags			Equipment
//	@Router			/createEquipment [post]
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

//	@Description	Get List of equipments
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//	@Tags			Equipment
//	@Router			/getEquipments [get]
func GetEquipmentHandler(context *gin.Context) {
	utils.SetHeader(context)
	equipment.GetEquipmentService(context)
}
