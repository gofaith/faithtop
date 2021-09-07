package faithtop

type IGroupBox interface {
	IWidget
	Title(title string) IGroupBox
	Checkable(b bool) IGroupBox
	Checked(b bool) IGroupBox
	IsChecked() bool
	OnToggled(fn func(self IGroupBox, b bool)) IGroupBox
	TitleAlign(align AlignmentFlag) IGroupBox
	Flat(b bool) IGroupBox
}

var newGroupBoxImpl func() IGroupBox

func GroupBox() IGroupBox {
	return newGroupBoxImpl()
}
