package models

import (
	"fmt"
	"log"
	"slices"

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

func NewDefaultMacro() *Macro {
	return &Macro{
		Name:    "MacroName",
		Actions: make([]*Action, 0),
	}
}

func (m *Macro) AddAction(a *Action) {
	m.Actions = append(m.Actions, a)
}

func (m *Macro) GetAction(idx int) (*Action, error) {
	if !m.isValidBoundsInActions(idx) {
		return nil, fmt.Errorf("idx out of bounds of Macro's actions")
	}
	return m.Actions[idx], nil
}

func (m *Macro) UpdateAction(idx int, updatedAction *Action) error {
	if !m.isValidBoundsInActions(idx) {
		return fmt.Errorf("idx out of bounds of Macro's actions")
	}

	m.Actions[idx] = updatedAction
	return nil
}

func (m *Macro) DeleteAction(idx int) error {
	if !m.isValidBoundsInActions(idx) {
		return fmt.Errorf("idx out of bounds of Macro's actions")
	}

	m.Actions = slices.Delete(m.Actions, idx, idx+1)
	return nil
}

// swap positions of idx1 and idx2
func (m *Macro) SwapActionPositions(idx1, idx2 int) error {
	if !m.isValidBoundsInActions(idx1) || !m.isValidBoundsInActions(idx2) {
		return fmt.Errorf("idx out of bounds of Macro's actions")
	}

	m.Actions[idx1], m.Actions[idx2] = m.Actions[idx2], m.Actions[idx1]
	return nil
}

// move srcIdx to where destIdx is, and shift over by 1
func (m *Macro) MoveActionPositions(srcIdx, destIdx int) error {
	if !m.isValidBoundsInActions(srcIdx) || !m.isValidBoundsInActions(destIdx) {
		return fmt.Errorf("idx out of bounds of Macro's actions")
	}
	if srcIdx == destIdx {
		return nil
	}

	item := m.Actions[srcIdx]
	// Remove the item at srcIdx by shifting left
	if srcIdx < destIdx {
		// Move item to a higher index
		copy(m.Actions[srcIdx:], m.Actions[srcIdx+1:destIdx+1])
	} else {
		// Move item to a lower index
		copy(m.Actions[destIdx+1:], m.Actions[destIdx:srcIdx])
	}

	m.Actions[destIdx] = item
	return nil
}

func (m *Macro) isValidBoundsInActions(idx int) bool {
	return idx <= len(m.Actions) && idx >= 0
}

func (m Macro) String() string {
	data, err := yaml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
