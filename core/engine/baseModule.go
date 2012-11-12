package engine

import (
	"../packet"
)

type BaseModule interface {
	GetModuleName() string
	GetModuleCode() byte
	ExcuteCommand(p *PlayerObj, cmd packet.IGozillaPacket)
}
