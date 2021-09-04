package faithtop

type ITextArea interface {
	IWidget
	Assign(v *ITextArea) ITextArea
	Text(s string) ITextArea
	Html(s string) ITextArea
	ReadOnly(b bool) ITextArea
	PlainText(s string) ITextArea
	GetText() string
}

var newTextAreaImpl func() ITextArea

func TextArea() ITextArea {
	return newTextAreaImpl()
}
