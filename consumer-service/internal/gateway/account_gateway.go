package gateway

import "consumer-service/internal/entity"

type AccountGateway interface {
	Create(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
	Update(account *entity.Account) error
	Update2(account *entity.Account) error
}
