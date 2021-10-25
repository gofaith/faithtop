//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type PlainTextAreaImpl struct {
	*WidgetImpl
	textarea *widgets.QPlainTextEdit
}

func init() {
	newPlainTextAreaImpl = func() IPlainTextArea {
		v := &PlainTextAreaImpl{
			textarea: widgets.NewQPlainTextEdit(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.textarea.QWidget_PTR())
		return v
	}
}
func (t *PlainTextAreaImpl) Assign(v *IPlainTextArea) IPlainTextArea {
	*v = t
	return t
}
func (t *PlainTextAreaImpl) PlainText(s string) IPlainTextArea {
	t.textarea.SetPlainText(s)
	return t
}

func (t *PlainTextAreaImpl) GetText() string {
	return t.textarea.ToPlainText()
}

func (t *PlainTextAreaImpl) ReadOnly(b bool) IPlainTextArea {
	t.textarea.SetReadOnly(b)
	return t
}
