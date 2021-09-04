//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type ProgressImpl struct {
	*WidgetImpl
	progress *widgets.QProgressBar
}

func init() {
	newProgressImpl = func() IProgress {
		v := &ProgressImpl{
			progress: widgets.NewQProgressBar(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.progress.QWidget_PTR())
		return v
	}
}
func (p *ProgressImpl) Assign(v *IProgress) IProgress {
	*v = p
	return p
}
func (p *ProgressImpl) Max(i int) IProgress {
	p.progress.SetMaximum(i)
	return p
}

func (p *ProgressImpl) Min(i int) IProgress {
	p.progress.SetMinimum(i)
	return p
}

func (p *ProgressImpl) Value(i int) IProgress {
	p.progress.SetValue(i)
	return p
}

func (p *ProgressImpl) Text(s string) IProgress {
	p.progress.SetFormat(s)
	return p
}

func (p *ProgressImpl) Inverted(b bool) IProgress {
	p.progress.SetInvertedAppearance(b)
	return p
}

func (p *ProgressImpl) GetValue() int {
	return p.progress.Value()
}

func (p *ProgressImpl) Orientation(v Orientation) IProgress {
	p.progress.SetOrientation(core.Qt__Orientation(v))
	return p
}

func (p *ProgressImpl) TextAlign(v AlignmentFlag) IProgress {
	p.progress.SetAlignment(core.Qt__AlignmentFlag(v))
	return p
}

func (p *ProgressImpl) TextVisible(b bool) IProgress {
	p.progress.SetTextVisible(b)
	return p
}
