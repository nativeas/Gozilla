package engine

type engine struct {
	Version string
	Modules map[byte]BaseModule
	players PlayerCollection
}

func (e *engine) Init(ver string) {
	e.Version = ver
	e.players = new(PlayerCollection)
}

func (e *engine) RegisterModule(m *BaseModule) {
	if e.Modules == nil {
		e.Modules = make(map[byte]BaseModule)
	}
	e.Modules[m.GetModuleCode()] = m
}

func (e *engine) Excute() bool {

}
