package frpmanager

import (
	"github.com/admpub/nging/v5/application/registry/navigate"

	"github.com/nging-plugins/frpmanager/application/handler"
)

var LeftNavigate = handler.LeftNavigate

func RegisterNavigate(nc *navigate.Collection) {
	nc.Backend.AddLeftItems(-1, LeftNavigate)
}
