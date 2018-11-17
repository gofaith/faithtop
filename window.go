package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FWindow struct {
	v           *gtk.Window
	wid         string
	showAfter   bool
	child       IView
	ondestroyFn func()
}

func Win() *FWindow {
	w := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	fw := &FWindow{}
	fw.v = w
	fw.wid = NewToken()
	setupWindow(fw)
	return fw
}
func PopupWin() *FWindow {
	w := gtk.NewWindow(gtk.WINDOW_POPUP)
	fw := &FWindow{}
	fw.v = w
	setupWindow(fw)
	return fw
}

func TopWin() *FWindow {
	return Win().Top().Size(200, 100)
}
func TopPopupWin() *FWindow {
	return PopupWin().Top().Size(200, 100)
}
func setupWindow(w *FWindow) {
	w.v.SetDefaultSize(230, 130)
	w.v.SetPosition(gtk.WIN_POS_CENTER)
	w.v.Connect("destroy", func() {
		if w.ondestroyFn != nil {
			w.ondestroyFn()
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
	v.child = i
	if v.showAfter {
		v.Show()
	}
	return v
}

func (v *FWindow) Show() *FWindow {
	v.v.ShowAll()
	if v.child.getBaseView().afterShownFn != nil {
		v.child.getBaseView().afterShownFn()
	}
	return v
}
func (v *FWindow) Close() *FWindow {
	v.v.Destroy()
	return v
}
func (v *FWindow) OnDestroy(f func()) *FWindow {
	v.ondestroyFn = f
	return v
}
func (v *FWindow) OnCloseClicked(f func()) *FWindow {
	v.v.Connect("delete-event", func() bool {
		f()
		return false
	})
	return v
}
func (v *FWindow) DeferShow() *FWindow {
	v.showAfter = true
	return v
}
func (v *FWindow) VBox(is ...IView) *FWindow {
	return v.Add(VBox().Padding(5).Append(is...))
}
func (v *FWindow) HBox(is ...IView) *FWindow {
	return v.Add(HBox().Padding(5).Append(is...))
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
func (v *FWindow) Title(t string) *FWindow {
	v.v.SetTitle(t)
	return v
}
func ShowWin(i IView) *FWindow {
	return TopWin().Size(200, 100).Title("Infomation").DeferShow().Add(i)
}
func ShowInfo(w *FWindow, title, info string) {
	dialog := gtk.NewMessageDialog(
		w.v,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		info)
	dialog.SetTitle(title)
	dialog.Response(func() {
		dialog.Destroy()
	})
	dialog.Run()
}
func ShowErr(w *FWindow, title, err string) {
	dialog := gtk.NewMessageDialog(
		w.v,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_ERROR,
		gtk.BUTTONS_OK,
		err)
	dialog.SetTitle(title)
	dialog.Response(func() {
		dialog.Destroy()
	})
	dialog.Run()
}
