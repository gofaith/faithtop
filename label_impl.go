//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type LabelImpl struct {
	*WidgetImpl
	label *widgets.QLabel
}

func init() {
	newLabelImpl = func() ILabel {
		v := &LabelImpl{
			label: widgets.NewQLabel(nil, 0),
		}
		v.WidgetImpl = widgetImplFrom(v.label.QWidget_PTR())
		return v
	}
}

func (l *LabelImpl) Text(s string) ILabel {
	l.label.SetText(s)
	return l
}

func (l *LabelImpl) Interaction(flags TextInteractionFlag) ILabel {
	l.label.SetTextInteractionFlags(core.Qt__TextInteractionFlag(flags))
	return l
}
