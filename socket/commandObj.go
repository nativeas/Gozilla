package socket

type SocketCommand struct {
	MainCMD       uint
	SubCMD        uint
	ComandContent string
}

func NewSocketCommand(m uint, s uint, c string) SocketCommand {
	obj := new(SocketCommand)
	obj.ComandContent = c
	obj.MainCMD = m
	obj.SubCMD = s
	return *obj
}
