package controller

import (
	"booking-room/delivery/middleware"
	"booking-room/shared/common"
	"booking-room/usecase"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type ReportController struct {
	reportUsecase usecase.ReportUsecase
	middleware    *middleware.Middleware
	rp            *gin.RouterGroup
}

func NewReportController(reportUsecase usecase.ReportUsecase, middleware *middleware.Middleware, rp *gin.RouterGroup) *ReportController {
	return &ReportController{
		reportUsecase: reportUsecase,
		middleware:    middleware,
		rp:            rp,
	}
}

func (r *ReportController) Route() {
	r.rp.GET("/download", r.GetReport)
}

func (r *ReportController) GetReport(c *gin.Context) {
	claims := r.middleware.GetUser(c)
	if ok := common.AuthorizationAdmin(claims); !ok {
		log.Println("Authorization failed because user is not admin")
		common.SendErrorResponse(c, http.StatusForbidden, "Forbidden")
		return
	}

	format := "2006-01-02"
	now := time.Now()
	start, _ := time.Parse(format, c.Query("start"))
	end, _ := time.Parse(format, c.Query("end"))

	//set default query, -1 month and +1 month
	if start.IsZero() {
		start = now.AddDate(0, -1, 0)
	}
	if end.IsZero() {
		end = now.AddDate(0, 1, 0)
	}
	employee, room, facilities, reserve, totalFacility, totalRoom, err := r.reportUsecase.GetReport(start, end)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//generate excel file
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("DataSheet")

	sheet.AddRow()
	title := sheet.AddRow()
	title.AddCell().SetString("Report Booking Room")
	title.AddCell().SetDateTime(start)
	title.AddCell().SetDateTime(end)
	sheet.AddRow()

	//create header employee
	headerEmployee := sheet.AddRow()
	headerEmployee.AddCell().SetString("Name")
	headerEmployee.AddCell().SetString("Email")
	headerEmployee.AddCell().SetString("Division")
	headerEmployee.AddCell().SetString("Position")
	headerEmployee.AddCell().SetString("Role")
	headerEmployee.AddCell().SetString("Contact")

	//write employee data
	for _, data := range employee {
		row := sheet.AddRow()
		row.AddCell().SetString(data.Name)
		row.AddCell().SetString(data.Email)
		row.AddCell().SetString(data.Division)
		row.AddCell().SetString(data.Position)
		row.AddCell().SetString(data.Role)
		row.AddCell().SetString(data.Contact)
	}

	sheet.AddRow()

	//create header room
	headerRoom := sheet.AddRow()
	headerRoom.AddCell().SetString("Code Room")
	headerRoom.AddCell().SetString("Room Type")
	headerRoom.AddCell().SetString("Capacity")
	headerRoom.AddCell().SetString("Facility")

	//write room data
	for _, data := range room {
		row := sheet.AddRow()
		row.AddCell().SetString(data.CodeRoom)
		row.AddCell().SetString(data.TypeRoom)
		row.AddCell().SetInt(data.Capacity)
		row.AddCell().SetString(data.Facilities)
	}

	sheet.AddRow()

	//create header facilities
	headerFacilities := sheet.AddRow()
	headerFacilities.AddCell().SetString("Code Facilities")
	headerFacilities.AddCell().SetString("Facilities Type")
	headerFacilities.AddCell().SetString("Status")

	//write facilities data
	for _, data := range facilities {
		row := sheet.AddRow()
		row.AddCell().SetString(data.CodeName)
		row.AddCell().SetString(data.FacilityType)
		row.AddCell().SetString(data.Status)
	}

	sheet.AddRow()

	//create header reservation
	headerReservation := sheet.AddRow()
	headerReservation.AddCell().SetString("Reservation ID")
	headerReservation.AddCell().SetString("Employee Name")
	headerReservation.AddCell().SetString("Code Room")
	headerReservation.AddCell().SetString("Start Date")
	headerReservation.AddCell().SetString("End Date")
	headerReservation.AddCell().SetString("Notes")
	headerReservation.AddCell().SetString("Approval Status")
	headerReservation.AddCell().SetString("Approval Note")

	//write reservation data
	for _, data := range reserve {
		row := sheet.AddRow()
		row.AddCell().SetString(data.ReservationId)
		row.AddCell().SetString(data.EmployeeName)
		row.AddCell().SetString(data.CodeRoom)
		row.AddCell().SetDateTime(data.StartDate)
		row.AddCell().SetDateTime(data.EndDate)
		row.AddCell().SetString(data.Note)
		row.AddCell().SetString(data.ApproveStatus)
		row.AddCell().SetString(data.ApproveNote)
	}

	sheet.AddRow()

	//create header total facility booked
	headerTotalFacility := sheet.AddRow()
	headerTotalFacility.AddCell().SetString("Facility Type")
	headerTotalFacility.AddCell().SetString("Total")

	//write total facility booked data
	for _, data := range totalFacility {
		row := sheet.AddRow()
		row.AddCell().SetString(data.FacilityType)
		row.AddCell().SetInt(data.Total)
	}

	sheet.AddRow()

	//create header total room booked
	headerTotalRoom := sheet.AddRow()
	headerTotalRoom.AddCell().SetString("Room Type")
	headerTotalRoom.AddCell().SetString("Total")

	//write total room booked data
	for _, data := range totalRoom {
		row := sheet.AddRow()
		row.AddCell().SetString(data.RoomType)
		row.AddCell().SetInt(data.Total)
	}

	fileName := "Report.xlsx"
	//Save file to excel
	err = file.Save(fileName)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Excel file created successfully.")
	}

	//set http header
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	//send file
	c.File(fileName)

}
