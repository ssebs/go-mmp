package models

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// TODO: generate this from macro
// TODO: enum?
var actionFunctionList = []string{
	"Delay",
	"PressRelease",
	"Press",
	"Release",
	"SendText",
	"Shortcut",
	"Repeat",
	"ClickAt",
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

func NewDefaultAction() *Action {
	return &Action{
		FuncName:  "SendText",
		FuncParam: "",
	}
}

func (a Action) Validate() error {
	found := false
	for _, actionFuncName := range actionFunctionList {
		if a.FuncName == actionFuncName {
			found = true
		}
	}
	if !found {
		return fmt.Errorf("%s is not a valid Function in Actions", a.FuncName)
	}

	// TODO: validate FuncParam
	switch a.FuncName {
	case "Delay":
		if _, err := time.ParseDuration(a.FuncParam); err != nil {
			return fmt.Errorf("%s is not a valid go time.duration", a.FuncParam)
		}
	case "Shortcut":
		if !strings.Contains(a.FuncParam, "+") {
			return fmt.Errorf("%s does not have a + for a shortcut", a.FuncParam)
		}
	}

	return nil
}

func GetActionFunctions() []string {
	return actionFunctionList
}

func (a Action) String() string {
	data, err := yaml.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
