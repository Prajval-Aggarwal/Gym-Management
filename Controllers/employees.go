package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/Models"
	"net/http"
	"time"
)

func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var employees []mod.GymEmp
	db.DB.Find(&employees)
	json.NewEncoder(w).Encode(&employees)

}

func CreateEmployeehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var employees mod.GymEmp
	json.NewDecoder(r.Body).Decode(&employees)

	// checking whether the employee role is from the emp_types table or not
	var emptypes mod.EmpTypes
	db.DB.Find(&emptypes)

	db.DB.Create(&employees)

	json.NewEncoder(w).Encode(employees)
}

func EmployeeRoleHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var employeeRole mod.EmpTypes
	err := json.NewDecoder(r.Body).Decode(&employeeRole)
	if err != nil {
		panic(err)
	}
	result := db.DB.Model(&mod.EmpTypes{}).Where("role =?", employeeRole.Role).Updates(&employeeRole)
	if result.Error != nil {
		fmt.Println("error in DB")
	} else if result.RowsAffected == 0 { //if the subs_name is not in record then create new record
		db.DB.Create(&employeeRole)
		fmt.Fprint(w, "New Employee type added")

	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" Old Price updated successfully"))
	}

}

func GetEmployeeRoleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var employeeRoles []mod.EmpTypes
	db.DB.Find(&employeeRoles)
	json.NewEncoder(w).Encode(&employeeRoles)
}

func GetUsersWithEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("You are watching employees having users under them")
	var employee []mod.EmpWithUser
	// db.DB.Find(&emp).Where("role=?","Trainer").First(&emp)
	query := "SELECT gym_emps.emp_id , gym_emps.emp_name , COUNT(gym_emps.emp_id) as alotted_members FROM gym_emps LEFT JOIN subscriptions ON subscriptions.emp_id = gym_emps.emp_id GROUP BY gym_emps.emp_id HAVING gym_emps.role = 'Trainer';"
	db.DB.Raw(query).Scan(&employee)
	json.NewEncoder(w).Encode(&employee)
}
func EmployeeAttendenceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	now := time.Now()

	var employee mod.EmpAttendence
	employee.User_Id = id
	employee.Present = "Present"
	employee.Date = now.Format("02 Jan 2006")
	db.DB.Create(&employee)

}
