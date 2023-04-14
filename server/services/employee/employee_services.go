package employee

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"time"

	"github.com/gin-gonic/gin"
)

func GetEmployeeService(context *gin.Context) {
	var employees []model.GymEmp
	query := "SELECT * FROM gym_emps"
	err := db.QueryExecutor(query, &employees)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Employees retrieved",
		employees,
		context,
	)
}

func GetEmployeeRoleService(context *gin.Context) {
	var employeeRoles []model.EmpTypes
	query := "SELECT * FROM emp_types"
	err := db.QueryExecutor(query, &employeeRoles)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Employee roles retrieved",
		employeeRoles,
		context,
	)
}

func GetUsersWithEmployeService(context *gin.Context) {
	var employee []model.EmpWithUser
	query := "SELECT gym_emps.emp_id , gym_emps.emp_name , COUNT(gym_emps.emp_id) as alotted_members FROM gym_emps LEFT JOIN subscriptions ON subscriptions.emp_id = gym_emps.emp_id GROUP BY gym_emps.emp_id HAVING gym_emps.role = 'Trainer';"
	err := db.QueryExecutor(query, &employee)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Data retrieved",
		employee,
		context,
	)
}

func CreateEmployeeService(context *gin.Context, Data request.CreateEmployeeRequest) {
	var createEmp model.GymEmp
	createEmp.Emp_name = Data.Emp_name
	createEmp.Gender = Data.Gender
	createEmp.Role = Data.Role
	err := db.CreateRecord(&createEmp)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
	}

	response.ShowResponse(
		"Success",
		200,
		"Employee created successfully",
		createEmp,
		context,
	)
}

func EmployeeAttendenceService(context *gin.Context, userId request.EmployeeRequest) {
	var empAttendence model.EmpAttendence
	now := time.Now()
	empAttendence.User_Id = userId.EmpId
	empAttendence.Present = "Present"
	empAttendence.Date = now.Format("02 Jan 2006")

	err := db.CreateRecord(&empAttendence)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Employee Attendence logged successfully",
		empAttendence,
		context,
	)

}

func EmployeeRoleService(context *gin.Context, empRoleData model.EmpTypes) {
	result := db.UpdateRecord(&empRoleData, empRoleData.Role, "role")
	if result.Error != nil {
		response.ErrorResponse(context, 400, result.Error.Error())
		return
	} else if result.RowsAffected == 0 {
		err := db.CreateRecord(&empRoleData)
		if err != nil {
			response.ErrorResponse(context, 400, err.Error())
			return
		}
		response.ShowResponse(
			"Success",
			200,
			"New Employee Role added",
			empRoleData,
			context,
		)

	} else {
		response.ShowResponse(
			"Success",
			200,
			"Old Role updated successfully",
			empRoleData,
			context,
		)
	}
}

func GetEmployeeByIdService(context *gin.Context, empId request.EmployeeRequest) {
	var empGetter model.User
	err := db.FindById(&empGetter, empId.EmpId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"Employee retrieved",
		empGetter,
		context,
	)
}
