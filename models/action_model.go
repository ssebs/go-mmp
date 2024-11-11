package models

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
