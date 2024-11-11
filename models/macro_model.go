package models

type Macro struct {
	Name    string   `yaml:"Name"`
	Actions []Action `yaml:"Actions"`
}

func NewMacro(name string, actions []Action) Macro {
	m := Macro{
		Name:    name,
		Actions: actions,
	}

	if actions == nil {
		m.Actions = make([]Action, 0)
	}

	return m
}
