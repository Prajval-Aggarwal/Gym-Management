package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
	"time"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emps []mod.GymEmp
	db.DB.Find(&emps)
	json.NewEncoder(w).Encode(&emps)

}

func CreateEmphandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp mod.GymEmp
	json.NewDecoder(r.Body).Decode(&emp)

	// checking whether the employee role is from the emp_types table or not
	var emptypes mod.EmpTypes
	db.DB.Find(&emptypes)

	db.DB.Create(&emp)

	json.NewEncoder(w).Encode(emp)
}

func SetEmpRole(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var emprole mod.EmpTypes
	err := json.NewDecoder(r.Body).Decode(&emprole)
	if err != nil {
		panic(err)
	}
	result := db.DB.Model(&mod.EmpTypes{}).Where("role =?", emprole.Role).Updates(&emprole)
	if result.Error != nil {
		fmt.Println("error in DB")
	} else if result.RowsAffected == 0 { //if the subs_name is not in record then create new record
		db.DB.Create(&emprole)
		fmt.Fprint(w, "New Employee type added")

	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" Old Price updated successfully"))
	}

}

func GetEmpRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emprole []mod.EmpTypes
	db.DB.Find(&emprole)
	json.NewEncoder(w).Encode(&emprole)
}

func GetEmployeesWithUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("You are watching employees having users under them")
	var emp []mod.EmpWithUser
	// db.DB.Find(&emp).Where("role=?","Trainer").First(&emp)
	query := "SELECT gym_emps.emp_id , gym_emps.emp_name , COUNT(gym_emps.emp_id) as alotted_members FROM gym_emps LEFT JOIN subscriptions ON subscriptions.emp_id = gym_emps.emp_id GROUP BY gym_emps.emp_id HAVING gym_emps.role = 'Trainer';"
	db.DB.Raw(query).Scan(&emp)
	json.NewEncoder(w).Encode(&emp)

func EmpAttendence(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	now := time.Now()

	var temp mod.EmpAttendence
	temp.User_Id = id
	temp.Present = "Present"
	temp.Date = now.Format("02 Jan 2006")
	db.DB.Create(&temp)

}
