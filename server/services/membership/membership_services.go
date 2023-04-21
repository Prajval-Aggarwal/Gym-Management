package membership

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateMembershipService(context *gin.Context, membershipData model.Membership) {
	membershipData.MemName = strings.ToLower(membershipData.MemName)
	err := db.CreateRecord(&membershipData)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Membership created successfully",
		membershipData,
		context,
	)
}

func GetMembershipsService(context *gin.Context) {

	var memberships []model.Membership

	query := "SELECT * FROM memberships ORDER BY price ASC"
	err := db.QueryExecutor(query, &memberships)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Membership retrieved",
		memberships,
		context,
	)

}

func UpdateMembershipService(context *gin.Context, updatedData model.Membership) {
	updatedData.MemName = strings.ToLower(updatedData.MemName)
	result := db.UpdateRecord(updatedData, updatedData.MemName, "mem_name")

	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Membership updated successfully",
		updatedData,
		context,
	)

}

func DeleteMembershipService(context *gin.Context, deletedData request.DeleteMembershipRequest) {

	query := "DELETE FROM memberships WHERE mem_name =?"
	err := db.QueryExecutor(query, nil, deletedData.MembershipName)

	deletedData.MembershipName = strings.ToLower(deletedData.MembershipName)
	if !db.RecordExist("memberships", "mem_name", deletedData.MembershipName) {
		response.ErrorResponse(context, 400, "Membership does not exist")
		return
	}

	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return

	}
	response.ShowResponse(
		"Success",
		200,
		"Membership deleted successfully",
		nil,
		context,
	)

}
