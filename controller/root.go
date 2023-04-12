package controller

import (
	"github.com/scaf-fold/srv-leaf-id/controller/id"
	"github.com/scaf-fold/tools/pkg/courier"
)

var serviceName = "srv-leaf-id"

var Root = courier.NewGroup(serviceName)

func init() {
	V1 := Root.Group("/v1")
	{
		V1.Register(&id.IdGenerator{})
		V1.Register(&id.IdParser{})
	}
}
