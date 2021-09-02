package faithtop

type IText interface {
	IWidget
	Text(s string) IText
}
