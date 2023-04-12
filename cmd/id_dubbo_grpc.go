package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/scaf-fold/srv-leaf-id/global"
	"github.com/scaf-fold/srv-leaf-id/pkg/id_rpc"
	"log"
)

func init() {
	global.Config(nil)
}

func init() {
	config.SetProviderService(&id_rpc.IdGenerator{})
}

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	log.Println("id rpc provider started")
	select {}
}
