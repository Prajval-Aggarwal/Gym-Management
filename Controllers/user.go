package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"net/http"
	"time"
)

// create a new user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	//show the slots table
	var slots []mod.Slot
	query := "SELECT * FROM slots"
	db.DB.Raw(query).Scan(&slots)
	json.NewEncoder(w).Encode(&slots)

	var user mod.User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("get users called")
	var output []mod.Display
	query := "SELECT users.user_id,users.user_name,users.gender, payments.amount, payments.offer_amount,payments.offer,payments.payment_type, payments.payment_id, subscriptions.subs_name, subscriptions.start_date, subscriptions.deleted_at,subscriptions.end_date, subscriptions.duration, subscriptions.emp_id FROM users JOIN payments ON users.user_id = payments.user_id JOIN subscriptions ON payments.payment_id = subscriptions.payment_id;"

	db.DB.Raw(query).Scan(&output)

	json.NewEncoder(w).Encode(&output)

}

// get user by id
func GetUserbyId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("get user by id called")
	id := r.URL.Query().Get("id")
	fmt.Println("id: ", id)
	w.Header().Set("Content-Type", "application/json")
	var u mod.User
	db.DB.Where("user_id = ?", id).Find(&u)
	fmt.Println("user: ", u)
	if u.User_Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with id %s not found", id)
		return
	}
	var user mod.Display
	// db.DB.Joins("User").Joins("Payment").Omit("User").Where("subscriptions.user_id = ?", id).First(&user)

	db.DB.Raw("SELECT users.user_id,users.user_name,users.gender, payments.amount,payments.offer_amount, payments.payment_type, payments.payment_id, subscriptions.subs_name, subscriptions.start_date, subscriptions.deleted_at,subscriptions.end_date, subscriptions.duration, subscriptions.emp_id FROM users JOIN payments ON users.user_id = payments.user_id JOIN subscriptions ON payments.payment_id = subscriptions.payment_id WHERE users.user_id = ?", id).Scan(&user)

	json.NewEncoder(w).Encode(&user)

}

func UserAttendence(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	now := time.Now()

	var temp mod.UAttendence
	temp.User_Id = id
	temp.Present = "Present"
	temp.Date = now.Format("02 Jan 2006")
	db.DB.Create(&temp)
}
