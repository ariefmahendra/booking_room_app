package controller

import (
	"booking-room/delivery/middleware"
	"booking-room/mocks/usecase_mock"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TrxRsvControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	tum *usecase_mock.RsvUseCaseMock
	// amm *midlleware_mock.AuthorMiddlewareMock
}

var expectedList = dto.TransactionDTO{
	Id:            "ID001",
	EmployeeId:    "EM001",
	EmplyName:     "Budi",
	RoomCode:      "R001",
	StartDate:     parseTime("2024-01-25T09:00:00Z"),
	EndDate:       parseTime("2024-01-27T11:00:00Z"),
	Note:          "Team Meeting",
	ApproveStatus: "PENDING",
	ApproveNote:   "Department Briefing",
	Facility: []dto.Facility{
		{
			Id:   "F001",
			Code: "PRJ3",
			Type: "projector",
		},
	},
}

var expectedPage = shared_model.Paging{
	Page:        1,
	TotalPages:  5,
	TotalRows:   1,
	RowsPerPage: 5,
}

var tesTransactionDTO = dto.PayloadReservationDTO{
	Id:            "ID001",
	Email:         "budi@mail.com",
	RoomCode:      "R001",
	StartDate:     pointerTime("2024-01-25T09:00:00Z"),
	EndDate:       pointerTime("2024-01-27T11:00:00Z"),
	Note:          "Team Meeting",
	Facilities: []dto.Facility{
		{
			Id:   "F001",
			Code: "PRJ3",
			Type: "projector",
		},
	},
}

func parseTime(timeStr string) time.Time{
	layout := "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return parsedTime
}

func pointerTime(timeStr string) *time.Time {
	layout := "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}
	return &parsedTime
}


func (s *TrxRsvControllerTestSuite) SetupTest() {
	s.tum = new(usecase_mock.RsvUseCaseMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	s.rg = r.Group("/api/v1/reservation")
}

func (s *TrxRsvControllerTestSuite) TestGetList_Success() {
	s.tum.On("List", 1, 5).Return([]dto.TransactionDTO{expectedList}, shared_model.Paging{}, nil)

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg) 
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.getAll(ctx)

	s.Equal(http.StatusOK, record.Code)
}


func (s *TrxRsvControllerTestSuite) TestGetID_Success() {
	s.tum.On("GetID", "ID001").Return(expectedList, nil)

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation/get/ID001", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.getID(ctx)

	s.Equal(http.StatusOK, record.Code)
	s.Contains(record.Body.String(), "ID001")
}

func (s *TrxRsvControllerTestSuite) TestGetID_Failed() {
	s.tum.On("GetID", "nonexistentID").Return(dto.TransactionDTO{}, errors.New("ID not found"))

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation/get/nonexistentID", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.getID(ctx)

	s.Equal(http.StatusBadRequest, record.Code)
	s.Contains(record.Body.String(), "ID not found")
}

func (s *TrxRsvControllerTestSuite) TestGetEmployee_Success() {
	expectedList := []dto.TransactionDTO{ expectedList }
	expectedPaging := expectedPage
	s.tum.On("GetEmployee", "employeeID", 1, 5).Return(expectedList, expectedPaging, nil)

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation/employee/employeeID?page=1&size=5", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.getEmployee(ctx)

	s.Equal(http.StatusOK, record.Code)
}

func (s *TrxRsvControllerTestSuite) TestGetEmployee_Failed() {
	s.tum.On("GetEmployee", "nonexistentEmployeeID", 1, 5).Return([]dto.TransactionDTO{}, shared_model.Paging{}, errors.New("Employee ID not found"))

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation/employee/nonexistentEmployeeID?page=1&size=5", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.getEmployee(ctx)

	s.Equal(http.StatusBadRequest, record.Code)
	s.Contains(record.Body.String(), "Employee ID not found")
}

// Test for /approval endpoint success scenario
func (s *TrxRsvControllerTestSuite) TestGetApproval_Success() {
	expectedList := []dto.TransactionDTO{ expectedList }
	expectedPaging := expectedPage
	s.tum.On("GetApprovalList", 1, 5).Return(expectedList, expectedPaging, nil)

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation/approval?page=1&size=5", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.acceptRSVP(ctx)

	s.Equal(http.StatusOK, record.Code)
}

func (s *TrxRsvControllerTestSuite) TestGetApproval_Failed() {
	s.tum.On("GetApprovalList", 1, 5).Return([]dto.TransactionDTO{}, shared_model.Paging{}, errors.New("Failed to get approval list"))

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	request, err := http.NewRequest("GET", "/api/v1/reservation/approval?page=1&size=5", nil)
	s.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.acceptRSVP(ctx)

	s.Equal(http.StatusInternalServerError, record.Code)
	s.Contains(record.Body.String(), "Failed to get approval list")
}


func (s *TrxRsvControllerTestSuite) TestAcceptRSVP_Success() {
	expectedTransaction := expectedList
	s.tum.On("UpdateStatus", mock.AnythingOfType("dto.TransactionDTO")).Return(expectedTransaction, nil)

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	jsonPayload := `{"id": "ID001", "approveStatus": "ACCEPT"}`
	request, err := http.NewRequest("PUT", "/api/v1/reservation/approval", strings.NewReader(jsonPayload))
	s.NoError(err)
	request.Header.Set("Content-Type", "application/json")

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.acceptRSVP(ctx)

	s.Equal(http.StatusCreated, record.Code)
}

func (s *TrxRsvControllerTestSuite) TestAcceptRSVP_Failed() {
	s.tum.On("UpdateStatus", mock.AnythingOfType("dto.TransactionDTO")).Return(dto.TransactionDTO{}, errors.New("Failed to update RSVP status"))

	mockMiddleware := &middleware.Middleware{}

	rsvController := NewTrxRsvpController(s.tum, mockMiddleware, s.rg)
	rsvController.Route()

	jsonPayload := `{"id": "ID001", "approveStatus": "ACCEPT"}`
	request, err := http.NewRequest("PUT", "/api/v1/reservation/approval", strings.NewReader(jsonPayload))
	s.NoError(err)
	request.Header.Set("Content-Type", "application/json")

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request

	rsvController.acceptRSVP(ctx)

	s.Equal(http.StatusBadRequest, record.Code)
	s.Contains(record.Body.String(), "Failed to update RSVP status")
}


