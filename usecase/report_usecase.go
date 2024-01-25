package usecase

import (
	"booking-room/model"
	"booking-room/repository"
	"fmt"
	"time"
)

type reportUsecase struct {
	reportRepository repository.ReportRepository
}

type ReportUsecase interface {
	GetReport(start, end time.Time) ([]model.EmployeeList, []model.RoomList, []model.FacilitiesList, []model.ReservationReport, []model.FacilityTotalReserved, []model.RoomTotalReserved, error)
}

func NewReportUsecase(reportRepository repository.ReportRepository) ReportUsecase {
	return &reportUsecase{reportRepository: reportRepository}
}
func (r *reportUsecase) GetReport(start, end time.Time) ([]model.EmployeeList, []model.RoomList, []model.FacilitiesList, []model.ReservationReport, []model.FacilityTotalReserved, []model.RoomTotalReserved, error) {
	employee, err := r.reportRepository.FindAll()
	room, err := r.reportRepository.FindAllRoom()
	facilities, err := r.reportRepository.FindAllFacilities()
	reserve, err := r.reportRepository.ReservationReport(start, end)
	totalFacility, err := r.reportRepository.FacilityTotalReserved(start, end)
	totalRoom, err := r.reportRepository.RoomTotalReserved(start, end)
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Problem with accesing Report Data")
	}
	return employee, room, facilities, reserve, totalFacility, totalRoom, nil
}
