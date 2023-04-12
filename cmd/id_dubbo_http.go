package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/srv-leaf-id/controller"
	"github.com/scaf-fold/srv-leaf-id/global"
	"github.com/scaf-fold/srv-leaf-id/pkg/id_rpc"
	"github.com/scaf-fold/tools/pkg/courier"
	"log"
)

func init() {
	global.Config(nil)
	log.Println("id service initial started")
}

func init() {
	config.SetProviderService(&id_rpc.IdGenerator{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	log.Println("id rpc provider started")
}

func main() {
	g := gin.Default()
	courier.App(g, controller.Root, func(e *gin.Engine) {
		e.Run(":8081")
	})
}
