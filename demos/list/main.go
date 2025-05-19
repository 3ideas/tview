// Demo code for the List primitive.
package main

import (
	"github.com/3ideas/tview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItemWithStyle("List item 1", "Some explanatory text", 'a', nil,
			tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true)).
		AddItemWithStyle("List item 2", "Some explanatory text", 'b', nil,
			tcell.StyleDefault.Foreground(tcell.ColorGreen).Italic(true)).
		AddItemWithStyle("List item 3", "Some explanatory text", 'c', nil,
			tcell.StyleDefault.Foreground(tcell.ColorBlue).Underline(true)).
		AddItemWithStyle("List item 4", "Some explanatory text", 'd', nil,
			tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true).Italic(true)).
		AddItemWithStyle("Quit", "Press to exit", 'q', func() {
			app.Stop()
		}, tcell.StyleDefault.Foreground(tcell.ColorCrimson).Bold(true).Underline(true))

	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
