package configure

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

const (
	jsonStram = `{
					"development":
					{
						"id": "master-server-1", "host": "127.0.0.1", "port": 3005
					},
				  	"production":
				  	{
				  		"id": "master-server-1", "host": "pomelo3.server.163.org", "port": 3005
				  	}
				}`
)

type config struct {
	Production  element
	Development element
}

type element struct {
	Id   string
	Host string
	Port int
}

func FillConfig(jsonStraem string) config {
	var jsontype config
	// json.Unmarshal(jsonStram, &jsontype)
	log.Printf("Results: %v\n", jsontype)
	dec := json.NewDecoder(strings.NewReader(jsonStraem))
	for {
		var m config
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		log.Println(m)
	}
	return jsontype
}

const (
	servers = `{
				    "development":{
				        "connector":[
				             {"id":"connector-server-1", "host":"127.0.0.1", "port":4050, "wsPort":3050},
				             {"id":"connector-server-2", "host":"127.0.0.1", "port":4051, "wsPort":3051},
				             {"id":"connector-server-3", "host":"127.0.0.1", "port":4052, "wsPort":3052}
				         ],
				        "chat":[
				             {"id":"chat-server-1", "host":"127.0.0.1", "port":6050},
				             {"id":"chat-server-2", "host":"127.0.0.1", "port":6051},
				             {"id":"chat-server-3", "host":"127.0.0.1", "port":6052}
				        ],
				        "gate":[
					     {"id": "gate-server-1", "host": "127.0.0.1", "wsPort": 3014}
					]
				    },
				    "production":{
				       "connector":[
				             {"id":"connector-server-1", "host":"127.0.0.1", "port":4050, "wsPort":3050},
				             {"id":"connector-server-2", "host":"127.0.0.1", "port":4051, "wsPort":3051},
				             {"id":"connector-server-3", "host":"127.0.0.1", "port":4052, "wsPort":3052}
				         ],
				        "chat":[
				             {"id":"chat-server-1", "host":"127.0.0.1", "port":6050},
				             {"id":"chat-server-2", "host":"127.0.0.1", "port":6051},
				             {"id":"chat-server-3", "host":"127.0.0.1", "port":6052}
				        ],
				        "gate":[
					     {"id": "gate-server-1", "host": "127.0.0.1", "wsPort": 3014}
					]
				  }
				}`
)
