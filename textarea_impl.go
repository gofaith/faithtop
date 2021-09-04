//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type TextAreaImpl struct {
	*WidgetImpl
	textarea *widgets.QTextEdit
}

func init() {
	newTextAreaImpl = func() ITextArea {
		v := &TextAreaImpl{
			textarea: widgets.NewQTextEdit(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.textarea.QWidget_PTR())
		return v
	}
}
func (t *TextAreaImpl) Assign(v *ITextArea) ITextArea {
	*v = t
	return t
}
func (t *TextAreaImpl) Text(s string) ITextArea {
	t.textarea.SetText(s)
	return t
}

func (t *TextAreaImpl) Html(s string) ITextArea {
	t.textarea.SetHtml(s)
	return t
}

func (t *TextAreaImpl) GetText() string {
	return t.textarea.ToPlainText()
}

func (t *TextAreaImpl) PlainText(s string) ITextArea {
	t.textarea.SetPlainText(s)
	return t
}
func (t *TextAreaImpl) ReadOnly(b bool) ITextArea {
	t.textarea.SetReadOnly(b)
	return t
}
