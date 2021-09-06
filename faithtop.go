package faithtop

type IApp interface {
	Run() int
	Quit()
}

const (
	MAX_SIZE = 16777215
)

var appImpl func() IApp

func NewApp() IApp {
	if appImpl == nil {
		panic("please build with impl tag: qtdeploy -tags=impl build desktop")
	}
	return appImpl()
}
