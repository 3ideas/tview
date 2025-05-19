// Demo code for the DropDown primitive.
package main

import (
	"fmt"

	"github.com/3ideas/tview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := tview.NewApplication()
	var dropdown *tview.DropDown

	dropdown = tview.NewDropDown().
		SetLabel("Select an option (hit Enter): ").
		SetOptionsColor([]string{"First", "Second", "Third", "Fourth", "Fifth"},
			[]tcell.Color{tcell.ColorRed, tcell.ColorGreen, tcell.ColorBlue, tcell.ColorYellow, tcell.ColorCrimson},
			func(text string, index int) {
				// Show the selected option
				fmt.Printf("Selected: %s (index: %d)\n", text, index)
				app.QueueUpdateDraw(func() {
					dropdown.SetCurrentOption(index)
					fmt.Printf("Current option set to: %d\n", index)
				})
			})

	// Configure the dropdown list to show items properly
	dropdown.SetFieldWidth(20).
		SetDoneFunc(func(key tcell.Key) {
			fmt.Printf("Done with key: %v\n", key)
		})

	if err := app.SetRoot(dropdown, true).EnableMouse(true).Run(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
