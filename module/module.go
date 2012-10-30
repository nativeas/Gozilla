package module

import (
	"../socket"
	"log"
)

type IModule interface {
	ModuleCMD() uint
	Input(cmd *socket.TargetdCommand)
}

type Module struct {
}

func (m *Module) ModuleCMD() uint {
	return 0
}

func (m *Module) Input(cmd *socket.TargetdCommand) {
	log.Println("Do something")
}
