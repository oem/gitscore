package ansi

import ui "github.com/gizak/termui"

func Draw() error {
	err := ui.Init()
	return err
}
