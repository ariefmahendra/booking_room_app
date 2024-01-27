package helper

import (
	"booking-room/model"
	"booking-room/model/dto"
	"database/sql"
	"log"
)

// Commit POST DATA
func TxQuery(db *sql.DB, payload model.Transaction) (string) {
	tx, err := db.Begin()
	if err != nil{
		return ""
	}

	idTrx := insertRSVP(payload, tx)
	insertFacility(payload, tx, idTrx)

	if err != nil {
		tx.Rollback()
		log.Println(err, "Transaction Rollback!")
		return idTrx
	}
	
	return idTrx
}


// Commit DELETE DATA
func TxDeleteRSVP(db *sql.DB, id string) error {
    tx, err := db.Begin()
	if err != nil{
		return err
	}

	query := "UPDATE tx_room_reservation SET deleted_at = (CURRENT_TIMESTAMP) WHERE id = $1"
	_, err = tx.Exec(query, id)
	if err != nil {
		Validate(err, "deleteRSVP", tx)
		return err
	} else {
		log.Println("Successfully delete reservation with ID:", id)
	}

	var idFcltys []string
	idFcltys, _ = facId(id, tx)

	for _, f := range idFcltys {
		updateStatusFacilityAvailable(f, tx)
		if err != nil {
			Validate(err, "updateFacilityStatus", tx)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err, "Transaction Rollback!")
		return err
	}
	
	return err
}

func TxApproval(db *sql.DB, payload dto.TransactionDTO) (string, error) {
	tx, err := db.Begin()
	if err != nil{
		return payload.Id, err
	}

	query := "UPDATE tx_room_reservation SET approval_status = $1, approval_note = $2 WHERE id = $3"
	_, err = tx.Exec(query, payload.ApproveStatus, payload.ApproveNote, payload.Id)
	if err != nil {
		Validate(err, "updatedRSVP", tx)
		return payload.Id, err
	} else {
		log.Println("Successfully update reservation status with ID:", payload.Id)
	}

	var idFcltys []string
	idFcltys, _ = facId(payload.Id, tx)

	if payload.ApproveStatus == "ACCEPT"{
		for _, f := range idFcltys {
			err := updateStatusFacilityBooked(f, tx)
			if err != nil {
				Validate(err, "updateFacilityStatus", tx)
				break
			}
		}
		
	} else if payload.ApproveStatus == "DECLINE"{
		for _, f := range idFcltys {
			err := updateStatusFacilityAvailable(f, tx)
			if err != nil {
				Validate(err, "updateFacilityStatus", tx)
				break
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err, "Transaction Rollback!")
		return payload.Id, err
	}
	

	return payload.Id, err
}







//  INSERT DATA
func insertRSVP(payload model.Transaction, tx *sql.Tx) string {
	idTrx := ""
	query := "INSERT INTO tx_room_reservation (employee_id, room_id, start_date, end_date, notes, approval_note) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err := tx.QueryRow(query, payload.EmployeeId, payload.RoomId, payload.StartDate, payload.EndDate, payload.Note, payload.ApproveNote).Scan(&idTrx)
	if err != nil {
		Validate(err, "insertRSVP", tx)
	} else {
		log.Println("Successfully inserted data with ID:", idTrx)
	}
	return idTrx
}

func insertFacility(payload model.Transaction, tx *sql.Tx, idRSVP string) {
	query := "INSERT INTO tx_additional (reservation_id, facilities_id) VALUES ($1, $2) RETURNING id"
	for _, f := range payload.Facility {
		_, err := tx.Exec(query, idRSVP, f.Id)
		updateStatusFacilityRequest(f.Id, tx)
		if err != nil {
			Validate(err, "insertFacility", tx)
			return
		}
	}
	log.Println("Facilities inserted successfully.")
}

// SELECT facilities Id
func facId(id string, tx *sql.Tx) ([]string, error) {
	var idFcltys []string
	query := "SELECT facilities_id FROM tx_additional WHERE reservation_id = $1"
	rows, _ := tx.Query(query, id)
	defer rows.Close()
	for rows.Next() {
		var idFclty string
		if err := rows.Scan(&idFclty); err != nil {
			Validate(err, "select id facility", tx)
			return idFcltys, err
		}
		idFcltys = append(idFcltys, idFclty)
	}
	return idFcltys, nil
}

// UPDATE Facilities status
func updateStatusFacilityRequest(id string, tx *sql.Tx) {
	query :=  "UPDATE mst_facilities SET status = 'REQUEST' WHERE id = $1"
	_, err := tx.Exec(query, id)
	if err != nil {
		Validate(err, "updateFacilityStatus", tx)
		return
	}
}

func updateStatusFacilityAvailable(id string, tx *sql.Tx) error{
	query :=  "UPDATE mst_facilities SET status = 'AVAILABLE' WHERE id = $1"
	_, err := tx.Exec(query, id)
	if err != nil {
		Validate(err, "updateFacilityStatus", tx)
		return err
	}
	return err
}

func updateStatusFacilityBooked(id string, tx *sql.Tx)error {
	query :=  "UPDATE mst_facilities SET status = 'BOOKED' WHERE id = $1"
	_, err := tx.Exec(query, id)
	if err != nil {
		Validate(err, "updateFacilityStatus", tx)
		return err
	}
	return err
}


func Validate(err error, message string, tx *sql.Tx)  {
	if err != nil{
		tx.Rollback()
		log.Println(err, "Transaction Rollback!")
	} else{
		log.Println("Successfully " + message + " data!")
	}
}

