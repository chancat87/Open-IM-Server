package main

import (
	"Open_IM/internal/rpc/organization"
	"Open_IM/pkg/common/config"
	promePkg "Open_IM/pkg/common/prometheus"
	"flag"
	"fmt"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImOrganizationPort
	rpcPort := flag.Int("port", defaultPorts[0], "get RpcOrganizationPort from cmd,default 11200 as port")
	prometheusPort := flag.Int("promethus-port", config.Config.Prometheus.OrganizationPrometheusPort[0], "organizationPrometheusPort default listen port")
	flag.Parse()
	fmt.Println("start organization rpc server, port: ", *rpcPort)
	rpcServer := organization.NewServer(*rpcPort)
	go func() {
		err := promePkg.StartPromeSrv(*prometheusPort)
		if err != nil {
			panic(err)
		}
	}()
	rpcServer.Run()
}
