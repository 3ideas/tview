// Demo code for the List primitive.
package main

import (
	"github.com/3ideas/tview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	// Set different colors and styles for each item
	list.SetItemColor(0, tcell.ColorRed)
	list.SetItemStyle(0, tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true))

	list.SetItemColor(1, tcell.ColorGreen)
	list.SetItemStyle(1, tcell.StyleDefault.Foreground(tcell.ColorGreen).Italic(true))

	list.SetItemColor(2, tcell.ColorBlue)
	list.SetItemStyle(2, tcell.StyleDefault.Foreground(tcell.ColorBlue).Underline(true))

	list.SetItemColor(3, tcell.ColorYellow)
	list.SetItemStyle(3, tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true).Italic(true))

	list.SetItemColor(4, tcell.ColorCrimson)
	list.SetItemStyle(4, tcell.StyleDefault.Foreground(tcell.ColorCrimson).Bold(true).Underline(true))

	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
