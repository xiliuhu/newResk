package gorpc

import (
	"github.com/ztaoing/newResk/services"
)

//用rpc实现法红包
type EnvelopeRpc struct {
}

//发红包
func (e EnvelopeRpc) Send(in services.RedEnvelopeSendDTO, out *services.RedEnvelopeActivity) error {
	s := services.GetRedEnvelopeService()
	activity, err := s.Send(in)
	activity.CopyTo(out)
	return err
}

//抢红包
func (e EnvelopeRpc) Receive(in services.RedEnvelopeReceiveDTO, out *services.RedEnvelopeItemDTO) error {
	s := services.GetRedEnvelopeService()
	item, err := s.Recive(in)
	item.CopyTo(out)
	return err
}
