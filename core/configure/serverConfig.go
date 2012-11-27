package configure

const (
	SERVERTYPE_GATE   = 1
	SERVERTYPE_SERVER = 2
	SERVERTYPE_MASTER = 3
)

type ServerConfig struct {
	Id   string
	Host string
	Port int
	Type int
}

func CreateServerConfig(id string, host string, port int, typ int) *ServerConfig {
	obj := new(ServerConfig)
	obj.Id = id
	obj.Host = host
	obj.Port = port
	obj.Type = typ
	return obj
}
