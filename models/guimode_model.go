package models

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type GUIMode int

var _ yaml.Marshaler = (*GUIMode)(nil)
var _ yaml.Unmarshaler = (*GUIMode)(nil)

const (
	NOTSET GUIMode = iota // For use in comparing cli flags
	NORMAL                // Serial listener + GUI
	GUIOnly
	// CLIOnly
	TESTING
)

func GetGUIModesList() []string {
	return []string{
		"NORMAL",
		"GUIOnly",
	}
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
	case TESTING:
		return "TESTING"
	}
	return "NOTSET"
}

func (g *GUIMode) Set(m string) error {
	switch strings.ToUpper(m) {
	case "NORMAL":
		*g = NORMAL
		return nil
	case "GUIONLY":
		*g = GUIOnly
		return nil
	case "TESTING":
		*g = TESTING
		return nil
	}

	*g = NOTSET
	return fmt.Errorf("could not find mode %s", m)
}
func ParseGUIModeString(m string) (GUIMode, error) {
	gm := NOTSET
	switch strings.ToUpper(m) {
	case "NORMAL":
		gm = NORMAL
	case "GUIONLY":
		gm = GUIOnly
	case "TESTING":
		gm = TESTING
	}

	if gm != NOTSET {
		return gm, nil
	}

	return gm, fmt.Errorf("could not find mode %s", m)
}

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
