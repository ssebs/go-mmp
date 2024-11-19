package models

import (
	"fmt"
	"log"
	"slices"

	"gopkg.in/yaml.v3"
)

/*
- `SaveConfig(destinationFullPath string)`
- `LoadConfig(sourceFullPath) *Config`
- `AddMacro(newMacro Macro)`
- `DeleteMacro(idx int)`
- `UpdateMacro(idx int, updatedMacro Macro)`
- `GetMacro(idx int)`

also functionality from old Config
*/
type ConfigM struct {
	*Metadata `yaml:"Metadata"`
	Macros    []*Macro `yaml:"Macros"`
}

func NewConfigM(meta *Metadata, macros []*Macro) *ConfigM {
	c := &ConfigM{
		Metadata: meta,
		Macros:   macros,
	}

	if c.Macros == nil {
		c.Macros = make([]*Macro, 0)
	}

	return c
}

func (c *ConfigM) AddMacro(m *Macro) {
	c.Macros = append(c.Macros, m)
}

func (c *ConfigM) GetMacro(idx int) (*Macro, error) {
	if !c.isValidBoundsInMacros(idx) {
		return nil, fmt.Errorf("idx out of bounds of Macros")
	}

	return c.Macros[idx], nil
}

func (c *ConfigM) UpdateMacro(idx int, updatedMacro *Macro) error {
	if !c.isValidBoundsInMacros(idx) {
		return fmt.Errorf("idx out of bounds of Macros")
	}

	c.Macros[idx] = updatedMacro
	return nil
}

func (c *ConfigM) DeleteMacro(idx int) error {
	if !c.isValidBoundsInMacros(idx) {
		return fmt.Errorf("idx out of bounds of Macros")
	}
	c.Macros = slices.Delete(c.Macros, idx, idx+1)
	return nil
}

func (c *ConfigM) SwapMacroPositions(idx1, idx2 int) error {
	if !c.isValidBoundsInMacros(idx1) || !c.isValidBoundsInMacros(idx2) {
		return fmt.Errorf("idx out of bounds of Macro's actions")
	}

	c.Macros[idx1], c.Macros[idx2] = c.Macros[idx2], c.Macros[idx1]
	return nil
}

func (c *ConfigM) isValidBoundsInMacros(idx int) bool {
	return idx <= len(c.Macros) && idx >= 0
}

func (c ConfigM) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
