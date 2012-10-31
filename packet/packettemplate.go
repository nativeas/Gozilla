package packet

type I interface {
}

type STPacketBase struct {
	BTMainCmd byte
	BTSubCmd  byte
}

type STLoginBase struct {
	STBase STPacketBase
}

func (s *STLoginBase) InitMainCmd() {
	s.STBase.BTMainCmd = 0x01
}

//0x0101	预登陆
type STPreLogin struct {
	STCmd STLoginBase
}

func (s *STPreLogin) InitSubCmd() {
	s.STCmd.InitMainCmd()
	s.STCmd.STBase.BTSubCmd = 0x01
}

//0x0102	预登陆返回
type STPreLoginRet struct {
	STCmd STLoginBase
}

func (s *STPreLoginRet) InitSubCmd() {
	s.STCmd.InitMainCmd()
	s.STCmd.STBase.BTSubCmd = 0x02
}
