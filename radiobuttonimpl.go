//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type RadioButtonImpl struct {
	*WidgetImpl
	radioButton *widgets.QRadioButton
}

func init() {
	newRadioButtonImpl = func() IRadioButton {
		v := &RadioButtonImpl{
			radioButton: widgets.NewQRadioButton(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.radioButton.QWidget_PTR())
		return v
	}
}

func (r *RadioButtonImpl) Checked(b bool) IRadioButton {
	r.radioButton.SetChecked(b)
	return r
}

func (r *RadioButtonImpl) Text(text string) IRadioButton {
	r.radioButton.SetText(text)
	return r
}

func (r *RadioButtonImpl) GetText() string {
	return r.radioButton.Text()
}

func (r *RadioButtonImpl) Assign(v *IRadioButton) IRadioButton {
	*v = r
	return r
}

func (r *RadioButtonImpl) IsChecked() bool {
	return r.radioButton.IsChecked()
}

func (r *RadioButtonImpl) Group(g *IButtonGroup) IRadioButton {
	(*g).(*ButtonGroupImpl).buttonGroup.AddButton(r.radioButton, 0)
	return r
}
