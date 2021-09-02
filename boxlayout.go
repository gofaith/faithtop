package faithtop

type IBoxLayout interface {
	ILayout
	Align(align AlignmentFlag) IBoxLayout
	AppendWidgets(children ...IWidget) IBoxLayout
	AppendItems(items ...ILayoutItem) IBoxLayout
}
