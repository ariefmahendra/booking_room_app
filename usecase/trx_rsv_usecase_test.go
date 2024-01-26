package usecase

import (
	"booking-room/mocks/repo_mock"
	"booking-room/model/dto"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

var expectedList = dto.TransactionDTO{
	Id:            "ID001",
	EmployeeId:    "",
	EmplyName:     "Budi",
	RoomCode:      "R001",
	StartDate:     parseTime("2024-01-25T09:00:00Z"),
	EndDate:       parseTime("2024-01-25T09:00:00Z"),
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

func parseTime(timeStr string) time.Time {
	layout := "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return parsedTime
}

type TrxRsvUsecaseSuite struct {
	suite.Suite
	trm *repo_mock.RsvRepoMock

}

func (s *TrxRsvUsecaseSuite) SetupTest() {
	s.trm = new(repo_mock.RsvRepoMock)
}

func TestReservationTestSuite(t *testing.T) {
	suite.Run(t, new(TrxRsvUsecaseSuite))
}
