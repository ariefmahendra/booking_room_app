package controller

import (
	"booking-room/delivery/middleware"
	"booking-room/model/dto"
	"booking-room/shared/common"
	"booking-room/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeControllerImpl struct {
	employeeUC usecase.EmployeeUC
	middleware *middleware.Middleware
	rg         *gin.RouterGroup
}

func NewEmployeeController(employeeUC usecase.EmployeeUC, middleware *middleware.Middleware, rg *gin.RouterGroup) *EmployeeControllerImpl {
	return &EmployeeControllerImpl{employeeUC: employeeUC, middleware: middleware, rg: rg}
}

func (e *EmployeeControllerImpl) Route() {
	er := e.rg
	er.POST("/", e.CreateEmployee)
	er.PATCH("/:id", e.UpdateEmployee)
	er.DELETE("/:id", e.DeleteEmployee)
	er.GET("/:id", e.GetEmployeeById)
	er.GET("/", e.GetEmployeeByEmail)
	er.GET("/", e.GetEmployees)
	er.GET("/deleted", e.GetDeletedEmployees)
}

func (e *EmployeeControllerImpl) GetDeletedEmployees(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	isOk := common.AuthorizationAdmin(claims)

	if !isOk {
		log.Println("authorization failed because user is not admin")
		common.SendErrorResponse(ctx, http.StatusForbidden, "forbidden")
		return
	}

	pageParam := ctx.Query("page")
	page, _ := strconv.Atoi(pageParam)

	sizeParam := ctx.Query("size")
	size, _ := strconv.Atoi(sizeParam)

	employees, paging, err := e.employeeUC.GetDeletedEmployees(page, size)
	if err != nil {
		log.Printf("failed to get deleted employees : %v\n", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employees")
		return
	}

	common.SendSuccessPagedResponse(ctx, http.StatusOK, employees, paging)
}

func (e *EmployeeControllerImpl) CreateEmployee(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	if ok := common.AuthorizationAdmin(claims); !ok {
		log.Println("authorization failed because user is not admin")
		common.SendErrorResponse(ctx, http.StatusForbidden, "Forbidden")
		return
	}

	var employeeReq dto.EmployeeCreateRequest

	err := ctx.Bind(&employeeReq)
	if err != nil {
		log.Printf("failed to bind json create employee : %v", err)
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	employee := common.RequestToEmployeeModel(employeeReq)

	if ok := common.ValidateEmail(employee.Email); !ok {
		log.Printf("invalid email : %v", employee.Email)
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid email")
		return
	}

	employeeDto, err := e.employeeUC.CreteEmployee(employee)
	if err != nil {
		log.Printf("failed to create employee : %v", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to create employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusCreated, employeeDto)
}

func (e *EmployeeControllerImpl) UpdateEmployee(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	if ok := common.AuthorizationAdmin(claims); !ok {
		log.Printf("authorization failed : %v", claims)
		common.SendErrorResponse(ctx, http.StatusForbidden, "Forbidden")
		return
	}

	var employeeReq dto.EmployeeCreateRequest

	err := ctx.Bind(&employeeReq)
	if err != nil {
		log.Printf("failed to bind json update employee: %v", err)
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	employee := common.RequestToEmployeeModel(employeeReq)

	employeeId := ctx.Param("id")
	employee.Id = employeeId

	if ok := common.ValidateEmail(employee.Email); !ok {
		log.Printf("invalid email : %v", employee.Email)
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid email")
		return
	}

	employeeDto, err := e.employeeUC.UpdateEmployee(employee)
	if err != nil {
		log.Println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to updating employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusCreated, employeeDto)
}

func (e *EmployeeControllerImpl) DeleteEmployee(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	if ok := common.AuthorizationAdmin(claims); !ok {
		log.Printf("authorization failed : %v", claims)
		common.SendErrorResponse(ctx, http.StatusForbidden, "Forbidden")
		return
	}

	employeeId := ctx.Param("id")

	err := e.employeeUC.DeleteEmployeeById(employeeId)
	if err != nil {
		log.Println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to delete employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, "success")
}

func (e *EmployeeControllerImpl) GetEmployeeById(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	ok := common.AuthorizationAdmin(claims)

	employeeId := ctx.Param("id")

	if employeeId != claims.Id && !ok {
		log.Println("authorization failed because user id not match and user is not admin")
		common.SendErrorResponse(ctx, http.StatusForbidden, "Forbidden access to other user")
		return
	}

	employeeById, err := e.employeeUC.GetEmployeeById(employeeId)
	if err != nil {
		log.Println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, employeeById)
}

func (e *EmployeeControllerImpl) GetEmployeeByEmail(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	ok := common.AuthorizationAdmin(claims)

	employeeEmail := ctx.Query("email")

	if employeeEmail != claims.Email && !ok {
		log.Println("authorization failed because user id not match and user is not admin")
		common.SendErrorResponse(ctx, http.StatusForbidden, "Forbidden access to other user")
		return
	}

	employeeByEmail, err := e.employeeUC.GetEmployeeByEmail(employeeEmail)
	if err != nil {
		log.Println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employee")
		return
	}

	common.SendSuccessResponse(ctx, http.StatusOK, employeeByEmail)
}

func (e *EmployeeControllerImpl) GetEmployees(ctx *gin.Context) {
	claims := e.middleware.GetUser(ctx)
	isOk := common.AuthorizationAdmin(claims)

	if !isOk {
		log.Println("authorization failed because user is not admin")
		common.SendErrorResponse(ctx, http.StatusForbidden, "forbidden")
		return
	}

	pageParam := ctx.Query("page")
	page, _ := strconv.Atoi(pageParam)

	sizeParam := ctx.Query("size")
	size, _ := strconv.Atoi(sizeParam)

	employees, paging, err := e.employeeUC.GetEmployees(page, size)
	if err != nil {
		log.Printf("failed to get employees : %v\n", err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, "failed to get employees")
		return
	}

	common.SendSuccessPagedResponse(ctx, http.StatusOK, employees, paging)
}
