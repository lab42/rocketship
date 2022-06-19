package modules

import "github.com/lab42/rocketship/config"

var Modules = make(map[string]*Module)

type Module struct {
	Name string
	Exec func(config *config.Config) (string, error)
}

func AddModule(module *Module) {
	Modules[module.Name] = module
}

func NewModule(name string, exec func(config *config.Config) (string, error)) *Module {
	return &Module{
		Name: name,
		Exec: exec,
	}
}
