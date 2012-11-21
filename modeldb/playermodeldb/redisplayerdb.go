package playermodeldb

/*[玩家redisdb结构]*/
type PlayerObjDB struct {
	NOnlyID     int    //唯一ID
	StrCretName string //角色名
	NGameGold   int    //游戏币数
	NRmbGold    int    //充值币数
	NGiftGold   int    //礼金数
}

var PlayerObjField = []string{
	"onlyid",
	"cretname",
	"gamegold",
	"rmbgold",
	"giftgold",
}
