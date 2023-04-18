package handler

import (
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/employee"
	"gym/server/utils"
	"gym/server/validation"

	"github.com/gin-gonic/gin"
)

//	@Description	Creates a new employee record in database
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//
// Param EmpDetails body request.CreateEmployeeRequest true "Employee details"
//
//	@Tags			Employee
//	@Router			/createEmp [post]
func CreateEmployeeHandler(context *gin.Context) {

	utils.SetHeader(context)

	var createEmployee request.CreateEmployeeRequest

	utils.RequestDecoding(context, &createEmployee)

	err := validation.CheckValidation(&createEmployee)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	employee.CreateEmployeeService(context, createEmployee)
}

//	@Description	Gets the list of employees
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//
//	@Tags			Employee
//	@Router			/getEmp [get]
func GetEmployeeHandler(context *gin.Context) {
	utils.SetHeader(context)
	employee.GetEmployeeService(context)
}

//	@Description	Gets the list of employee roles
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//
//	@Tags			Employee
//	@Router			/getEmpRole [get]
func GetEmployeeRoleHandler(context *gin.Context) {
	utils.SetHeader(context)
	employee.GetEmployeeRoleService(context)
}

//	@Description	Give the the count how many trainer are training how many people
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	response.Success
//	@Failure		400	{object}	response.Error
//
//	@Tags			Employee
//	@Router			/empWithuser [get]
func GetUsersWithEmployeesHandler(context *gin.Context) {
	utils.SetHeader(context)
	employee.GetUsersWithEmployeService(context)
}

//	@Description	Marks the employee present for that day
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	response.Success
//	@Failure		400				{object}	response.Error
//	@Param			EmployeeDetails	body		request.EmployeeRequest	true	"Details of employee whose attendence is to be marked"
//	@Tags			Employee
//	@Router			/empAttendence [post]
func EmployeeAttendenceHandler(context *gin.Context) {
	utils.SetHeader(context)

	var empId request.EmployeeRequest

	utils.RequestDecoding(context, &empId)

	err := validation.CheckValidation(&empId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	employee.EmployeeAttendenceService(context, empId)
}

//	@Description	Create the type of employees
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	response.Success
//	@Failure		400				{object}	response.Error
//	@Param			EmployeeTypes	body		model.EmpTypes	true	"Employee type like tranier,cleaner"
//	@Tags			Employee
//	@Router			/empAttendence [post]
func EmployeeRoleHandler(context *gin.Context) {
	utils.SetHeader(context)
	var createEmpRole model.EmpTypes
	utils.RequestDecoding(context, &createEmpRole)
	err := validation.CheckValidation(&createEmpRole)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	employee.EmployeeRoleService(context, createEmpRole)
}

//	@Description	Gets a singler employee
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	response.Success
//	@Failure		400			{object}	response.Error
//	@Param			EmployeeId	body		request.EmployeeRequest	true	"Employee details"
//	@Tags			Employee
//	@Router			/getEmpById [post]
func GetEmployeeByIdHandler(context *gin.Context) {
	utils.SetHeader(context)

	var empId request.EmployeeRequest

	utils.RequestDecoding(context, &empId)
	err := validation.CheckValidation(&empId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	employee.GetEmployeeByIdService(context, empId)
}
