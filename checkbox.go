package faithtop

type ICheckBox interface {
	IWidget
	Assign(v *ICheckBox) ICheckBox
	IsChecked() bool
	GetCheckState() CheckState
	Checked(b bool) ICheckBox
	CheckState(state CheckState) ICheckBox
	Text(s string) ICheckBox
	OnChanged(fn func(b bool)) ICheckBox
	OnStateChanged(fn func(state CheckState)) ICheckBox
}

var newCheckBoxImpl func() ICheckBox

func CheckBox() ICheckBox {
	return newCheckBoxImpl()
}
