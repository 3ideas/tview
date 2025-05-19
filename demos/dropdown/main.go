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
		SetOptions([]string{"First", "Second", "Third", "Fourth", "Fifth"}, nil)
		// SetOptionsStyle(
		// 	[]string{"First", "Second", "Third", "Fourth", "Fifth"},
		// 	[]tcell.Style{
		// 		tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true),
		// 		tcell.StyleDefault.Foreground(tcell.ColorGreen).Italic(true),
		// 		tcell.StyleDefault.Foreground(tcell.ColorBlue).Underline(true),
		// 		tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true).Italic(true),
		// 		tcell.StyleDefault.Foreground(tcell.ColorCrimson).Bold(true).Underline(true),
		// 	},
		// 	func(text string, index int) {
		// 		// Show the selected option
		// 		fmt.Printf("Selected: %s (index: %d)\n", text, index)
		// 		app.QueueUpdateDraw(func() {
		// 			dropdown.SetCurrentOption(index)
		// 			fmt.Printf("Current option set to: %d\n", index)
		// 		})
		// 	})

	// Configure the dropdown list to show items properly
	dropdown.SetFieldWidth(20).
		SetDoneFunc(func(key tcell.Key) {
			fmt.Printf("Done with key: %v\n", key)
		}).
		SetFieldBackgroundColor(tcell.ColorGray).
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorBlack),
			tcell.StyleDefault.Background(tcell.ColorDarkGray).Foreground(tcell.ColorWhite),
		)

	if err := app.SetRoot(dropdown, true).EnableMouse(true).Run(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
