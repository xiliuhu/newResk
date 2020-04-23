package main

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/ztaoing/newResk/services"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":18082")
	if err != nil {
		log.Panic(err)
	}
	in := services.RedEnvelopeSendDTO{
		Amount:       decimal.NewFromFloat(1),
		UserId:       "",
		Username:     "测试用户",
		EnvelopeType: services.GeneralEnvelope,
		Quantity:     2,
		Blessing:     "",
	}
	out := services.RedEnvelopeActivity{}
	err = client.Call("EnvelopeRpc.Send", in, &out)
	if err != nil {
		log.Panic(err)
	}
	log.Infof("%+v", out)

}
