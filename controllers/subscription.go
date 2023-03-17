package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"log"
	"math"
	"net/http"
	"time"
)

func CreateSubsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var sub mod.Subscription
	json.NewDecoder(r.Body).Decode(&sub)
	id := r.URL.Query().Get("id")
	fmt.Println("id is", id)
	dateStr := time.Now().Truncate(time.Hour)
	sub.StartDate = dateStr.Format("02 Jan 2006")

	sub.EndDate = dateStr.AddDate(0, 0, int(sub.Duration*30)).Format("02 Jan 2006")

	sub.User_Id = id
	db.DB.Create(&sub)
	json.NewEncoder(w).Encode(&sub)

}

func UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	var currSubs mod.Subscription
	var updatedSubs mod.Subscription
	var payment mod.Payment

	db.DB.Where("user_id=?", id).Find(&currSubs)
	db.DB.Where("user_id=?", id).Find(&payment)
	json.NewDecoder(r.Body).Decode(&updatedSubs)

	if currSubs.Subs_Name == updatedSubs.Subs_Name {
		w.Write([]byte("User already accquires that subscription"))
		return
	}

	var newAmount float64
	var memShip mod.SubsType
	db.DB.Where("subs_name=?", updatedSubs.Subs_Name).First(&memShip)
	if updatedSubs.Duration == 0 {
		newAmount = memShip.Price * currSubs.Duration
	} else {
		newAmount = memShip.Price * updatedSubs.Duration
	}
	oldAmount := payment.Amount
	fmt.Println("new amount is:", newAmount)
	if newAmount > oldAmount {
		diff := newAmount - oldAmount
		fmt.Fprintf(w, "You need to pay %v amount to upgrade your subscription\n", diff)
		payment.Amount = newAmount

	} else {

		newDuration := oldAmount / memShip.Price
		currSubs.Duration = newDuration
	}
	currSubs.Subs_Name = updatedSubs.Subs_Name

	db.DB.Where("user_id=?", id).Updates(&currSubs)
	db.DB.Where("user_id=?", id).Updates(&payment)

	w.Write([]byte("Subscription updated successfully"))

}

func EndSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	now := time.Now().Truncate(24 * time.Hour)
	var subs mod.Subscription
	var payment mod.Payment
	db.DB.Where("user_id=?", id).Find(&subs)
	if subs.Payment_Id == "" {
		fmt.Println("Payment not done")
		db.DB.Where("user_id=?", id).Delete(&subs)
		return
	}
	db.DB.Where("payment_id=?", subs.Payment_Id).First(&payment)
	startDate, err := time.Parse("02 Jan 2006", subs.StartDate)
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
	oneDayMoney := (payment.Amount / (float64(subs.Duration) * 30))
	MoneyRefund := math.Round((payment.Amount - (duration * oneDayMoney)) / 2)
	subs.Duration = duration / 30

	payment.Amount -= MoneyRefund
	db.DB.Where("user_id=?", id).Updates(&payment)
	db.DB.Where("user_id=?", id).Updates(&subs)
	db.DB.Where("user_id=?", id).Delete(&subs)
	w.Write([]byte("Deleted user sucessfully.."))

}
