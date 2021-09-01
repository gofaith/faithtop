package faithtop

type IApp interface {
	Run() int
}

var appImpl func() IApp

func NewApp() IApp {
	if appImpl == nil {
		panic("please build with impl tag: qtdeploy -tags=impl build desktop")
	}
	return appImpl()
}
