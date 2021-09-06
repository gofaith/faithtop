package faithtop

type IComboBox interface {
	IWidget
	Assign(v *IComboBox) IComboBox
	Data(texts []string) IComboBox
	Current(i int)IComboBox
	Clear() IComboBox
	Editable(b bool) IComboBox
	OnTextChanged(fn func(text string)) IComboBox
	OnIndexChanged(fn func(idx int)) IComboBox
}

var newComboBoxImpl func() IComboBox

func ComboBox() IComboBox {
	return newComboBoxImpl()
}

