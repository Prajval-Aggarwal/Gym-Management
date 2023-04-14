package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/subscriptions"
	"gym/server/utils"
	"gym/server/validation"
	"github.com/gin-gonic/gin"
)

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

// func UpdateSubscriptionHandler(context *gin.Context){
// 	context.Writer.Header().Set("Content-Type", "application/json")

// 	var subscriptionEnd request.UpdateSubRequest

// 	reqBody, _ := ioutil.ReadAll(context.Request.Body)

// 	err := json.Unmarshal(reqBody, &subscriptionEnd)
// 	if err != nil {
// 		response.ErrorResponse(context, 400, err.Error())
// 		return
// 	}

// 	subscriptions.UpdateSubscriptionService(context, subscriptionEnd)
// }
