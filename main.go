package main

import (
	"github.com/hyperpilotio/snap-plugin-collector-prometheus/prometheus"
	"github.com/jpra1113/snap-plugin-lib-go/v1/plugin"
	"google.golang.org/grpc"
)

const (
	pluginName     = "snap-plugin-collector-prometheus"
	pluginVersion  = 1
	maxMessageSize = 100 << 20
)

func main() {
	plugin.StartCollector(
		prometheus.New(),
		pluginName,
		pluginVersion,
		plugin.GRPCServerOptions(grpc.MaxMsgSize(maxMessageSize)),
		plugin.GRPCServerOptions(grpc.MaxSendMsgSize(maxMessageSize)),
		plugin.GRPCServerOptions(grpc.MaxRecvMsgSize(maxMessageSize)),
	)
}
