package utils

import (
	"flag"
)

func ParseFlag() (string, string) {
	var (
		id  = flag.String("server", "-1", "serverId")
		typ = flag.String("type", "master", "server type")
	)
	flag.Parse()
	return *id, *typ
}
