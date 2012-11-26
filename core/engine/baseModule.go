package engine

import (
	"../packet"
)

type BaseModule interface {
	GetModuleName() string
	GetModuleCode() byte
	ExcuteCommand(NclientId int, cmd packet.IGozillaPacket)
}
