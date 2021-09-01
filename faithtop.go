package faithtop

type IApp interface {
	Run() int
}

var appImpl func() IApp

func NewApp() IApp {
	if appImpl == nil {
		panic("please build go code with impl tag: go build -tags impl")
	}
	return appImpl()
}
