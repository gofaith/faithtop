package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

func init() {
	gtk.Init(nil)
}

var (
	idMap        = make(map[string]interface{})
	currentImage *FImage
)
