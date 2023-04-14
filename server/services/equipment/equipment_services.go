package equipment

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/response"

	"github.com/gin-gonic/gin"
)

func CreateEquipmentService(context *gin.Context, equipData model.Equipment) {
	result := db.UpdateRecord(&equipData, equipData.Equip_Name, "equip_name")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		err := db.CreateRecord(&equipData)
		if err != nil {
			response.ErrorResponse(context, 400, err.Error())
			return
		}
		response.ShowResponse(
			"Success",
			200,
			"New Equipment added",
			equipData,
			context,
		)

	} else {
		response.ShowResponse(
			"Success",
			200,
			"Old Equipment updated successfully",
			equipData,
			context,
		)
	}
}

func GetEquipmentService(context *gin.Context) {
	var equipments []model.Equipment
	query := "SELECT * FROM equipment ORDER BY equip_name ASC"
	err := db.QueryExecutor(query, &equipments)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Equipments retrieved",
		equipments,
		context,
	)
}
