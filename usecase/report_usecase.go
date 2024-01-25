package usecase

import (
	"booking-room/model"
	"booking-room/repository"
	"fmt"
)

type reportUsecase struct {
	reportRepository repository.ReportRepository
}

type ReportUsecase interface {
	GetReport() ([]model.EmployeeList, []model.RoomList, []model.FacilitiesList, []model.ReservationReport, []model.FacilityTotalReserved, []model.RoomTotalReserved, error)
}

func NewReportUsecase(reportRepository repository.ReportRepository) ReportUsecase {
	return &reportUsecase{reportRepository: reportRepository}
}
func (r *reportUsecase) GetReport() ([]model.EmployeeList, []model.RoomList, []model.FacilitiesList, []model.ReservationReport, []model.FacilityTotalReserved, []model.RoomTotalReserved, error) {
	employee, err := r.reportRepository.FindAll()
	room, err := r.reportRepository.FindAllRoom()
	facilities, err := r.reportRepository.FindAllFacilities()
	reserve, err := r.reportRepository.ReservationReport()
	totalFacility, err := r.reportRepository.FacilityTotalReserved()
	totalRoom, err := r.reportRepository.RoomTotalReserved()
	if err != nil {
		return nil, nil, nil, nil, nil, nil, fmt.Errorf("Problem with accesing Report Data")
	}
	return employee, room, facilities, reserve, totalFacility, totalRoom, nil
}
