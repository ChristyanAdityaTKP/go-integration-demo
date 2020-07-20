package usecase

import (
	"time"

	"github.com/dh258/go-integration-demo/domain"
	"github.com/dh258/go-integration-demo/utils"
)

var (
	AddressUsecase addressUsecaseInterface = &addressUsecase{}
)

type addressUsecase struct{}

type addressUsecaseInterface interface {
	CreateAddress(*domain.Address) (*domain.Address, utils.MessageErr)
	GetAllAddresses() ([]*domain.Address, utils.MessageErr)
	GetByID(addressID int64) (*domain.Address, utils.MessageErr)
}

func (a *addressUsecase) CreateAddress(address *domain.Address) (*domain.Address, utils.MessageErr) {
	err := address.Validate()
	if err != nil {
		return nil, err
	}

	address.CreatedAt = time.Now()
	address.UpdatedAt = time.Now()

	result, err := domain.AddressRepo.Create(address)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *addressUsecase) GetAllAddresses() ([]*domain.Address, utils.MessageErr) {
	result, err := domain.AddressRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *addressUsecase) GetByID(addressID int64) (*domain.Address, utils.MessageErr) {
	result, err := domain.AddressRepo.Get(addressID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
