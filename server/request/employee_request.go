package request

type CreateEmployeeRequest struct {
	Emp_name string `json:"empName" validate:"required"`
	Gender   string `json:"gender" validate:"required,oneof=Male Female Other"`
	Role     string `json:"role" validate:"required"`
}

type EmployeeRequest struct {
	EmpId string `json:"empId" validate:"required"`
}
