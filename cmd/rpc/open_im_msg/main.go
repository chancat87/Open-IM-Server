package main

import (
	"Open_IM/internal/rpc/msg"
	"Open_IM/pkg/common/config"
	promePkg "Open_IM/pkg/common/prometheus"
	"flag"
	"fmt"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImMessagePort
	rpcPort := flag.Int("port", defaultPorts[0], "rpc listening port")
	prometheusPort := flag.Int("promethus-port", config.Config.Prometheus.MessagePrometheusPort[0], "msgPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start msg rpc server, port: ", *rpcPort)
	rpcServer := msg.NewRpcChatServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
