package faithtop

import (
	"github.com/StevenZack/livedata"
	"log"

	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FImage struct {
	FBaseView
	v   *walk.ImageView
	dec declarative.ImageView
}

func Image() *FImage {
	f := &FImage{}
	f.dec = declarative.ImageView{
		AssignTo: &f.v,
		Mode:     declarative.ImageViewModeZoom,
	}
	return f
}

func (f *FImage) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FImage) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}

func (f *FImage) declarative() declarative.Widget {
	return f.dec
}

func (f *FImage) Assign(v **FImage) *FImage {
	*v = f
	return f
}

func (f *FImage) ScaleTypeFitCenter() *FImage {
	if f.v != nil {
		f.v.SetMode(walk.ImageViewModeZoom)
		return f
	}
	f.dec.Mode = declarative.ImageViewModeZoom
	return f
}

func (f *FImage) ScaleTypeFitXY() *FImage {
	if f.v != nil {
		f.v.SetMode(walk.ImageViewModeStretch)
		return f
	}
	f.dec.Mode = declarative.ImageViewModeStretch
	return f
}

func (f *FImage) ScaleTypeFitImage() *FImage {
	if f.v != nil {
		f.v.SetMode(walk.ImageViewModeCorner)
		return f
	}
	f.dec.Mode = declarative.ImageViewModeCenter
	return f
}

func (f *FImage) Expand() *FImage {
	f.expand = true
	return f
}

func (f *FImage) Src(path string) *FImage {
	image, e := walk.NewImageFromFile(path)
	if e != nil {
		log.Println(e)
		return f
	}
	if f.v != nil {
		f.v.SetImage(image)
		return f
	}
	f.dec.Image = image
	return f
}

func (f *FImage) Size(w, h int) *FImage {
	if f.v != nil {
		f.v.SetSize(walk.Size{Width: w, Height: h})
		return f
	}
	f.dec.MaxSize = declarative.Size{Width: w, Height: h}
	f.dec.MinSize = declarative.Size{Width: w, Height: h}
	return f
}

func (f *FImage) BindSrc(l *livedata.String)*FImage  {
	l.ObserveForever(func(str string){
		if str!=""{
			f.Src(str)
		}
	})
	return f
}