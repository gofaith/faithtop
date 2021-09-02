package faithtop

type IButton interface {
	IWidget
	Text(s string) IButton
	OnClick(fn func()) IButton
	Enabled(b bool) IButton
	Flat(b bool) IButton
}

var newButtonImpl func() IButton

func Button() IButton {
	return newButtonImpl()
}
