package faithtop

type ICheckBox interface {
	IWidget
	Assign(v *ICheckBox) ICheckBox
	IsChecked() bool
	Checked(b bool) ICheckBox
	Text(s string) ICheckBox
	OnChanged(fn func(b bool)) ICheckBox
}

var newCheckBoxImpl func() ICheckBox

func CheckBox() ICheckBox {
	return newCheckBoxImpl()
}
