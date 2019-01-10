package faithtop

import (
	"github.com/mattn/go-webkit/webkit"
)

type FWebView struct {
	FBaseView
	v *webkit.WebView
}

func WebView() *FWebView {
	v := FWebView{}
	v.v = webkit.NewWebView()
	v.view = v.v
	v.widget = &v.v.Widget
	setupWidget(&v)
	return &v
}

// ================================================================
func (v *FWebView) Size(w, h int) *FWebView {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetWebViewByItemId(itemId string) *FWebView {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FWebView); ok {
			return lv
		}
	}
	return nil
}
func (v *FWebView) SetItemId(lv *FListView, itemId string) *FWebView {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetWebViewById(id string) *FWebView {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FWebView); ok {
			return b
		}
	}
	return nil
}
func (v *FWebView) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FWebView) SetId(id string) *FWebView {
	idMap[id] = v
	return v
}
func (v *FWebView) Expand() *FWebView {
	v.expand = true
	return v
}
func (v *FWebView) NotFill() *FWebView {
	v.notFill = true
	return v
}
func (v *FWebView) Disable() *FWebView {
	v.v.SetSensitive(false)
	return v
}
func (v *FWebView) Enable() *FWebView {
	v.v.SetSensitive(true)
	return v
}
func (v *FWebView) Visible() *FWebView {
	v.v.SetVisible(true)
	return v
}
func (v *FWebView) Invisible() *FWebView {
	v.FBaseView.Invisible()
	return v
}

func (v *FWebView) Tooltips(s string) *FWebView {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FWebView) Focus() *FWebView {
	v.FBaseView.Focus()
	return v
}
func (v *FWebView) Padding(i uint) *FWebView {
	v.padding = i
	return v
}
func (v *FWebView) OnDragDrop(f func([]string)) *FWebView {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FWebView) LoadHtmlString(s, baseUri string) *FWebView {
	v.v.LoadHtmlString(s, baseUri)
	return v
}
