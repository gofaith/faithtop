package faithtop

type IWidget interface {
	getWidget() IWidget
	AssignWidget(v *IWidget) IWidget
	SizePolicy(width, height SizePolicy) IWidget
	MinWidth(width int) IWidget
	MinHeight(height int) IWidget
	MaxWidth(width int) IWidget
	MaxHeight(height int) IWidget
	Layout(layout ILayout) IWidget
	Align(align AlignmentFlag) IWidget
	/** Style set a CSS stylesheet string to a widget.
	e.g.
	Button().Style(`QPushButton{background-color:blue;color:white;} QPushButton:hover{background-color:red;} QPushButton:pressed{background-color:white;color:grey}`)
	Button().Style(`border-image: url(:/qml/a.png);`)
	*/
	Style(styleSheet string) IWidget
	OnDragEnter(fn func(widget IWidget)) IWidget
	OnDragLeave(fn func(widget IWidget)) IWidget
	OnDrop(fn func(widget IWidget, urls []string)) IWidget
	Enabled(b bool) IWidget
	Cursor(shape CursorShape) IWidget
	OnMousePress(fn func(widget IWidget)) IWidget
	OnMouseRelease(fn func(widget IWidget)) IWidget
	OnKeyPress(fn func(widget IWidget, key Key)) IWidget
	ToolTip(tip string) IWidget
}

var newWidgetImpl func() IWidget

func Widget() IWidget {
	return newWidgetImpl()
}
