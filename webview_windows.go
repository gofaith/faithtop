package faithtop

import (
	"github.com/zserge/webview"
	"net/url"
)

type FWebView struct {
	v          webview.WebView
	w, h       int
	title      string
	htmlString string
}

func WebView(w, h int, title string) *FWebView {
	fb := &FWebView{}
	fb.w = w
	fb.h = h
	fb.title = title
	return fb
}
func (v *FWebView) LoadHtmlString(s string) *FWebView {
	v.htmlString = s
	return v
}
func (v *FWebView) Show() *FWebView {
	v.v = webview.New(webview.Settings{
		Width:  v.w,
		Height: v.h,
		Title:  v.title,
		URL:    "data:text/html," + url.PathEscape(v.htmlString),
	})
	go v.v.Run()
	v.v.Exit()
	return v
}
func (v *FWebView) Close() *FWebView {
	v.v.Terminate()
	return v
}
