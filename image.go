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

func Image() IImage {
	return newImageImpl()
}

func Image2(url string) IImage {
	return newImageImpl().Src(url)
}
