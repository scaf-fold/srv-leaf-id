package global

import (
	"fmt"
	"github.com/scaf-fold/common/distribute_id"
	"github.com/scaf-fold/common/util"
	"github.com/scaf-fold/common/zk"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type LeafConfig struct {
	ServiceName string   `yaml:"serviceName"`
	Port        int64    `yaml:"port"`
	Cluster     []string `yaml:"cluster"`
	CachePath   string   `yaml:"cachePath"`
	StartTime   string   `yaml:"startTime"`
}

var timeFormat = "2006-01-02 15:04:05"

var ID *distribute_id.DistributeIdImpl
var StartTime time.Time

func Config(conf *string) {
	validPath := ""
	if conf == nil {
		validPath = "../conf/local.yaml"
		_, err := os.Stat(validPath)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := os.Stat(*conf)
		if err != nil {
			panic(err)
		}
		validPath = *conf
	}
	d, err := os.ReadFile(validPath)
	if err != nil {
		panic(err)
	}
	c := &LeafConfig{}
	err = yaml.Unmarshal(d, c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
	start, err := time.Parse(timeFormat, c.StartTime)
	if err != nil {
		panic(err)
	}
	ip := make(chan string)
	go util.Ip(ip)
	zkConf := &zk.Configuration{
		ConnCluster: c.Cluster,
		RootNode:    c.ServiceName,
		CachePath:   c.CachePath,
		EndPoint: &zk.Node{
			Ip:   <-ip,
			Port: c.Port,
		},
	}
	StartTime = start
	ID = distribute_id.NewDistributeIdImpl(&start, zkConf)
	ID.Start()
}
