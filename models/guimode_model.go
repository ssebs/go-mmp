package models

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type GUIMode int

// var _ yaml.Marshaler = (*GUIMode)(nil)
var _ yaml.Unmarshaler = (*GUIMode)(nil)

const (
	NOTSET GUIMode = iota // For use in comparing cli flags
	NORMAL                // Serial listener + GUI
	GUIOnly
	// CLIOnly
	TESTING
)

/* YAML Marshaller implementations */
func (g *GUIMode) UnmarshalYAML(node *yaml.Node) error {
	var modeStr string
	var modeInt int

	// Check if the node contains an integer
	if err := node.Decode(&modeInt); err == nil {
		*g = GUIMode(modeInt)
		return nil
	}

	// If not an integer, attempt to decode as a string
	if err := node.Decode(&modeStr); err != nil {
		return err
	}

	// Use the Set method to handle string-based GUIMode values
	if err := g.Set(modeStr); err != nil {
		return fmt.Errorf("failed to unmarshal GUIMode: %w", err)
	}

	return nil
}

func (g GUIMode) MarshalYAML() (interface{}, error) {
	return g.Type(), nil
}

/* GUIMode pflag.Value implementation */
func (g *GUIMode) String() string {
	return g.Type()
}
func (g *GUIMode) Type() string {
	switch *g {
	case NORMAL:
		return "NORMAL"
	case GUIOnly:
		return "GUIOnly"
	}
	return ""
}
func (g *GUIMode) Set(m string) error {
	switch strings.ToUpper(m) {
	case "NORMAL":
		*g = NORMAL
		return nil
	case "GUIONLY":
		*g = GUIOnly
		return nil
	}

	return fmt.Errorf("could not find mode %s", m)
}
