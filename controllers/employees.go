package cont

import (
	"encoding/json"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
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

	db.DB.Create(&emp)

	json.NewEncoder(w).Encode(emp)
}
