package modules

import "github.com/lab42/rocketship/config"

var Modules []*Module

type Module struct {
	Name string
	Exec func(config *config.Config) (string, error)
}

func AddModule(m *Module) {
	Modules = append(Modules, m)
}

func NewModule(name string, exec func(config *config.Config) (string, error)) *Module {
	return &Module{
		Name: name,
		Exec: exec,
	}
}
