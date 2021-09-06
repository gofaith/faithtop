//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type ButtonGroupImpl struct {
	buttonGroup *widgets.QButtonGroup
}

func init() {
	newButtonGroupImpl = func() IButtonGroup {
		return &ButtonGroupImpl{
			buttonGroup: widgets.NewQButtonGroup(nil),
		}
	}
}

func (b *ButtonGroupImpl) OnClicked(fn func(text string)) IButtonGroup {
	b.buttonGroup.ConnectButtonClicked(func(button *widgets.QAbstractButton) {
		if fn != nil {
			fn(button.Text())
		}
	})
	return b
}

func (b *ButtonGroupImpl) BindValue(v *string) IButtonGroup {
	b.OnClicked(func(text string) {
		if *v != text {
			*v = text
		}
	})
	return b
}
