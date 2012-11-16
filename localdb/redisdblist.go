package localdb

const (
	REDISDBINDEX_PLAYER  = 1
	REDISDBINDEX_ACCOUNT = 2
	REDISDBINDEX_MAX     = 3
)

var RedisDBHashKeyArray = []string{
	"dbplayer",
	"dbaccount",
}
