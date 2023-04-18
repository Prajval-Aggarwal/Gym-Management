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

//	@Description	Creates a new membership
//	@Accept			json
//	@Produce		json
//	@Success		200					{object}	response.Success
//	@Failure		400					{object}	response.Error
//	@Param			MembershipDetails	body		model.Membership	true	"MembershipDetails"
//	@Tags			Membership
//	@Router			/createMembership [post]
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

//	@Description	Gets the list of memberships
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//
//	@Tags			Membership
//	@Router			/getMembership [get]
func GetMembershipHandler(context *gin.Context) {
	utils.SetHeader(context)
	membership.GetMembershipsService(context)
}

//	@Description	updates the membership
//	@Accept			json
//	@Produce		json
//
//	@Success		200					{object}	response.Success
//	@Failure		400					{object}	response.Error
//
//	@Param			MembershipDetails	body		model.Membership	true	"Membership name"
//	@Tags			Membership
//	@Router			/updateMembership [put]
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

//	@Description	delete a membership
//	@Accept			json
//	@Produce		json
//
//	@Success		200					{object}	response.Success
//	@Failure		400					{object}	response.Error
//
//	@Param			MembershipDetails	body		request.DeleteMembershipRequest	true	"Membership name"
//	@Tags			Membership
//	@Router			/deleteMembership [delete]
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
