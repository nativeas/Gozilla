package localdb

/*[db总列表]*/
const (
	REDISDBINDEX_PLAYER  = 0 //playermodeldb
	REDISDBINDEX_ACCOUNT = 1
	REDISDBINDEX_MAX     = 2
)

var RedisDBHashKeyArray = []string{
	"dbplayer",
	"dbaccount",
}
