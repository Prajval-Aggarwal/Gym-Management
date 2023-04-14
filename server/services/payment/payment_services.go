package payment

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"

	"github.com/gin-gonic/gin"
)

func MakePaymentService(context *gin.Context, PaymentData request.CreatePaymentRequest) {
	var payment model.Payment
	var subscription model.Subscription
	var membership model.Membership

	payment.Payment_Type = PaymentData.PaymentType
	err := db.FindById(&subscription, PaymentData.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = db.FindById(&membership, subscription.Subs_Name, "mem_name")
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	var billamount float64
	totalAmount := membership.Price * subscription.Duration
	if subscription.Duration == 6 {
		//10% discount
		billamount = totalAmount * 0.9
		payment.Offer = "10%"

	} else if subscription.Duration == 12 {
		//20% discount
		billamount = totalAmount * 0.8
		payment.Offer = "20%"

	} else {
		payment.Offer = "0"
		billamount = totalAmount
	}
	payment.Amount = totalAmount
	payment.OfferAmount = billamount
	payment.User_Id = PaymentData.UserId

	//add razor pay to add payemnt id for now it is uuid
	err = db.CreateRecord(&payment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	subscription.Payment_Id = payment.Payment_Id

	result := db.UpdateRecord(&subscription, PaymentData.UserId, "user_id")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	}

	response.ShowResponse("Success", 200, "Payment done", payment, context)

}
