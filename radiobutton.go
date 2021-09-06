package faithtop

type (
	IRadioButton interface {
		IWidget
		Assign(v *IRadioButton) IRadioButton
		Group(g *IButtonGroup) IRadioButton
		IsChecked() bool
		Checked(b bool) IRadioButton
		Text(text string) IRadioButton
		GetText() string
	}
)

var (
	newRadioButtonImpl func() IRadioButton
)

func RadioButton() IRadioButton {
	return newRadioButtonImpl()
}

func RadioButton2(text string) IRadioButton {
	return newRadioButtonImpl().Text(text)
}
