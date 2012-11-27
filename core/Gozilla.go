package core

import (
	conf "./configure"
	"./utils"
	"log"
)

const (
	LIFECYCLE_PREINIT = 1
)

type Gozilla struct {
	status       int // lifecycle status
	serverConfig *conf.ServerConfig
	// engine Base
}

func CreateGozilla() *Gozilla {
	goz := new(Gozilla)
	return goz
}

//检查环境
func (g *Gozilla) init() {
	log.Println("Gozilla Initing...")
	sid, stype := utils.ParseFlag()
	if stype == "master" {
		g.masterConfigure()
	} else {
		g.configureGozilla(sid)
	}
	log.Println("Gozilla Init Complete!")
}

//read master json , to start master gozilla
func (g *Gozilla) masterConfigure() {
	g.serverConfig = conf.CreateServerConfig("dev", "127.0.0.1", 8001, conf.SERVERTYPE_MASTER)
}

// read json define
func (g *Gozilla) configureGozilla(id string) {

}

//启动
func (g *Gozilla) Start() {
	g.init()

	g.daemon()
}

//停止
func (g *Gozilla) Stop() {

}

//全局等待
func (g *Gozilla) daemon() {

	for {

	}
}

func (g *Gozilla) LoadModule() {

}
