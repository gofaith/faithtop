package faithtop

type IProgress interface {
	IWidget
	Assign(v *IProgress) IProgress
	Max(i int) IProgress
	Min(i int) IProgress
	Value(i int) IProgress
	/** text on progress bar
	e.g., "%p%" => 24%, "loading" => loading
	*/
	Text(s string) IProgress
	Inverted(b bool) IProgress
	GetValue() int
	Orientation(v Orientation) IProgress
	TextAlign(v AlignmentFlag) IProgress
	TextVisible(b bool) IProgress
}

var newProgressImpl func() IProgress

func Progress() IProgress {
	return newProgressImpl()
}
