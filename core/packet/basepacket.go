package packet

// 
//  定义了所有命令包的结构
//
type GozillaPacket struct {
	MainCmd byte
	SubCmd  byte
}

//
//  定义了所有命令包需要实现的方法
//
type IGozillaPacket interface {
	GetMainCmd() byte
	GetSubCmd() byte
	Init()
}

// 获取命令包的MainCmd
// 主要用于标识模块
func (g *GozillaPacket) GetMainCmd() byte {
	return g.MainCmd
}

// 获取命令包的SubCmd
// 主要用于标识实际的操作
func (g *GozillaPacket) GetSubCmd() byte {
	return g.SubCmd
}

//
//  初始化命令包数据
//	涵盖了对MainCMD、SubCMD以及其他一切的修改
//
func (g *GozillaPacket) Init() {

}

//login packet sample

// type LoginPacket struct {
// 	GozillaPacket
// }

// const (
// 	LOGIN_MOD    = 1
// 	PRELOGIN_CMD = 1
// )

// func (l *LoginPacket) Init() {
// 	l.MainCmd = LOGIN_MOD
// }

// type PreloginPacket struct {
// 	LoginPacket
// }

// func (p *PreloginPacket) Init() {
// 	p.LoginPacket.Init()
// 	p.SubCmd = PRELOGIN_CMD
// }
