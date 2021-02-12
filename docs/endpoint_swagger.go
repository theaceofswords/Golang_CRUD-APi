package docs

import "golang-training/app/models"

// swagger:route GET /employee employee-tag idOfEmpEndpoint1
// Retrieves single empoyee record based on parameter ID.
// responses:
//   200: EmpGETResponse
//   404: _ Record Not Found
//   422: _ Invalid Parameter type

//successful operation.
// swagger:response EmpGETResponse
type EmpRequestWrapper struct {
	// in:body
	Body models.Employee
}

//================================================

// swagger:route POST /employee employee-tag idOfEmpEndpoint2
// Add new Employee record.
// responses:
//   200: EmpPOSTResponse
//   409: _ Record already exists
//   422: _ Invalid Body type

//successful operation.
// swagger:response EmpPOSTResponse
type EmpResponseWrapper struct {
	// in:body
	Body models.Employee
}

// swagger:parameters idOfEmpEndpoint2
type EmpParamsWrapper2 struct {
	// Object to be inserted.
	// in:body
	Body models.Employee
}

//====================================================================================

// swagger:route PUT /employee employee-tag idOfEmpEndpoint3
// Update existing Employee record.
// responses:
//   200: EmpPUTResponse
//   404: _ Record not found
//   422: _ Invalid Body type

//successful operation.
// swagger:response EmpPUTResponse
type EmpResponseWrapper3 struct {
	// in:body
	Body models.Employee
}

// swagger:parameters idOfEmpEndpoint3
type EmpParamsWrapper3 struct {
	// Object to update.
	// in:body
	Body models.Employee
}

//===================================================================

// swagger:route DELETE /employee employee-tag idOfEmpEndpoint4
// Update existing Employee record.
// responses:
//   200: _ Successful operation.
//   404: _ Record not found
//   422: _ Invalid Parameter type

// swagger:parameters idOfEmpEndpoint4
type EmpParamsWrapper4 struct {
	// This text will appear as description of your request body.
	// in:body
	Body models.GetParam
}
