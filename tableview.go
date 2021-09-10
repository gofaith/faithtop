package faithtop

type ITableView interface {
	IWidget
	Assign(v *ITableView) ITableView
	DataChanged(left, top, right, bottom int) ITableView
	Reload() ITableView
	RowCount() int
	HeaderCount() int
	RemoveRow(i int)
	RemoveRows(from, count int)
	AddRow(i int)
	AddRows(from, count int)
	ShowGrid(b bool) ITableView
	NoHorizontalHeader() ITableView
	NoVerticalHeader() ITableView
	NoSelection()ITableView
}

var newTableViewImpl func(rowCount func() int, columCount func() int, headerData func(section int) string, data func(row, column int) interface{}) ITableView

func TableView(rowCount func() int, columCount func() int, headerData func(section int) string, data func(row, column int) interface{}) ITableView {
	return newTableViewImpl(rowCount, columCount, headerData, data)
}
