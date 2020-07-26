package faithtop

import (
	"github.com/gofaith/walk/declarative"

	"github.com/gofaith/walk"
)

type IView interface {
	baseView() *FBaseView
	widget(builder *declarative.Builder) walk.Widget
	declarative() declarative.Widget
}
type FBaseView struct {
	expand bool
}

//

func RunOnUIThread(fn func()) {
	go fn()
}
