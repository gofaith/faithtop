package faithtop

type IPlainTextArea interface {
	IWidget
	Assign(v *IPlainTextArea) IPlainTextArea
	PlainText(s string) IPlainTextArea
	ReadOnly(b bool) IPlainTextArea
	GetText() string
}

var newPlainTextAreaImpl func() IPlainTextArea

func PlainTextArea() IPlainTextArea {
	return newPlainTextAreaImpl()
}
