package services

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"

	"audiman/gu-project/recharger-ser/entity"
	msgbroker "audiman/gu-project/recharger-ser/messageBroker"
)

// RechargerService interface for recharge service
type RechargerService interface {
	InitRechargesListener() error
	CloseListener()
}

type service struct{}

var msgbkr msgbroker.MessageBroker

// NewRechargerService creates a new recharger service
func NewRechargerService(messageBroker msgbroker.MessageBroker) RechargerService {
	msgbkr = messageBroker
	return &service{}
}

func (*service) InitRechargesListener() error {
	err := msgbkr.Subscribe("addRecharge", handleMessages)
	if err != nil {
		errors.Wrap(err, "services.rechargerService.InitRechargesListener")
	}
	return nil
}

func (*service) CloseListener() {
	msgbkr.Close()
}

// refactor

func handleMessages(data []byte) {
	fmt.Println(string(data))

	rmsg := entity.RechargeMsg{}
	err := json.Unmarshal(data, &rmsg)
	if err != nil {
		log.Printf("%s : service.rechargeService.handleMessages", err.Error())
	}

	time.Sleep(time.Second * 5)
	rmsg.ResolvedAt = time.Now().UTC().Unix()

	body, err := json.Marshal(rmsg)
	if err != nil {
		log.Printf("%s : service.rechargeService.handleMessages", err.Error())
	}

	fmt.Println("SEND RESPONSE")
	msgbkr.PublishOnQueue(body, "rechargeFromRechargerService")
}
