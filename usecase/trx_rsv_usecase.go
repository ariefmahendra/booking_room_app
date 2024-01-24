package usecase

import (
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/shared_model"
)

type TrxRsvUsecase interface {
	List(page,size int) ([]dto.TransactionDTO,  shared_model.Paging, error)
	GetID(id string) (dto.TransactionDTO, error)
	GetEmployee(id string, page,size int) ([]dto.TransactionDTO, shared_model.Paging, error)
}

type trxRsvUsecase struct {
	trxRsvRepo repository.TrxRsvRepository
}

// GetEmployee implements TrxRsvUsecase.
func (t *trxRsvUsecase) GetEmployee(id string, page,size int) ([]dto.TransactionDTO, shared_model.Paging, error) {
	p,s :=noneQuery(page,size)
	return t.trxRsvRepo.GetEmployee(id, p, s)
}

// Get implements TrxRsvUsecase.
func (t *trxRsvUsecase) GetID(id string) (dto.TransactionDTO, error) {
	return t.trxRsvRepo.GetID(id)
}

// List implements TrxRsvUsecase.
func (t *trxRsvUsecase) List(page,size int) ([]dto.TransactionDTO, shared_model.Paging, error) {
	p,s :=noneQuery(page,size)
	return t.trxRsvRepo.List(p,s)
}

func NewTrxRsvUseCase(trxRsvRepo repository.TrxRsvRepository) TrxRsvUsecase {
	return &trxRsvUsecase{
		trxRsvRepo: trxRsvRepo,
	}
}

func noneQuery(page,size int) (int, int) {
	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 5
	}

	return page, size
}
