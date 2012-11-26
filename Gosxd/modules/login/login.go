package login

import (
	"../../../core/packet"
	"../../modules"
)

type Login_Module struct {
}

func (l *Login_Module) GetModuleName() string {
	return "Login_Module"
}

func (l *Login_Module) GetModuleCode() byte {
	return modules.MOD_LOGIN
}

func (l *Login_Module) ExcuteCommand(NclientId int, cmd packet.IGozillaPacket) {

}
