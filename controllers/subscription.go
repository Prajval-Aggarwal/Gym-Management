package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
	"time"
)

func CreateSubsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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

	// jase hi subscribe krwao to use employee alot kro
	 AddEmptoSub(sub)
	

	json.NewEncoder(w).Encode(&sub)

}

// adding employee to subscription
func AddEmptoSub(sub mod.Subscription)  {
	fmt.Println("Adding employee to subscription")
	var emp mod.GymEmp
	// random employee id
	//! pick the latest updated trainer
	fmt.Println("from random id: ",SelectRand().Emp_Id)
	// db.DB.Order("created_at desc").Limit(1).Find(&emp)

	fmt.Println("employeeid: ", SelectRand().Emp_Id)
	sub.Emp_Id = SelectRand().Emp_Id
	sub.Emp_name = SelectRand().Emp_name
	if emp.Role != "Trainer" {
		fmt.Println("Alotted employee can only be a trainer , please add a trainer!!")
		 
	}
	fmt.Println("userid: ", sub.User_Id)
	db.DB.Model(&mod.Subscription{}).Where("user_id =?", sub.User_Id).Updates(&sub)
}

func SelectRand()mod.GymEmp{
	fmt.Println("fetching random users")
	var emp  mod.GymEmp
	query := "SELECT * FROM gym_emps  WHERE gym_emps.role = 'Trainer' ORDER BY RANDOM()  LIMIT 1;"
	db.DB.Raw(query).Scan(&emp)
	return emp
}

func GetSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subs []mod.Subscription
	db.DB.Find(&subs)
	json.NewEncoder(w).Encode(&subs)

}

