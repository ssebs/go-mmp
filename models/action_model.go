package models

import (
	"log"

	"gopkg.in/yaml.v3"
)

type Action struct {
	FuncName  string `yaml:"FuncName"`
	FuncParam string `yaml:"FuncParam"`
}

func NewAction(funcName, funcParam string) Action {
	return Action{
		FuncName:  funcName,
		FuncParam: funcParam,
	}
}

func (a Action) String() string {
	data, err := yaml.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
