package api

import (
	"github.com/it-chain/it-chain-Engine/gateway"
	"github.com/it-chain/it-chain-Engine/gateway/infra"
)

type MessageApi struct {
	grpcService gateway.GrpcService
}

func NewMessageApi(grpcService gateway.GrpcService) *MessageApi {
	return &MessageApi{
		grpcService: grpcService,
	}
}

//todo validation rule added example( check length of recipent)
func (c MessageApi) DeliverMessage(command gateway.MessageDeliverCommand) {

	//validation rule add
	hostService := infra.NewGrpcHostService()
	c.grpcService.SendMessages(command.Body, command.Protocol, command.Recipients...)
}