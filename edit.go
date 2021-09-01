package faithtop

type IEdit interface {
	IWidget
	Assign(v *IEdit) IEdit
	Text(s string) IEdit
	GetText() string
}

var newEditImpl func() IEdit

func Edit() IEdit {
	return newEditImpl()
}
