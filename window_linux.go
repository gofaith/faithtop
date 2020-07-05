package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type FWindow struct {
	v         *gtk.Window
	wid       string
	showAfter bool
	child     IView
}

func Win() *FWindow {
	w := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	fw := &FWindow{}
	fw.v = w
	fw.wid = NewToken()
	setupWindow(fw)
	return fw
}

func TopWin() *FWindow {
	return Win().Top().Size(200, 100)
}

func setupWindow(w *FWindow) {
	w.v.SetDefaultSize(230, 130)
	w.v.SetPosition(gtk.WIN_POS_CENTER)
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
	v.v.Add(i.baseView().widget)
	v.child = i
	if v.showAfter {
		v.Show()
	}
	return v
}

func (v *FWindow) Show() *FWindow {
	v.v.ShowAll()
	return v
}
func (v *FWindow) Close() {
	v.v.Destroy()
}

func (f *FWindow) Hide() {
	f.v.Hide()
}

func (f *FWindow) IsVisible() bool {
	return f.v.GetVisible()
}

func (v *FWindow) OnCloseClicked(f func() bool) *FWindow {
	v.v.Connect("delete-event", func() bool {
		return f()
	})
	return v
}
func (v *FWindow) DeferShow() *FWindow {
	v.showAfter = true
	return v
}
func (v *FWindow) VBox(is ...IView) *FWindow {
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
func (v *FWindow) Title(t string) *FWindow {
	v.v.SetTitle(t)
	return v
}
func (f *FWindow) BindInfo(title string, str *livedata.String) *FWindow {
	str.ObserveForever(func(s string) {
		if s != "" {
			RunOnUIThread(func() {
				ShowInfo(f, title, s)
			})
		}
	})
	return f
}
func ShowWin(i IView) *FWindow {
	return TopWin().Size(200, 100).Title("").DeferShow().Add(i)
}
func ShowInfo(w *FWindow, title, info string) {
	dialog := gtk.NewMessageDialog(
		w.v,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_INFO,
		gtk.BUTTONS_OK,
		info)
	dialog.SetTitle(title)
	dialog.Response(func(ctx *glib.CallbackContext) {
		dialog.Destroy()
	})
	dialog.Run()
}
func ShowConfirm(w *FWindow, title, info, ok, cancel string, yes, no func()) {
	dialog := gtk.NewMessageDialog(
		w.v,
		gtk.DIALOG_MODAL,
		gtk.MESSAGE_QUESTION,
		gtk.BUTTONS_NONE,
		info,
	)
	dialog.AddButton(ok, gtk.RESPONSE_YES)
	dialog.AddButton(cancel, gtk.RESPONSE_NO)
	dialog.SetDefaultResponse(gtk.RESPONSE_YES)
	dialog.SetTitle(title)
	switch dialog.Run() {
	case gtk.RESPONSE_YES:
		if yes != nil {
			yes()
		}
	case gtk.RESPONSE_NO:
		if no != nil {
			no()
		}
	}
	dialog.Destroy()
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
