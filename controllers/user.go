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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var user mod.User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("get users called")
	var output mod.Display
	query := "SELECT users.user_id,users.user_name,users.gender, payments.amount, payments.payment_type, payments.payment_id, subscriptions.subs_name, subscriptions.start_date, subscriptions.deleted_at,subscriptions.end_date, subscriptions.duration, subscriptions.emp_id FROM users JOIN payments ON users.user_id = payments.user_id JOIN subscriptions ON payments.payment_id = subscriptions.payment_id; "

	db.DB.Raw(query).Scan(&output)

	json.NewEncoder(w).Encode(&output)

}

func EndSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	now := time.Now().Truncate(24 * time.Hour)
	var subs mod.Subscription
	var payment mod.Payment
	db.DB.Where("user_id=?", id).First(&subs)
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
