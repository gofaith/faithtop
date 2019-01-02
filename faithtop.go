package faithtop

import (
	"fmt"

	"github.com/mattn/go-gtk/gtk"
)

func init() {
	gtk.Init(nil)
	fmt.Println("gtk initialized")
}

var (
	idMap        = make(map[string]interface{})
	groupIdMap   = make(map[string]*FRadio)
	currentImage *FImage
)

func MainQuit() {
	gtk.MainQuit()
}

func Main() {
	gtk.Main()
}
