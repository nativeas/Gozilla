package login

import (
	"../../core/engine"
)

type Login_Module struct {
}

func (l *Login_Module) GetModuleName() {
	return "Login_Module"
}

func (l *Login_Module) GetModuleCode() {
	return 1
}
