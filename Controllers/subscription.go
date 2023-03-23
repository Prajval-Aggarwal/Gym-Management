package Controllers

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"log"
	"math"
	"net/http"
	"time"
)

func CreatSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var subscription mod.Subscription
	json.NewDecoder(r.Body).Decode(&subscription)
	id := r.URL.Query().Get("id")
	var u mod.User
	db.DB.Where("user_id = ?", id).Find(&u)
	fmt.Println("user: ", u)
	if u.User_Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with id %s not found", id)
		return
	}
	// fmt.Println("id is", id)
	dateStr := time.Now().Truncate(time.Hour)
	subscription.StartDate = dateStr.Format("02 Jan 2006")

	subscription.EndDate = dateStr.AddDate(0, 0, int(subscription.Duration*30)).Format("02 Jan 2006")

	subscription.User_Id = id
	// slot selection for user
	var slots mod.Slot
	db.DB.Where("id=?", subscription.Slot_id).Find(&slots)
	slots.Available_space -= 1
	db.DB.Where("id=?",slots.ID).Updates(&slots)
  
	// sort kr dena -> rajan

	db.DB.Create(&subscription)

	AddEmptoSub(subscription)

	json.NewEncoder(w).Encode(&subscription)
	// json.NewEncoder(w).Encode(&sub)

	//show bill according to the subscription chosen and duration given
	var subcription_type mod.SubsType
	db.DB.Where("subs_name=?", subscription.Subs_Name).First(&subcription_type)

	//bill amount if duration is 6 months or 12 months
	var billamount float64
	if subscription.Duration == 6 {
		//10% discount
		billamount = (subcription_type.Price * subscription.Duration) * 0.9
		fmt.Fprintln(w, "10% Discount applied")

	} else if subscription.Duration == 12 {
		//20% discount
		billamount = (subcription_type.Price * subscription.Duration) * 0.8
		fmt.Fprintln(w, "20% Discount applied")

	} else {
		billamount = subcription_type.Price * subscription.Duration
	}
	fmt.Fprint(w, "\n\n")
	fmt.Fprint(w, "BILL OF SUBSCRIPTION\n\n")

	fmt.Fprintf(w, " Subscription:%s \n Duration:%d \n Bill Amount:%d ", subscription.Subs_Name, int(subscription.Duration), int(billamount))

}

func UpdateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	var currentSubscription mod.Subscription
	var updatedSubscription mod.Subscription
	var payment mod.Payment

	db.DB.Where("user_id=?", id).Find(&currentSubscription)
	db.DB.Where("user_id=?", id).Find(&payment)
	json.NewDecoder(r.Body).Decode(&updatedSubscription)

	if currentSubscription.Subs_Name == updatedSubscription.Subs_Name {
		w.Write([]byte("User already accquires that subscription"))
		return
	}

	var newAmount float64
	var memShip mod.SubsType
	db.DB.Where("subs_name=?", updatedSubscription.Subs_Name).First(&memShip)
	if updatedSubscription.Duration == 0 {
		newAmount = memShip.Price * currentSubscription.Duration
	} else {
		newAmount = memShip.Price * updatedSubscription.Duration
	}
	oldAmount := payment.Amount
	fmt.Println("new amount is:", newAmount)
	if newAmount > oldAmount {
		diff := newAmount - oldAmount
		fmt.Fprintf(w, "You need to pay %v amount to upgrade your subscription\n", diff)
		payment.Amount = newAmount

	} else {

		newDuration := oldAmount / memShip.Price
		currentSubscription.Duration = newDuration
	}
	currentSubscription.Subs_Name = updatedSubscription.Subs_Name

	db.DB.Where("user_id=?", id).Updates(&currentSubscription)
	db.DB.Where("user_id=?", id).Updates(&payment)

	w.Write([]byte("Subscription updated successfully"))

}

func EndSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	now := time.Now().Truncate(24 * time.Hour)
	var subcription mod.Subscription
	var payment mod.Payment
	db.DB.Where("user_id=?", id).Find(&subcription)
	if subcription.Payment_Id == "" {
		fmt.Println("Payment not done")
		db.DB.Where("user_id=?", id).Delete(&subcription)
		return
	}
	db.DB.Where("payment_id=?", subcription.Payment_Id).First(&payment)
	startDate, err := time.Parse("02 Jan 2006", subcription.StartDate)
	if err != nil {
		log.Fatal(err)
	}
	temp := now.Sub(startDate).Hours() / 24
	duration := float64(temp)
	fmt.Println("Duration is", duration)
	// if duration < 30 {
	// 	http.Error(w, "Cannot end membership before one month", http.StatusBadRequest)
	// 	return
	// }
	oneDayMoney := (payment.OfferAmount / (float64(subcription.Duration) * 30))
	MoneyRefund := math.Round((payment.OfferAmount - (duration * oneDayMoney)) / 2)
	subcription.Duration = duration / 30

	payment.OfferAmount -= MoneyRefund
	db.DB.Where("user_id=?", id).Updates(&payment)
	db.DB.Where("user_id=?", id).Updates(&subcription)
	db.DB.Where("user_id=?", id).Delete(&subcription)
	w.Write([]byte("Deleted user sucessfully.."))

}

// adding employee to subscription
func AddEmptoSub(sub mod.Subscription) {
	fmt.Println("Adding employee to subscription")
	var employee mod.GymEmp
	// random employee id
	//! pick the latest updated trainer
	fmt.Println("from random id: ", SelectRand().Emp_Id)
	// db.DB.Order("created_at desc").Limit(1).Find(&emp)

	fmt.Println("employeeid: ", SelectRand().Emp_Id)
	sub.Emp_Id = SelectRand().Emp_Id
	sub.Emp_name = SelectRand().Emp_name
	if employee.Role != "Trainer" {
		fmt.Println("Alotted employee can only be a trainer , please add a trainer!!")

	}
	fmt.Println("userid: ", sub.User_Id)
	db.DB.Model(&mod.Subscription{}).Where("user_id =?", sub.User_Id).Updates(&sub)
}

func SelectRand() mod.GymEmp {
	fmt.Println("fetching random users")
	var employee mod.GymEmp
	query := "SELECT * FROM gym_emps  WHERE gym_emps.role = 'Trainer' ORDER BY RANDOM()  LIMIT 1;"
	db.DB.Raw(query).Scan(&employee)
	return employee
}

func GetSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subscriptions []mod.Subscription
	db.DB.Find(&subscriptions)
	json.NewEncoder(w).Encode(&subscriptions)

}
