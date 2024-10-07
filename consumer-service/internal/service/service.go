package service

import (
	"consumer-service/internal/entity"
	"consumer-service/internal/gateway"
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

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

func (s *ConsumeService) ExecuteConsumerKafka() {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"kafka:29092"}
	//brokers := []string{"localhost:9092"}

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic("Erro no consumer")
	}

	parttition, err := consumer.ConsumePartition("balances", 0, sarama.OffsetOldest)
	if err != nil {
		panic("Erro na partição")
	}

	messages := parttition.Messages()

	for {
		select {
		case msg := <-messages:
			fmt.Println(string(msg.Value))
			var mesasgeKafka MensageKafka
			err = json.Unmarshal(msg.Value, &mesasgeKafka)
			if err != nil {
				println("Erro ao efetuar parse da mensagem do kafka")
			}
			accountFrom := entity.NewAccount(mesasgeKafka.Payload.AccountIDFrom, float64(mesasgeKafka.Payload.BalanceAccountIDFrom))
			accountTo := entity.NewAccount(mesasgeKafka.Payload.AccountIDTo, float64(mesasgeKafka.Payload.BalanceAccountIDTo))

			err = s.gateway.Update2(accountFrom)
			if err != nil {
				fmt.Println("erro ao gravar primeiro valor", err)
			}
			err = s.gateway.Update2(accountTo)
			if err != nil {
				fmt.Println("erro ao gravar segundo valor", err)
			}
		}
	}

}
