package faithtop

type IEdit interface {
	IWidget
	Assign(v *IEdit) IEdit
	Text(s string) IEdit
	GetText() string
	Placeholder(s string) IEdit
	PasswordMode(b bool) IEdit
	OnReturn(fn func(self IEdit)) IEdit
}

var newEditImpl func() IEdit

func Edit() IEdit {
	return newEditImpl()
}
