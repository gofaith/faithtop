package faithtop

type ITableView interface {
	IWidget
}

var newTableViewImpl func(rowCount func() int, columCount func() int, headerData func(section int) string, data func(row, column int) interface{}) ITableView

func TableView(rowCount func() int, columCount func() int, headerData func(section int) string, data func(row, column int) interface{}) ITableView {
	return newTableViewImpl(rowCount, columCount, headerData, data)
}
