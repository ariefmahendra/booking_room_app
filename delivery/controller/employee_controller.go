package controller

import (
	"booking-room/model/dto"
	"booking-room/shared/common"
	"booking-room/shared/shared_model"
	"booking-room/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EmployeeControllerImpl struct {
	employeeUC usecase.EmployeeUC
}

func NewEmployeeController(employeeUC usecase.EmployeeUC) *EmployeeControllerImpl {
	return &EmployeeControllerImpl{employeeUC: employeeUC}
}

func (e *EmployeeControllerImpl) CreateEmployee(ctx *gin.Context) {
	var employeeReq dto.EmployeeCreateRequest

	err := ctx.Bind(&employeeReq)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	employee := common.RequestToEmployeeModel(employeeReq)

	employeeDto, err := e.employeeUC.CreteEmployee(employee)
	if err != nil {
		fmt.Printf("failed to create employee : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to create employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusCreated, employeeDto)
}

func (e *EmployeeControllerImpl) UpdateEmployee(ctx *gin.Context) {
	var employeeReq dto.EmployeeCreateRequest

	err := ctx.Bind(&employeeReq)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	employee := common.RequestToEmployeeModel(employeeReq)

	employeeId := ctx.Param("id")
	employee.Id = employeeId

	employeeDto, err := e.employeeUC.UpdateEmployee(employee)
	if err != nil {
		fmt.Printf("failed to update employee : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to updating employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusCreated, employeeDto)
}

func (e *EmployeeControllerImpl) DeleteEmployee(ctx *gin.Context) {
	employeeId := ctx.Param("id")

	err := e.employeeUC.DeleteEmployeeById(employeeId)
	if err != nil {
		fmt.Printf("failed to delete employee : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to delete employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, "success")
}

func (e *EmployeeControllerImpl) GetEmployeeById(ctx *gin.Context) {
	employeeId := ctx.Param("id")

	employeeById, err := e.employeeUC.GetEmployeeById(employeeId)
	if err != nil {
		fmt.Printf("failed to get employee by id: %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, employeeById)
}

func (e *EmployeeControllerImpl) GetEmployeeByEmail(ctx *gin.Context) {
	employeeEmail := ctx.Param("email")

	employeeByEmail, err := e.employeeUC.GetEmployeeByEmail(employeeEmail)
	if err != nil {
		fmt.Printf("failed to get employee by email : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, employeeByEmail)
}

func (e *EmployeeControllerImpl) GetEmployees(ctx *gin.Context) {
	pageParam := ctx.Query("page")
	page, _ := strconv.Atoi(pageParam)

	sizeParam := ctx.Query("size")
	size, _ := strconv.Atoi(sizeParam)

	employees, paging, err := e.employeeUC.GetEmployees(page, size)
	if err != nil {
		fmt.Printf("failed to get employees : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employees")
		return
	}

	response := shared_model.PagedResponse{
		Status: shared_model.Status{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data:   employees,
		Paging: paging,
	}

	common.SendSuccessResponse(ctx, http.StatusOK, response)
}
