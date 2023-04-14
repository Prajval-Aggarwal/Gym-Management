package subscriptions

import (
	"fmt"
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSubscriptionService(context *gin.Context, subscriptionCreate request.CreateSubRequest) {

	var user model.User
	err := db.FindById(&user, subscriptionCreate.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	dateStr := time.Now().Truncate(time.Hour)

	var subscription model.Subscription
	subscription.Slot_id = subscriptionCreate.SlotId
	subscription.Subs_Name = subscriptionCreate.SubsName
	subscription.StartDate = dateStr.Format("02 Jan 2006")
	subscription.Duration = float64(subscriptionCreate.Duration)
	subscription.EndDate = dateStr.AddDate(0, 0, int(subscription.Duration*30)).Format("02 Jan 2006")

	subscription.User_Id = subscriptionCreate.UserId

	var slots model.Slot

	err = db.FindById(&slots, subscription.Slot_id, "slot_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	fmt.Println("slots data is :", slots)
	slots.Available_space -= 1
	fmt.Println("slots is", slots)
	result := db.UpdateRecord(&slots, slots.SlotId, "slot_id")
	if result.Error != nil {
		response.ErrorResponse(context, 500, result.Error.Error())
		return
	}
	fmt.Println("asdasdsadsadad")

	err = db.CreateRecord(&subscription)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}

	err = db.FindById(&subscription, subscriptionCreate.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	//AddEmptoSub(subscription)

	response.ShowResponse("Success", 200, "Subscription added successfully", subscription, context)

}

func AddEmptoSub(sub model.Subscription) {

	var employee model.GymEmp

	sub.Emp_Id = SelectRand().Emp_Id
	sub.Emp_name = SelectRand().Emp_name
	if employee.Role != "Trainer" {
		fmt.Println("Alotted employee can only be a trainer , please add a trainer!!")

	}

	db.UpdateRecord(&sub, sub.User_Id, "user_id")

}

func SelectRand() *model.GymEmp {

	var employee model.GymEmp
	query := "SELECT * FROM gym_emps  WHERE gym_emps.role = 'Trainer' ORDER BY RANDOM()  LIMIT 1;"
	err := db.QueryExecutor(query, &employee)
	if err != nil {
		return nil
	}
	return &employee
}

func EndSubscriptionService(context *gin.Context, subscriptionEnd request.EndSubRequest) {

	now := time.Now().Truncate(24 * time.Hour)
	var subscription model.Subscription
	var payment model.Payment

	err := db.FindById(&subscription, subscriptionEnd.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	// if subscription.Payment_Id == "" {
	// 	fmt.Println("Payment not done")
	// 	// db.DB.Where("user_id=?", id).Delete(&subcription)
	// 	return
	// }

	// db.FindById(&payment, subscription.Payment_Id, "payment_id")
	startDate, _ := time.Parse("02 Jan 2006", subscription.StartDate)
	temp := now.Sub(startDate).Hours() / 24
	duration := float64(temp)

	oneDayMoney := (payment.OfferAmount / (float64(subscription.Duration) * 30))
	MoneyRefund := math.Round((payment.OfferAmount - (duration * oneDayMoney)) / 2)
	subscription.Duration = duration / 30
	subscription.EndDate = now.Format("02 Jan 2006")
	payment.OfferAmount -= MoneyRefund

	result := db.UpdateRecord(&payment, subscriptionEnd.UserId, "user_id")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	}

	result = db.UpdateRecord(&subscription, subscriptionEnd.UserId, "user_id")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	}

	err = db.DeleteRecord(&subscription, subscriptionEnd.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	var slots model.Slot

	db.FindById(&slots, subscription.Slot_id, "slot_id")
	fmt.Println("slots data is :", slots)
	slots.Available_space += 1
	fmt.Println("slots is", slots)
	result = db.UpdateRecord(&slots, slots.SlotId, "slot_id")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"User Deleted successfully",
		subscription,
		context,
	)

}

// func UpdateSubscriptionService(context *gin.Context, subscriptionUpdate request.UpdateSubRequest){
// 	var currentSubscription model.Subscription
// 	var updatedSubscription model.Subscription
// 	var payment model.Payment

// 	err := db.FindById(&currentSubscription , subscriptionUpdate.UserId ,  "user_id")

// 	if err!=nil{
// 		response.ErrorResponse(context, 400, err.Error())
// 		return
// 	}

// 	err = db.FindById(&payment , subscriptionUpdate.UserId , "user_id")

// 	if err!=nil{
// 		response.ErrorResponse(context, 400, err.Error())
// 		return
// 	}

// 	if currentSubscription.Subs_Name == updatedSubscription.Subs_Name {
// 		response.ErrorResponse(context , 409 , "User already accquires that subscription")
// 		return
// 	}

// 	var newAmount float64
// 	var memShip model.Membership

// 	err = db.FindById(&memShip , subscriptionUpdate.SubsName , "subs_name")

// 	if updatedSubscription.Duration == 0 {
// 		newAmount = memShip.Price * currentSubscription.Duration
// 	} else {
// 		newAmount = memShip.Price * updatedSubscription.Duration
// 	}
// 	fmt.Println("new amount is:", newAmount)
// 	if newAmount > oldAmount {
// 		diff := newAmount - oldAmount
// 		fmt.Fprintf(w, "You need to pay %v amount to upgrade your subscription\n", diff)
// 		payment.Amount = newAmount

// 	} else {

// 		newDuration := oldAmount / memShip.Price
// 		currentSubscription.Duration = newDuration
// 	}
// 	currentSubscription.Subs_Name = updatedSubscription.Subs_Name

// 	db.DB.Where("user_id=?", id).Updates(&currentSubscription)
// 	db.DB.Where("user_id=?", id).Updates(&payment)

// 	w.Write([]byte("Subscription updated successfully"))
// }
