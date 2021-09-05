package faithtop

type IButton interface {
	IWidget
	Text(s string) IButton
	OnClick(fn func()) IButton
	Flat(b bool) IButton
}

var newButtonImpl func() IButton

func Button() IButton {
	return newButtonImpl()
}

func Button2(text string) IButton {
	return newButtonImpl().Text(text)
}
