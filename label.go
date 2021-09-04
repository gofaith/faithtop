package faithtop

type ILabel interface {
	IWidget
	Assign(v *ILabel) ILabel
	/** set text of the label. e.g.,
	Text("hello world")
	we also support rich HTML text:
	Text(`<html><head/><body><p>asdkhkhqw</p><p><br/></p><p><span style=" font-weight:600;">213</span></p><p><br/></p><p>12</p></body></html>``)
	*/
	Text(s string) ILabel
	Image(url string, w, h int) ILabel
	Interaction(interaction TextInteractionFlag) ILabel
	GetText() string
}

var newLabelImpl func() ILabel

func Label() ILabel {
	return newLabelImpl()
}

func Label2(text string) ILabel {
	return newLabelImpl().Text(text)
}

func LabelImage(url string, w, h int) ILabel {
	return Label().Image(url, w, h)
}
