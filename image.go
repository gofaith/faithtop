package faithtop

type IImage interface {
	IWidget
	ScaleTo(width, height int) IImage
	Assign(v *IImage) IImage
	/** set image source.
	 */
	Src(url string) IImage
}

var (
	newImageImpl func() IImage
)

func Image(w, h int) IImage {
	return newImageImpl().ScaleTo(w, h)
}

func Image2(w, h int, url string) IImage {
	return Image(w, h).Src(url)
}
