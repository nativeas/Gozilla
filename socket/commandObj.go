package socket

type SocketCommand struct {
	MainCMD       uint   //主命令
	SubCMD        uint   //子命令
	ComandContent string //命令内容
}

func NewSocketCommand(m uint, s uint, c string) SocketCommand {
	obj := new(SocketCommand)
	obj.ComandContent = c
	obj.MainCMD = m
	obj.SubCMD = s
	return *obj
}

//这个结构用来从服务里面输出到业务逻辑
//
type TargetdCommand struct {
	TaretId int           //目标对象id，用来对应remoteObj的id
	Command SocketCommand //发送的数据包
}
