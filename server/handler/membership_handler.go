package handler

import (
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/membership"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

func CreateMembershipHandler(context *gin.Context) {

	utils.SetHeader(context)

	var createMembership model.Membership

	utils.RequestDecoding(context, &createMembership)

	err := validation.CheckValidation(&createMembership)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	membership.CreateMembershipService(context, createMembership)
}

func GetMembershipHandler(context *gin.Context) {
	utils.SetHeader(context)
	membership.GetMembershipsService(context)
}

func UpdateMembershipHandler(context *gin.Context) {

	utils.SetHeader(context)

	var updateMembership model.Membership

	utils.RequestDecoding(context, &updateMembership)

	err := validation.CheckValidation(&updateMembership)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	membership.UpdateMembershipService(context, updateMembership)

}

func DeleteMembershipHandler(context *gin.Context) {

	utils.SetHeader(context)

	var deleteMembersip request.DeleteMembershipRequest

	utils.RequestDecoding(context, &deleteMembersip)

	err := validation.CheckValidation(&deleteMembersip)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	membership.DeleteMembershipService(context, deleteMembersip)
}
