// github.com/gotk3/gotk3/blob/master/gtk/color_chooser.go

package main

import (
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize gtk
	gtk.Init(nil)

	// Create new window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		panic(err)
	}

	// Set the title of the window
	win.SetTitle("gtk.Color_Chooser Example")

	// Quit gtk when the window is closed
	win.Connect("destroy", func(){
		gtk.MainQuit()
	})

	// Create a new ColorChooserDialogNew object
	colorchooser, err := gtk.ColorChooserDialogNew("Color", win)
	if err != nil {
		panic(err)
	}

	// Add the ColorChooser to the window and show it
	win.Add(colorchooser)
	win.ShowAll()
	gtk.Main()
}
