package engine

import (
	"../packet"
)

type Engine struct {
	Version string
	Modules map[byte]BaseModule
	players PlayerCollection
}

func (e *engine) Init(ver string) {
	e.Version = ver
	e.players = new(PlayerCollection)
}

func (e *engine) RegisterModule(m *BaseModule) {
	if e.Modules == nil {
		e.Modules = make(map[byte]BaseModule)
	}
	e.Modules[m.GetModuleCode()] = m
}

func (e *engine) excute() bool {
	p, cmd := e.players.Pump()
	if p == nil {
		return false
	}
	mod_code := cmd.GetMainCmd()
	mod, ok = e.Modules[mod_code]
	if ok {
		mod.ExcuteCommand(p, cmd)
		return true
	}
	if ok != true {
		return false
	}
}

func (e *engine) ExcuteCycle() {
	mark := e.excute()
	for mark {
		mark = e.excute()
	}
}

func (e *Engine) PushPacket(NclientId int, packet packet.IGozillaPacket) {
	e.players.PushPacket(NclientId, packet)
}
