package models

import (
	"log"

	"gopkg.in/yaml.v3"
)

// TODO: generate this from macro
var actionFunctionList = []string{
	"Delay",
	"PressRelease",
	"Press",
	"Release",
	"SendText",
	"Shortcut",
	"Repeat",
}

// Action describes an "action" to take. e.g. "Shortcut" => "CTRL+Z"
// Should be used within a Macro
type Action struct {
	FuncName  string `yaml:"FuncName"`
	FuncParam string `yaml:"FuncParam"`
}

func NewAction(funcName, funcParam string) *Action {
	return &Action{
		FuncName:  funcName,
		FuncParam: funcParam,
	}
}

func GetActionFunctions() []string {
	return actionFunctionList
}

func (a *Action) String() string {
	data, err := yaml.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
