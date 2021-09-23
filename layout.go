package faithtop

type ILayout interface {
	getLayout() ILayout
	ContentMargin(left, top, right, bottom int) ILayout
}
