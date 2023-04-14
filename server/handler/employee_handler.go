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

func GetEmployeeHandler(context *gin.Context) {
	utils.SetHeader(context)
	employee.GetEmployeeService(context)
}

func GetEmployeeRoleHandler(context *gin.Context) {
	utils.SetHeader(context)
	employee.GetEmployeeRoleService(context)
}

func GetUsersWithEmployeesHandler(context *gin.Context) {
	utils.SetHeader(context)
	employee.GetUsersWithEmployeService(context)
}

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
