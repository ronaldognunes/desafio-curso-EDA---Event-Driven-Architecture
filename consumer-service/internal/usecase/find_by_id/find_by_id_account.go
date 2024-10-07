package find_by_id

import "consumer-service/internal/gateway"

type AccountOutPutDTO struct {
	AccountId      string
	AccountBalance float64
}

type FindByIdAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewFindByIdAccountUseCase(a gateway.AccountGateway) *FindByIdAccountUseCase {
	return &FindByIdAccountUseCase{
		AccountGateway: a,
	}
}

func (a *FindByIdAccountUseCase) Execute(id string) (*AccountOutPutDTO, error) {
	result := AccountOutPutDTO{}
	dados, err := a.AccountGateway.FindById(id)
	if err != nil {
		return nil, err
	}
	result.AccountId = dados.AccountId
	result.AccountBalance = dados.AccountBalance

	return &result, nil
}
