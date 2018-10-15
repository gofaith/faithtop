package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

var (
	windowCounter int
)

type FWindow struct {
	v         *gtk.Window
	showAfter bool
}

func Win() *FWindow {
	w := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	setupWindow(w)
	fw := &FWindow{}
	fw.v = w
	return fw
}
func PopupWin() *FWindow {
	w := gtk.NewWindow(gtk.WINDOW_POPUP)
	setupWindow(w)
	fw := &FWindow{}
	fw.v = w
	return fw
}

func TopWin() *FWindow {
	return Win().Top()
}
func TopPopupWin() *FWindow {
	return PopupWin().Top()
}
func setupWindow(w *gtk.Window) {
	windowCounter++
	w.SetDefaultSize(600, 400)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.Connect("destroy", func() {
		windowCounter--
		if windowCounter < 1 {
			gtk.MainQuit()
		}
	})
}

func GetWinById(id string) *FWindow {
	v, ok := idMap[id]
	if ok {
		if w, is := v.(*FWindow); is {
			return w
		}
	}
	return nil
}

func (v *FWindow) SetId(id string) *FWindow {
	idMap[id] = v
	return v
}
func (v *FWindow) Size(width, height int) *FWindow {
	v.v.SetDefaultSize(width, height)
	return v
}
func (v *FWindow) Top() *FWindow {
	v.v.SetModal(true)
	v.v.SetKeepAbove(true)
	return v
}

func (v *FWindow) Add(i IView) *FWindow {
	v.v.Add(i.getBaseView().view)
	i.getBaseView().alreadyAdded = true
	if i.getBaseView().afterAppend != nil {
		i.getBaseView().afterAppend()
		i.getBaseView().afterAppend = nil
	}
	if v.showAfter {
		v.Show()
	}
	return v
}

func (v *FWindow) Show() *FWindow {
	v.v.ShowAll()
	if windowCounter == 1 {
		gtk.Main()
	}
	return v
}
func (v *FWindow) Close() *FWindow {
	v.v.Destroy()
	return v
}

func (v *FWindow) DeferShow() *FWindow {
	v.showAfter = true
	return v
}
func (v *FWindow) Vbox(is ...IView) *FWindow {
	return v.Add(VBox().Append(is...))
}
func (v *FWindow) HBox(is ...IView) *FWindow {
	return v.Add(HBox().Append(is...))
}
func (v *FWindow) VScroll(is ...IView) *FWindow {
	return v.Add(VScroll().Append(is...))
}
func (v *FWindow) HScroll(is ...IView) *FWindow {
	return v.Add(HScroll().Append(is...))
}
func (v *FWindow) Scroll(child IView) *FWindow {
	return v.Add(Scroll(child))
}
