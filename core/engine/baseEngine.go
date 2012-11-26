package engine

import (
	"../packet"
)

const (
	CORE_VER = "alpha1"
)

type Engine struct {
	Version string
	Modules map[byte]BaseModule
	players *PlayerCollection
}

func (e *Engine) Init() {
	e.Version = CORE_VER
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
		mod.ExcuteCommand(1, cmd)
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
