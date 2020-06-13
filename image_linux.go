package faithtop

import (
	"time"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

type FImage struct {
	FBaseView
	v            *gtk.Image
	onLoad       func(string)
	scaleType    int
	alreadyShown bool
	onLoadFailed func(string)
	// 0 : fitCenter , 1 : fitXY , 2 : origin
}

func Image() *FImage {
	v := gtk.NewImage()
	fb := &FImage{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	currentImage = fb
	fb.scaleType = 2
	fb.Src("/")
	return fb
}

// ================================================================

func (v *FImage) Size(w, h int) *FImage {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FImage) Assign(v **FImage) *FImage {
	*v = f
	return f
}
func (vh *ViewHolder) GetImageByItemId(itemId string) *FImage {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FImage); ok {
			return lv
		}
	}
	return nil
}
func (v *FImage) SetItemId(lv *FListView, itemId string) *FImage {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetImageById(id string) *FImage {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FImage); ok {
			return b
		}
	}
	return nil
}
func (v *FImage) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FImage) SetId(id string) *FImage {
	idMap[id] = v
	return v
}
func (v *FImage) Expand() *FImage {
	v.expand = true
	return v
}
func (v *FImage) NotFill() *FImage {
	v.notFill = true
	return v
}
func (v *FImage) Disable() *FImage {
	v.v.SetSensitive(false)
	return v
}
func (v *FImage) Enable() *FImage {
	v.v.SetSensitive(true)
	return v
}
func (v *FImage) Visible() *FImage {
	v.v.SetVisible(true)
	return v
}
func (v *FImage) Invisible() *FImage {
	v.FBaseView.Invisible()
	return v
}

func (v *FImage) Tooltips(s string) *FImage {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FImage) Focus() *FImage {
	v.FBaseView.Focus()
	return v
}
func (v *FImage) Padding(i uint) *FImage {
	v.padding = i
	return v
}

func (v *FImage) OnDragDrop(f func([]string)) *FImage {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FImage) Src(url string) *FImage {
	if StartsWidth(url, "file://") {
		setImageFileSrc(v, url[len("file://"):])
	} else if StartsWidth(url, "http") {
		go CacheNetFile(url, GetCacheDir(), func(fpath string) {
			RunOnUIThread(func() {
				setImageFileSrc(v, fpath)
			})
		})
	} else {
		setImageFileSrc(v, url)
	}
	return v
}
func (v *FImage) SrcGif(url string) *FImage {
	v.scaleType = 2
	return v.Src(url)
}

func setImageFileSrc(v *FImage, url string) {
	var setup = func() {
		v.alreadyShown = true
		porigin, e := gdkpixbuf.NewPixbufFromFile(url)
		if e != nil {
			return
		}
		var w, h int = v.v.GetAllocation().Width, v.v.GetAllocation().Height
		if v.initWidth == -56 { //initialize it
			v.initWidth = v.v.GetAllocation().Width
			v.initHeight = v.v.GetAllocation().Height
		}
		var widthIsWrap, heightIsWrap = (v.initWidth < 2), (v.initHeight < 2)
		if v.scaleType == 0 { //fitCenter
			if widthIsWrap && !heightIsWrap {
				w = int(float64(v.v.GetAllocation().Height) / float64(porigin.GetHeight()) * float64(porigin.GetWidth()))
			} else if heightIsWrap && !widthIsWrap {
				h = int(float64(v.v.GetAllocation().Width) / float64(porigin.GetWidth()) * float64(porigin.GetHeight()))
			} else if widthIsWrap && heightIsWrap {
				w = porigin.GetWidth()
				h = porigin.GetHeight()
			}
			p, e := gdkpixbuf.NewPixbufFromFileAtScale(url, w, h, true)
			if e != nil {
				if v.onLoadFailed != nil {
					v.onLoadFailed(e.Error())
				}
				return
			}
			v.v.SetFromPixbuf(p)
		} else if v.scaleType == 1 { //fitXY
			if widthIsWrap && !heightIsWrap {
				w = porigin.GetWidth()
			}
			if heightIsWrap && !widthIsWrap {
				h = porigin.GetHeight()
			}
			p, e := gdkpixbuf.NewPixbufFromFileAtScale(url, w, h, false)
			if e != nil {
				if v.onLoadFailed != nil {
					v.onLoadFailed(e.Error())
				}
				return
			}
			v.v.SetFromPixbuf(p)
		} else {
			v.v.SetFromFile(url)
		}
		if v.onLoad != nil {
			v.onLoad(url)
		}
	}
	go func() {
		time.Sleep(time.Millisecond * 50)
		RunOnUIThread(func() {
			setup()
		})
	}()
	return
}
func (v *FImage) OnLoad(f func(string)) *FImage {
	v.onLoad = f
	return v
}
func (v *FImage) GetPixbuf() *gdkpixbuf.Pixbuf {
	return v.v.GetPixbuf()
}
func (v *FImage) SetPixbuf(p *gdkpixbuf.Pixbuf) *FImage {
	v.v.SetFromPixbuf(p)
	return v
}
func (v *FImage) GetImageWidth() int {
	return v.v.GetPixbuf().GetWidth()
}
func (v *FImage) GetImageHeight() int {
	return v.v.GetPixbuf().GetHeight()
}
func (v *FImage) FlipHorizontally() {
	p := v.v.GetPixbuf().Flip(true)
	v.v.SetFromPixbuf(p)
}
func (v *FImage) FlipVertically() {
	p := v.v.GetPixbuf().Flip(false)
	v.v.SetFromPixbuf(p)
}
func (v *FImage) ScaleTypeFitCenter() *FImage {
	v.scaleType = 0
	return v
}

func (v *FImage) ScaleTypeFitXY() *FImage {
	v.scaleType = 1
	return v
}

func (v *FImage) ScaleTypeFitImage() *FImage {
	v.scaleType = 2
	return v
}
func (v *FImage) OnLoadFailed(f func(string)) *FImage {
	v.onLoadFailed = f
	return v
}
