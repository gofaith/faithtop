package faithtop

import (
	"github.com/mattn/go-webkit/webkit"
)

type FWebView struct {
	w *FWindow
	v *webkit.WebView
}

func WebView(w, h int, title string) *FWebView {
	v := &FWebView{}
	v.v = webkit.NewWebView()
	v.w = Win().Title(title).Size(w, h).VScroll(
		v,
	)
	return v
}
func (f *FWebView) getBaseView() *FBaseView {
	return &FBaseView{view: f.v, widget: &f.v.Widget}
}
func (f *FWebView) LoadHtmlString(s string) *FWebView {
	f.v.LoadHtmlString(s, ".")
	return f
}
func (f *FWebView) Show() *FWebView {
	f.w.Show()
	return f
}
func (f *FWebView) Close() *FWebView {
	f.w.Close()
	return f
}
