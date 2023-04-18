package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/subscriptions"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

//	@Description	Creates a new subscription for the user
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			SubsDetails	body		request.CreateSubRequest	true	"Subscription details"
//	@Tags			Subscription
//	@Router			/createSubscription [post]
func CreateSubscriptionHandler(context *gin.Context) {

	utils.SetHeader(context)

	var subscriptionCreate request.CreateSubRequest

	utils.RequestDecoding(context, &subscriptionCreate)

	err := validation.CheckValidation(&subscriptionCreate)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	subscriptions.CreateSubscriptionService(context, subscriptionCreate)
}

//	@Description	Ends the subscription for the user
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			SubsDetails	body		request.EndSubRequest	true	"Subscription details"
//	@Tags			Subscription
//	@Router			/endSubscription [delete]
func EndSubscriptionHandler(context *gin.Context) {

	utils.SetHeader(context)

	var subscriptionEnd request.EndSubRequest

	utils.RequestDecoding(context, &subscriptionEnd)

	err := validation.CheckValidation(&subscriptionEnd)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	subscriptions.EndSubscriptionService(context, subscriptionEnd)
}

//	@Description	updates the subscription for the user
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			SubsDetails	body		request.UpdateSubRequest	true	"Subscription details"
//	@Tags			Subscription
//	@Router			/updateSubscription [put]
func UpdateSubscriptionHandler(context *gin.Context) {
	utils.SetHeader(context)

	var subscriptionEnd request.UpdateSubRequest

	utils.RequestDecoding(context, &subscriptionEnd)

	subscriptions.UpdateSubscriptionService(context, subscriptionEnd)
}
