//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"strconv"
	"strings"
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

func (l *LabelImpl) Assign(v *ILabel) ILabel {
	*v = l
	return l
}

func (l *LabelImpl) Interaction(flags TextInteractionFlag) ILabel {
	l.label.SetTextInteractionFlags(core.Qt__TextInteractionFlag(flags))
	return l
}

func (l *LabelImpl) Image(url string, w, h int) ILabel {
	builder := new(strings.Builder)
	builder.WriteString("<img src='" + url + "' ")
	if w > 0 {
		builder.WriteString("width='" + strconv.Itoa(w) + "' ")
	}
	if h > 0 {
		builder.WriteString("height='" + strconv.Itoa(h) + "' ")
	}
	builder.WriteString(">")
	return l.Text(builder.String())
}

func (l *LabelImpl) GetText() string {
	return l.label.Text()
}
