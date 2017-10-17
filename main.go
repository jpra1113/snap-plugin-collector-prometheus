package main

import (
	"github.com/hyperpilotio/snap-plugin-collector-prometheus/prometheus"
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

const (
	pluginName    = "snap-plugin-collector-prometheus"
	pluginVersion = 1
)

func main() {
	plugin.StartCollector(prometheus.New(), pluginName, pluginVersion)
}
