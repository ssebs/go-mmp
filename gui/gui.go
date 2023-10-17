package gui

import "fmt"

// GUISize
type GUISize struct {
	Width  int
	Height int
}

func (g *GUISize) GetSizeAsString() string {
	return fmt.Sprintf("%dx%d", g.Width, g.Height)
}

// GUI
type GUI struct {
	Size    GUISize
	Monitor int
}

// Create a new GUI, given a GUISize and a monitor #
func NewGUI(gs GUISize, monitor int) *GUI {
	return &GUI{Size: gs, Monitor: monitor}
}
