package login

import (
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
