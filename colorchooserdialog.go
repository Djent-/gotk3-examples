// github.com/gotk3/gotk3/blob/master/gtk/color_chooser.go

package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
)

func main() {
	// Initialize gtk
	gtk.Init(nil)

	// Create a new ColorChooserDialog object
	colorchooser, err := gtk.ColorChooserDialogNew("Color", nil)
	if err != nil {
		panic(err)
	}

	// Show the ColorChooserDialog
	colorchooser.ShowAll()

	// Get the results from the ColorChooserDialog
	if colorchooser.Run() == int(gtk.RESPONSE_OK) {
		// Grab the results
		color := colorchooser.GetRGBA()
		log.Println("Added color:", color)
		// Destroy the dialog and exit the program
		colorchooser.Destroy()
		gtk.MainQuit()
		// gtk actually throws an error using MainQuit()
		// so I'll exit the program fully with golang
		// may have unintended side effects though
		os.Exit(0)
	}

	// Run the main gtk loop
	gtk.Main()
}
