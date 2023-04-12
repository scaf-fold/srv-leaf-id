package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scaf-fold/srv-leaf-id/controller"
	"github.com/scaf-fold/srv-leaf-id/global"
	"github.com/scaf-fold/tools/pkg/courier"
	"log"
	"path/filepath"
	"strings"
)

func init() {
	global.Config(nil)
}

func main() {
	g := gin.Default()
	courier.App(g, controller.Root, func(e *gin.Engine) {
		e.Run(":8081")
	})
}

func TT(path string) {
	var zkPath string
	paths := strings.Split(path, "/")
	size := len(paths)
	for index, node := range paths {
		log.Println("size", size, ", index", index, ", node:", node)
		if node == "" {
			continue
		}
		zkPath = zkPath + "/" + node

		if index < size-1 {
			//zkPath, _ = cr.zkClient.Create(zkPath, []byte(""), 0, acl)
			log.Println("index=", index, "zkPath =", zkPath)
		} else {
			//zkPath, _ = cr.zkClient.Create(zkPath, data, flags, acl)
			log.Println("index=", index, "zkPath =", zkPath)
		}
	}
}

func TT3(path string) {
	var nodePath string
	paths := strings.Split(path, "/")
	size := len(paths)
	for index, sub := range paths {
		log.Println("size", size, ", index", index, ", sub", sub)
		if sub == "" {
			continue
		}
		nodePath += fmt.Sprintf("%c%s", filepath.Separator, sub)
		if index < size-1 {
			log.Println("index=", index, "nodePath=", nodePath)
		} else {
			log.Println("index=", index, "nodePath=", nodePath)
		}
	}
}
