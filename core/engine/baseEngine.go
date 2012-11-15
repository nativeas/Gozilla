package engine

import (
	"../packet"
)

type Engine struct {
	Version string
	Modules map[byte]BaseModule
	players *PlayerCollection
}

func (e *Engine) Init(ver string) {
	e.Version = ver
	e.players = new(PlayerCollection)
	e.players.Init()
}

func (e *Engine) RegisterModule(m BaseModule) {
	if e.Modules == nil {
		e.Modules = make(map[byte]BaseModule)
	}
	e.Modules[m.GetModuleCode()] = m
}

func (e *Engine) excute() bool {
	p, cmd := e.players.Pump()
	if p == nil {
		return false
	}
	mod_code := cmd.GetMainCmd()
	mod, ok := e.Modules[mod_code]
	if ok {
		mod.ExcuteCommand(p, cmd)
		return true
	}
	return false
}

func (e *Engine) ExcuteCycle() {
	mark := e.excute()
	for mark {
		mark = e.excute()
	}
}

func (e *Engine) PushPacket(NclientId int, packet packet.IGozillaPacket) {
	e.players.PushPacket(NclientId, packet)
}
