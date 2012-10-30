package module

import (
	"../socket"
	"log"
)

type ModuleRouter struct {
	modules map[uint]IModule
}

func NewModuleRouter() *ModuleRouter {
	var r = new(ModuleRouter)
	r.modules = make(map[uint]IModule)
	return r
}

// 注册模块
// 
//
func (m *ModuleRouter) RegisterModule(module IModule) {
	log.Println(module)
	var cmdtype uint = module.ModuleCMD()
	m.modules[cmdtype] = module
}

func (m *ModuleRouter) RoutingCommand(cmd *socket.TargetdCommand) {
	lm, ok := m.modules[cmd.Command.MainCMD]
	if ok {
		lm.Input(cmd)
	} else {
		log.Println("error!")
	}
}

func (m *ModuleRouter) Hello() {
	log.Println("Hello,module router")
}
