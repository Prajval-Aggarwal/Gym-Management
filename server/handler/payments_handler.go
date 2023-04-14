package handler

import (
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/payment"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

func MakePaymentHandler(context *gin.Context) {
	utils.SetHeader(context)
	var createPayment request.CreatePaymentRequest

	utils.RequestDecoding(context, &createPayment)

	err := validation.CheckValidation(&createPayment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	payment.MakePaymentService(context, createPayment)
}
