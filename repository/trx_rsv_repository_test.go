package repository

import (
	"booking-room/helper"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type TrxRsvRepoTestSuite struct {
	suite.Suite
	mockDB *sql.DB
	mockSql sqlmock.Sqlmock
	repo TrxRsvRepository
}

func (t *TrxRsvRepoTestSuite) SetupSuite()  {
	db, mock, _ := sqlmock.New()
	t.mockDB = db
	t.mockSql = mock
	t.repo = NewTrxRsvRepository(db)
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

func (t *TrxRsvRepoTestSuite) TestGetID_Success() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "code_room", "start_date", "end_date", "notes", "approval_status", "approval_note", "created_at", "updated_at", "deleted_at"}).
		AddRow(expectedList.Id, expectedList.EmployeeId, expectedList.EmplyName, expectedList.RoomCode, expectedList.StartDate, expectedList.EndDate,
			expectedList.Note, expectedList.ApproveStatus, expectedList.ApproveNote, expectedList.CreateAt, expectedList.UpdateAt, expectedList.DeleteAt)

	t.mockSql.ExpectQuery(regexp.QuoteMeta(`
			SELECT 
				tx.id, 
				tx.employee_id, 
				emp.name, 
				room.code_room, 
				tx.start_date, 
				tx.end_date,
				tx.notes,
				tx.approval_status,
				tx.approval_note,
				tx.created_at,
				tx.updated_at,
				tx.deleted_at
			FROM
				tx_room_reservation tx
			JOIN
				mst_employee emp ON tx.employee_id = emp.id
			JOIN
				mst_room room ON tx.room_id = room.id
			WHERE tx.id = $1`)).
		WithArgs(expectedList.Id).
		WillReturnRows(rows)

	result, err := t.repo.GetID(expectedList.Id)

	t.NoError(err)
	t.Equal(expectedList, result)
}

func (t *TrxRsvRepoTestSuite) TestGetID_Failed() {
    rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "code_room", "start_date", "end_date", "notes", "approval_status", "approval_note", "created_at", "updated_at", "deleted_at"}).
        AddRow(expectedList.Id, expectedList.EmployeeId, expectedList.EmplyName, expectedList.RoomCode, expectedList.StartDate, expectedList.EndDate,
            expectedList.Note, expectedList.ApproveStatus, expectedList.ApproveNote, expectedList.CreateAt, expectedList.UpdateAt, expectedList.DeleteAt)

    expectedID := "ID005"

    t.mockSql.ExpectQuery(regexp.QuoteMeta(`
			SELECT 
				tx.id, 
				tx.employee_id, 
				emp.name, 
				room.code_room, 
				tx.start_date, 
				tx.end_date,
				tx.notes,
				tx.approval_status,
				tx.approval_note,
				tx.created_at,
				tx.updated_at,
				tx.deleted_at
			FROM
				tx_room_reservation tx
			JOIN
				mst_employee emp ON tx.employee_id = emp.id
			JOIN
				mst_room room ON tx.room_id = room.id
			WHERE tx.id = $1`)).
        WithArgs(expectedID).
        WillReturnRows(rows)

    result, err := t.repo.GetID(expectedID)

    t.NotEqual(expectedList, result)
    t.NoError(err)
}


func (t *TrxRsvRepoTestSuite) TestGetEmployee_Success() {
    rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "code_room", "start_date", "end_date", "notes", "approval_status", "approval_note", "created_at", "updated_at", "deleted_at"}).
        AddRow(expectedList.Id, expectedList.EmployeeId, expectedList.EmplyName, expectedList.RoomCode, expectedList.StartDate, expectedList.EndDate,
            expectedList.Note, expectedList.ApproveStatus, expectedList.ApproveNote, expectedList.CreateAt, expectedList.UpdateAt, expectedList.DeleteAt)
    expectedQuery := `
        SELECT 
            tx.id, 
            tx.employee_id, 
            emp.name, 
            room.code_room, 
            tx.start_date, 
            tx.end_date,
            tx.notes,
            tx.approval_status,
            tx.approval_note,
            tx.created_at,
            tx.updated_at,
            tx.deleted_at
        FROM
            tx_room_reservation tx
        JOIN
            mst_employee emp ON tx.employee_id = emp.id
        JOIN
            mst_room room ON tx.room_id = room.id
        WHERE tx.employee_id = $1 AND tx.deleted_at IS NULL
        ORDER BY tx.created_at DESC
        LIMIT $2 OFFSET $3`

    t.mockSql.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
        WithArgs(expectedList.EmployeeId, 1, 5).
        WillReturnRows(rows)

    result, _, err := t.repo.GetEmployee(expectedList.EmployeeId, 1, 5)

    t.NoError(err)
    t.Equal(expectedList, result[0])
}

func (t *TrxRsvRepoTestSuite) TestGetEmployee_Failed() {
    rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "code_room", "start_date", "end_date", "notes", "approval_status", "approval_note", "created_at", "updated_at", "deleted_at"}).
        AddRow("ID005", "DifferentEmployeeID", "DifferentName", "DifferentRoomCode", time.Now(), time.Now(), "DifferentNote", "PENDING", "DifferentApprovalNote", time.Now(), time.Now(), nil)

    expectedQuery := `
        SELECT 
            tx.id, 
            tx.employee_id, 
            emp.name, 
            room.code_room, 
            tx.start_date, 
            tx.end_date,
            tx.notes,
            tx.approval_status,
            tx.approval_note,
            tx.created_at,
            tx.updated_at,
            tx.deleted_at
        FROM
            tx_room_reservation tx
        JOIN
            mst_employee emp ON tx.employee_id = emp.id
        JOIN
            mst_room room ON tx.room_id = room.id
        WHERE tx.employee_id = $1 AND tx.deleted_at IS NULL
        ORDER BY tx.created_at DESC
        LIMIT $2 OFFSET $3`

    expectedArg1 := expectedList.EmployeeId
    expectedArg2 := 1
    expectedArg3 := 5

    t.mockSql.ExpectQuery(regexp.QuoteMeta(expectedQuery)).
        WithArgs(expectedArg1, expectedArg2, expectedArg3).
        WillReturnRows(rows)

    result, _, err := t.repo.GetEmployee(expectedArg1, expectedArg2, expectedArg3)

    t.Nil(result)

	expectedErrorMsg := fmt.Sprintf("expected query argument 1: %s, but got: %d", expectedArg1, 5)
    t.Require().Error(err)
    t.Contains(err.Error(), expectedErrorMsg)
}

func (t *TrxRsvRepoTestSuite) TestList() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "code_room", "start_date", "end_date", "notes", "approval_status", "approval_note", "created_at", "updated_at", "deleted_at"}).
		AddRow(expectedList.Id, expectedList.EmployeeId, expectedList.EmplyName, expectedList.RoomCode, expectedList.StartDate, expectedList.EndDate,
			expectedList.Note, expectedList.ApproveStatus, expectedList.ApproveNote, expectedList.CreateAt, expectedList.UpdateAt, expectedList.DeleteAt)

	t.mockSql.ExpectQuery("SELECT").
		WithArgs(1, 1).
		WillReturnRows(rows)

	result, _, err := t.repo.List(1, 1)

	t.NoError(err)
	t.Equal(expectedList, result[0])
}

func facId(id string, tx *sql.Tx) ([]string, error) {
    var idFcltys []string
    query := "SELECT facilities_id FROM tx_additional WHERE reservation_id = $1"
    rows, err := tx.Query(query, id)
    if err != nil {
        helper.Validate(err, "select id facility", tx)
        return idFcltys, err
    }
    defer rows.Close()
    for rows.Next() {
        var idFclty string
        if err := rows.Scan(&idFclty); err != nil {
            helper.Validate(err, "select id facility", tx)
            return idFcltys, err
        }
        idFcltys = append(idFcltys, idFclty)
    }
    return idFcltys, nil
}


func (t *TrxRsvRepoTestSuite) TestDeleteResv_Failure() {

	t.mockDB.Begin()
	t.mockSql.ExpectBegin()
	t.mockSql.ExpectExec(regexp.QuoteMeta("UPDATE tx_room_reservation SET deleted_at = (CURRENT_TIMESTAMP) WHERE id = $1")).
		WithArgs(expectedList.Id).
		WillReturnError(errors.New("delete failed"))
	t.mockSql.ExpectRollback()

	result, err := t.repo.DeleteResv(expectedList.Id)

	t.Equal("Reservation Deleted Failed", result)
	t.Error(err)
}



func TestTrxRsvRepoTestSuite(t *testing.T) {
	suite.Run(t, new(TrxRsvRepoTestSuite))
}