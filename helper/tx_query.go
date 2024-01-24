package helper

import (
	"booking-room/model"
	"database/sql"
	"log"
)

func TxQuery(db *sql.DB, payload model.Transaction) (string) {
	tx, err := db.Begin()
	if err != nil{
		return ""
	}

	idTrx := insertRSVP(payload, tx)
	insertFacility(payload, tx, idTrx)

	err = tx.Commit()
	if err != nil{
		panic(err)
	}else {
		log.Println("Transaction Commited!")
	}

	return idTrx
}

func insertRSVP(payload model.Transaction, tx *sql.Tx) string {
	idTrx := ""
	query := "INSERT INTO tx_room_reservation (employee_id, room_id, start_date, end_date, notes, approval_note) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err := tx.QueryRow(query, payload.EmployeeId, payload.RoomId, payload.StartDate, payload.EndDate, payload.Note, payload.ApproveNote).Scan(&idTrx)
	if err != nil {
		validate(err, "insertRSVP", tx)
	} else {
		log.Println("Successfully inserted data with ID:", idTrx)
	}
	return idTrx
}

func insertFacility(payload model.Transaction, tx *sql.Tx, idRSVP string) {
	query := "INSERT INTO tx_additional (reservation_id, facilities_id) VALUES ($1, $2) RETURNING id"
	for _, f := range payload.Facility {
		_, err := tx.Exec(query, idRSVP, f.Id)
		if err != nil {
			validate(err, "insertFacility", tx)
			return
		}
	}
	log.Println("Facilities inserted successfully.")
}

func validate(err error, message string, tx *sql.Tx)  {
	if err != nil{
		tx.Rollback()
		log.Println(err, "Transaction Rollback!")
	} else{
		log.Println("Successfully " + message + " data!")
	}
}

