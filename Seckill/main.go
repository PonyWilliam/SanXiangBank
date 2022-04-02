package main

import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/PonyWilliam/tarsSeckill/tars-protocol/SanXiangBank"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(PublishImp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("PublishImp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(SanXiangBank.Publish)
	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".PublishObj")

	// Run application
	tars.Run()
}
