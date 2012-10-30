package main

import (
	"../module"
	"../socket"
	"log"
)

type sampleModule struct {
	module.Module
}

func (s *sampleModule) ModuleCMD() uint {
	return 1
}

func (s *sampleModule) Input(cmd *socket.TargetdCommand) {
	log.Println("sampleModule do cmd")
	log.Println(cmd.Command.ComandContent)
}

func main() {
	r := module.NewModuleRouter()
	var b module.IModule = new(sampleModule)

	r.RegisterModule(b)
	cmd := new(socket.TargetdCommand)
	cmd.Command = socket.NewSocketCommand(1, 2, "hello")
	r.RoutingCommand(cmd)
}
