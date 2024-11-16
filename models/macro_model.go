package models

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Macro struct {
	Name    string    `yaml:"Name"`
	Actions []*Action `yaml:"Actions"`
}

func NewMacro(name string, actions []*Action) *Macro {
	m := &Macro{
		Name:    name,
		Actions: actions,
	}

	if actions == nil {
		m.Actions = make([]*Action, 0)
	}

	return m
}

func (m Macro) String() string {
	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
