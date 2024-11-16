package models

import (
	"log"

	"gopkg.in/yaml.v3"
)

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

func (c ConfigM) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
