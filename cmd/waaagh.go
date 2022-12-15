package main

import (
	"fmt"
	"os"

	"github.com/gateway-fm/scriptorium/logger"

	"github.com/gateway-fm/warp_external/cmd/infra_summoner"
	"github.com/gateway-fm/warp_external/cmd/mkdir"
	"github.com/gateway-fm/warp_external/cmd/proxy_summoner"
)

func main() {
	logger.SetLoggerMode("local")
	summoner := proxy_summoner.Cmd()
	summoner.AddCommand(proxy_summoner.Cmd())
	summoner.AddCommand(infra_summoner.Cmd())
	summoner.AddCommand(mkdir.Cmd())
	if err := summoner.Execute(); err != nil {
		logger.Log().Error(fmt.Errorf("generate command failed: %w", err).Error())
		os.Exit(1)
	}
}
