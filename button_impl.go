//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type ButtonImpl struct {
	*WidgetImpl
	button *widgets.QPushButton
}

func init() {
	newButtonImpl = func() IButton {
		v := &ButtonImpl{
			button: widgets.NewQPushButton(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.button.QWidget_PTR())
		return v
	}
}

func (b *ButtonImpl) Text(s string) IButton {
	b.button.SetText(s)
	return b
}

func (b *ButtonImpl) OnClick(fn func()) IButton {
	b.button.ConnectClicked(func(checked bool) {
		if fn != nil {
			fn()
		}
	})
	return b
}

func (b *ButtonImpl) Enabled(enabled bool) IButton {
	b.button.SetEnabled(enabled)
	return b
}

func (b *ButtonImpl) Flat(flat bool) IButton {
	b.button.SetFlat(flat)
	return b
}
