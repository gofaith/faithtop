//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type FrameImpl struct {
	*WidgetImpl
	frame *widgets.QFrame
}

func init() {
	newFrameImpl = func() IFrame {
		v := &FrameImpl{
			frame: widgets.NewQFrame(nil, 0),
		}
		v.WidgetImpl = widgetImplFrom(v.frame.QWidget_PTR())
		return v
	}
}

func (f *FrameImpl) Shape(shape FrameShape) IFrame {
	f.frame.SetFrameShape(widgets.QFrame__Shape(shape))
	return f
}

func (f *FrameImpl) Shadow(shadow FrameShadow) IFrame {
	f.frame.SetFrameShadow(widgets.QFrame__Shadow(shadow))
	return f
}

func (f *FrameImpl) LineWidth(i int) IFrame {
	f.frame.SetLineWidth(i)
	return f
}

func (f *FrameImpl) MidLineWidth(i int) IFrame {
	f.frame.SetMidLineWidth(i)
	return f
}
