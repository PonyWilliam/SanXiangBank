package main

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars"

	"github.com/PonyWilliam/tarsSeckill/tars-protocol/SanXiangBank"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("SanXiangBank.Seckill.PublishObj@tcp -h 127.0.0.1 -p 10015 -t 60000")
	app := new(SanXiangBank.Publish)
	comm.StringToProxy(obj, app)
}
