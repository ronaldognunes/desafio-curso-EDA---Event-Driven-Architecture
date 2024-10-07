package consumer_service

import "consumer-service/internal/gateway"

type ConsumeService struct {
	gateway gateway.AccountGateway
}

type MensageKafka struct {
	Name    string `json:"Name"`
	Payload struct {
		AccountIDFrom        string `json:"account_id_from"`
		AccountIDTo          string `json:"account_id_to"`
		BalanceAccountIDFrom int    `json:"balance_account_id_from"`
		BalanceAccountIDTo   int    `json:"balance_account_id_to"`
	} `json:"Payload"`
}

func NewConsumeService(a gateway.AccountGateway) *ConsumeService {
	return &ConsumeService{
		gateway: a,
	}
}


