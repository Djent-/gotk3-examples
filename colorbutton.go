// github.com/gotk3/gotk3/blob/master/gtk/gtk.go

package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
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
	win.SetTitle("gtk.ColorButton Example")

	// Quit gtk when the window is closed
	win.Connect("destroy", func(){
		log.Println("Main window being destroyed")
		gtk.MainQuit()
	})

	// Create a new ColorButton object
	colorbutton, err := gtk.ColorButtonNew()
	if err != nil {
		panic(err)
	}

	colorbutton.Connect("color-set", func(){
		// When the user selects a color
		color := colorbutton.GetRGBA()
		log.Println("Added color:", color)
		// Quit the program
		//gtk.MainQuit()
	})

	// Add the ColorButton to the window and show it
	win.Add(colorbutton)
	win.ShowAll()

	// Run the main gtk loop
	gtk.Main()
}
