package usecase

import (
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/shared_model"
	"strings"
)

type TrxRsvUsecase interface {
	List(page, size int) ([]dto.TransactionDTO, shared_model.Paging, error)
	GetID(id string) (dto.TransactionDTO, error)
	GetEmployee(id string, page, size int) ([]dto.TransactionDTO, shared_model.Paging, error)
	PostReservation(payload dto.PayloadReservationDTO) (dto.TransactionDTO, error)
	UpdateStatus(payload dto.TransactionDTO) (dto.TransactionDTO, error)
	DeleteResv(id string) (string, error)
}

type trxRsvUsecase struct {
	trxRsvRepo repository.TrxRsvRepository
}

// DeleteResv implements TrxRsvUsecase.
func (t *trxRsvUsecase) DeleteResv(id string) (string, error) {
	return t.trxRsvRepo.DeleteResv(id)
}

// UpdateStatus implements TrxRsvUsecase.
func (t *trxRsvUsecase) UpdateStatus(payload dto.TransactionDTO) (dto.TransactionDTO, error) {
	acc := strings.ToUpper(payload.ApproveStatus)
	payload.ApproveStatus = acc
	if payload.ApproveNote == "" {
		payload.ApproveNote = "Viewed by GA"
	}

	return t.trxRsvRepo.UpdateStatus(payload)

}

// PostReservation implements TrxRsvUsecase.
func (t *trxRsvUsecase) PostReservation(payload dto.PayloadReservationDTO) (dto.TransactionDTO, error) {
	idRSVP, err := t.trxRsvRepo.PostReservation(payload)
	if err != nil {
		return dto.TransactionDTO{}, err
	}

	return t.trxRsvRepo.GetID(idRSVP)
}

// GetEmployee implements TrxRsvUsecase.
func (t *trxRsvUsecase) GetEmployee(id string, page, size int) ([]dto.TransactionDTO, shared_model.Paging, error) {
	p, s := noneQuery(page, size)
	return t.trxRsvRepo.GetEmployee(id, p, s)
}

// Get implements TrxRsvUsecase.
func (t *trxRsvUsecase) GetID(id string) (dto.TransactionDTO, error) {
	return t.trxRsvRepo.GetID(id)
}

// List implements TrxRsvUsecase.
func (t *trxRsvUsecase) List(page, size int) ([]dto.TransactionDTO, shared_model.Paging, error) {
	p, s := noneQuery(page, size)
	return t.trxRsvRepo.List(p, s)
}

func NewTrxRsvUseCase(trxRsvRepo repository.TrxRsvRepository) TrxRsvUsecase {
	return &trxRsvUsecase{
		trxRsvRepo: trxRsvRepo,
	}
}

func noneQuery(page, size int) (int, int) {
	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 5
	}

	return page, size
}
