package repository

import (
	"booking-room/model"
	"database/sql"
	"log"
	"time"
)

type ReportRepository interface {
	FindAll() ([]model.EmployeeList, error)
	FindAllRoom() ([]model.RoomList, error)
	FindAllFacilities() ([]model.FacilitiesList, error)
	ReservationReport(start, end time.Time) ([]model.ReservationReport, error)
	FacilityTotalReserved(start, end time.Time) ([]model.FacilityTotalReserved, error)
	RoomTotalReserved(start, end time.Time) ([]model.RoomTotalReserved, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) FindAll() ([]model.EmployeeList, error) {
	var Employees []model.EmployeeList
	rows, err := r.db.Query("SELECT name, email, division, position, role, contact FROM mst_employee WHERE deleted_at IS NULL")
	if err != nil {
		log.Println("reportRepository.FindAll", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var employee model.EmployeeList
		if err := rows.Scan(&employee.Name, &employee.Email, &employee.Division, &employee.Position, &employee.Role, &employee.Contact); err != nil {
			log.Println("reportRepository.FindAll", err.Error())
		}
		Employees = append(Employees, employee)
	}
	return Employees, nil
}

func (r *reportRepository) FindAllRoom() ([]model.RoomList, error) {
	var Rooms []model.RoomList
	rows, err := r.db.Query("SELECT code_room, room_type, capacity, facilities FROM mst_room WHERE deleted_at IS NULL")
	if err != nil {
		log.Println("reportRepository.FindAllRoom", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var room model.RoomList
		if err := rows.Scan(&room.CodeRoom, &room.TypeRoom, &room.Capacity, &room.Facilities); err != nil {
			log.Println("reportRepository.FindAllRoom", err.Error())
		}
		Rooms = append(Rooms, room)
	}
	return Rooms, nil
}

func (r *reportRepository) FindAllFacilities() ([]model.FacilitiesList, error) {
	var Facilities []model.FacilitiesList
	rows, err := r.db.Query("SELECT code_name, facilities_type, status FROM mst_facilities WHERE deleted_at IS NULL")
	if err != nil {
		log.Println("reportRepository.FindAllFacilities", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var facilities model.FacilitiesList
		if err := rows.Scan(&facilities.CodeName, &facilities.FacilityType, &facilities.Status); err != nil {
			log.Println("reportRepository.FindAllFacilities", err.Error())
		}
		Facilities = append(Facilities, facilities)
	}
	return Facilities, nil
}

func (r *reportRepository) ReservationReport(start, end time.Time) ([]model.ReservationReport, error) {
	var Reports []model.ReservationReport
	sqlquery := `SELECT
					rr.id AS reservation_id,
					e.name AS employee_name,
					r.code_room,
					rr.start_date,
					rr.end_date,
					rr.notes,
					rr.approval_status,
					rr.approval_note
				FROM
					tx_room_reservation rr
				JOIN
					mst_employee e ON rr.employee_id = e.id
				JOIN
					mst_room r ON rr.room_id = r.id
				WHERE 
					rr.start_date > $1 AND rr.start_date < $2;`
	rows, err := r.db.Query(sqlquery, start, end)
	if err != nil {
		log.Println("reportRepository.FindAllReservationReport", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var report model.ReservationReport
		if err := rows.Scan(&report.ReservationId, &report.EmployeeName, &report.CodeRoom, &report.StartDate, &report.EndDate, &report.Note, &report.ApproveStatus, &report.ApproveNote); err != nil {
			log.Println("reportRepository.FindAllReservationReportAppend", err.Error())
		}
		rows2, err := r.db.Query(`SELECT
									f.facilities_type,
									COUNT(*) AS total_reservations
								FROM
									tx_room_reservation rr
								JOIN
									tx_additional a ON rr.id = a.reservation_id
								JOIN
									mst_facilities f ON a.facilities_id = f.id
								WHERE
									rr.start_date >= $2 AND rr.end_date < $3 AND rr.id = $1
								GROUP BY
									f.facilities_type;`, report.ReservationId, start, end)
		if err != nil {
			log.Println("reportRepository.FindAllReservationReportDetail", err.Error())
		}
		defer rows2.Close()
		additionals := []model.AdditionalReport{}
		for rows2.Next() {
			var additional model.AdditionalReport
			if err := rows2.Scan(&additional.FacilitiesName, &additional.FacilitieQty); err != nil {
				log.Println("reportRepository.FindAllReservationReportDetailAppend", err.Error())
			}
			additionals = append(additionals, additional)
		}
		report.AdditionalFacilitys = additionals
		Reports = append(Reports, report)
	}
	return Reports, nil
}

func (r *reportRepository) FacilityTotalReserved(start, end time.Time) ([]model.FacilityTotalReserved, error) {
	var FacilityTotalReserved []model.FacilityTotalReserved
	rows, err := r.db.Query(`SELECT
								f.facilities_type,
								COUNT(*) AS total_reservations
							FROM
								tx_room_reservation rr
							JOIN
								tx_additional a ON rr.id = a.reservation_id
							JOIN
								mst_facilities f ON a.facilities_id = f.id
							WHERE
								rr.start_date >= $1 AND rr.end_date < $2
							GROUP BY
								f.facilities_type;`, start, end)
	if err != nil {
		log.Println("reportRepository.FacilityTotalReserved", err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var facility model.FacilityTotalReserved
		if err := rows.Scan(&facility.FacilityType, &facility.Total); err != nil {
			log.Println("reportRepository.FacilityTotalReserved", err.Error())
		}
		FacilityTotalReserved = append(FacilityTotalReserved, facility)
	}
	return FacilityTotalReserved, nil
}

func (r *reportRepository) RoomTotalReserved(start, end time.Time) ([]model.RoomTotalReserved, error) {
	var RoomTotalReserved []model.RoomTotalReserved
	rows, err := r.db.Query(`SELECT
								r.room_type,
								COUNT(*) AS total_reservations
							FROM
								tx_room_reservation rr
							JOIN
								mst_room r ON rr.room_id = r.id
							WHERE
								rr.start_date >= $1 AND rr.end_date < $2
							GROUP BY
								r.room_type;`, start, end)
	if err != nil {
		log.Println("reportRepository.RoomTotalReserved", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var room model.RoomTotalReserved
		if err := rows.Scan(&room.RoomType, &room.Total); err != nil {
			log.Println("reportRepository.RoomTotalReserved", err.Error())
		}
		RoomTotalReserved = append(RoomTotalReserved, room)
	}
	return RoomTotalReserved, nil
}
