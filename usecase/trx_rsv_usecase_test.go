package usecase

import (
	"booking-room/mocks/repo_mock"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

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

var expectedPaging = shared_model.Paging{
	Page:        1,
	TotalPages:  3,
	TotalRows:   1,
	RowsPerPage: 1,
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

type TrxRsvUsecaseSuite struct {
	suite.Suite
	trm *repo_mock.RsvRepoMock
	rrm *repo_mock.RoomRepositoryMock
	ruc RoomUseCase
	tuc TrxRsvUsecase
}

func (s *TrxRsvUsecaseSuite) SetupTest() {
	s.trm = new(repo_mock.RsvRepoMock)
	s.rrm = new(repo_mock.RoomRepositoryMock)
	s.ruc = NewRoomUseCase(s.rrm)
	s.tuc = NewTrxRsvUseCase(s.trm, s.ruc)
}

func (s *TrxRsvUsecaseSuite) TestList() {
	s.trm.On("List", 1, 5).Return([]dto.TransactionDTO{expectedList}, shared_model.Paging{}, nil)

	actual, _, err := s.tuc.List(1, 5)
	s.NoError(err)
	s.Nil(err)
	s.Equal(expectedList, actual[0])
}

func (s *TrxRsvUsecaseSuite) TestGetID_Success() {
	s.trm.On("GetID", expectedList.Id).Return(expectedList, nil)
	actual, err := s.tuc.GetID(expectedList.Id)
	s.NoError(err)
	s.Nil(err)
	s.Equal(expectedList, actual)
}

func (s *TrxRsvUsecaseSuite) TestGetID_Failed() {
	s.trm.On("GetID", expectedList.Id).Return(dto.TransactionDTO{}, nil)
	_, err := s.tuc.GetID(expectedList.Id)
	s.NoError(err)
}

func (s *TrxRsvUsecaseSuite) TestGetEmployee() {
	s.trm.On("GetEmployee", expectedList.EmployeeId, 1, 5).Return([]dto.TransactionDTO{expectedList}, shared_model.Paging{}, nil)
	actual, _, err := s.tuc.GetEmployee(expectedList.EmployeeId, 1, 5)
	s.NoError(err)
	s.Nil(err)
	s.Equal(expectedList, actual[0])
}

func (s *TrxRsvUsecaseSuite) TestPostReservation() {
	s.trm.On("PostReservation", tesTransactionDTO).Return("ID001", nil)
	_, err := s.tuc.PostReservation(tesTransactionDTO)
	s.NoError(err)
}

func TestReservationTestSuite(t *testing.T) {
	suite.Run(t, new(TrxRsvUsecaseSuite))
}
