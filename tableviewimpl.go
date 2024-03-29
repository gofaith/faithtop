//go:build impl

package faithtop

import (
	"sync"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type (
	TableViewImpl struct {
		*WidgetImpl
		tableView *widgets.QTableView
		model     *tableViewModel

		rowCount, columnCount func() int
		headerData            func(section int) string
		data                  func(row, column int) interface{}
	}
	tableViewModel struct {
		core.QAbstractTableModel
		_          func() `constructor:"init"`
		headerData []*core.QVariant
		modelData  [][]*core.QVariant
		widgetData []tableViewModelWidgetData
		locker     sync.Mutex
	}
	tableViewModelWidgetData struct {
		row, column int
		widget      widgets.QWidget_ITF
	}
)

func init() {
	newTableViewImpl = func(rowCount, columCount func() int, headerData func(section int) string, data func(row, column int) interface{}) ITableView {
		v := &TableViewImpl{
			tableView:   widgets.NewQTableView(nil),
			rowCount:    rowCount,
			columnCount: columCount,
			headerData:  headerData,
			data:        data,
		}
		v.WidgetImpl = widgetImplFrom(v.tableView.QWidget_PTR())

		v.model = NewTableViewModel(nil)
		widgetData := v.createDataAll()

		v.tableView.SetModel(v.model)

		v.model.widgetData = widgetData
		v.loadWidgets(widgetData)
		return v
	}
}

func (t *tableViewModel) init() {
	t.ConnectHeaderData(func(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) || orientation == core.Qt__Vertical {
			return t.HeaderDataDefault(section, orientation, role)
		}
		if len(t.headerData) > section {
			return t.headerData[section]
		}

		return core.NewQVariant()
	})
	t.ConnectRowCount(func(*core.QModelIndex) int {
		return len(t.modelData)
	})
	t.ConnectColumnCount(func(*core.QModelIndex) int {
		return len(t.headerData)
	})
	t.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		//return widget
		row, column := index.Row(), index.Column()
		if len(t.modelData) > row {
			vs := t.modelData[row]
			if len(vs) > column {
				return vs[column]
			}
		}

		return core.NewQVariant()
	})
}

func (t *TableViewImpl) createDataAll() []tableViewModelWidgetData {
	headerData := []*core.QVariant{}
	columnCount := t.columnCount()
	for i := 0; i < columnCount; i++ {
		headerData = append(headerData, core.NewQVariant1(t.headerData(i)))
	}

	modelData := [][]*core.QVariant{}
	widgetData := []tableViewModelWidgetData{}
	rowCount := t.rowCount()
	for rowNum := 0; rowNum < rowCount; rowNum++ {
		row := []*core.QVariant{}
		for columnNum := 0; columnNum < columnCount; columnNum++ {

			//load widget
			v := t.data(rowNum, columnNum)
			var variant *core.QVariant
			if str, ok := v.(string); ok {
				variant = core.NewQVariant1(str)
			} else if iwidget, ok := v.(IWidget); ok {
				widgetItem := tableViewModelWidgetData{
					row:    rowNum,
					column: columnNum,
					widget: iwidget.getWidget().(*WidgetImpl).widget,
				}
				widgetData = append(widgetData, widgetItem)
			} else {
				panic("the return value type can only be [string] or [IWidget]")
			}
			if variant == nil {
				variant = core.NewQVariant()
			}
			row = append(row, variant)
		}
		modelData = append(modelData, row)
	}

	t.model.locker.Lock()
	defer t.model.locker.Unlock()
	t.model.headerData = headerData
	t.model.modelData = modelData
	return widgetData
}
func (t *TableViewImpl) createData(row, column int) (*core.QVariant, []tableViewModelWidgetData) {
	v := t.data(row, column)
	var variant *core.QVariant
	widgetData := []tableViewModelWidgetData{}

	if str, ok := v.(string); ok {
		variant = core.NewQVariant1(str)
	} else if iwidget, ok := v.(IWidget); ok {
		widgetItem := tableViewModelWidgetData{
			row:    row,
			column: column,
			widget: iwidget.getWidget().(*WidgetImpl).widget,
		}
		widgetData = append(widgetData, widgetItem)
	} else {
		panic("the return value type can only be [string] or [IWidget]")
	}
	if variant == nil {
		variant = core.NewQVariant()
	}
	return variant, widgetData
}
func (t *TableViewImpl) updateData(row, column int) []tableViewModelWidgetData {
	variant, widgetData := t.createData(row, column)
	t.model.modelData[row][column] = variant
	return widgetData
}

func (t *TableViewImpl) loadWidgets(widgetData []tableViewModelWidgetData) {
	for _, item := range widgetData {
		t.tableView.SetIndexWidget(t.model.Index(item.row, item.column, core.NewQModelIndex()), item.widget)
	}
}

func (t *TableViewImpl) Assign(v *ITableView) ITableView {
	*v = t
	return t
}

func (t *TableViewImpl) DataChanged(left, top, right, bottom int) ITableView {
	r, r2 := t.rowCount(), t.RowCount()
	c, c2 := t.columnCount(), t.ColumnCount()
	widgetData := []tableViewModelWidgetData{}
	for columnNum := left; columnNum <= right; columnNum++ {
		for rowNum := top; rowNum <= bottom; rowNum++ {
			if columnNum < c && columnNum < c2 && rowNum < r && rowNum < r2 {
				widget := t.updateData(rowNum, columnNum)
				widgetData = append(widgetData, widget...)
			}
		}
	}
	t.loadWidgets(widgetData)
	t.model.DataChanged(t.model.Index(top, left, core.NewQModelIndex()), t.model.Index(bottom, right, core.NewQModelIndex()), []int{int(core.Qt__DisplayRole)})
	return t
}

func (t *TableViewImpl) Reload() ITableView {
	rowCount := t.RowCount()
	columnCount := t.ColumnCount()
	t.loadWidgets(t.createDataAll())
	t.DataChanged(0, 0, columnCount-1, rowCount-1)
	return t
}

func (t *TableViewImpl) RowCount() int {
	return len(t.model.modelData)
}
func (t *TableViewImpl) ColumnCount() int {
	return len(t.model.headerData)
}

func (t *TableViewImpl) HeaderCount() int {
	return len(t.model.headerData)
}

func (t *TableViewImpl) RemoveRow(row int) {
	if len(t.model.modelData) == 0 {
		return
	}
	t.model.BeginRemoveRows(core.NewQModelIndex(), row, row)
	t.model.modelData = append(t.model.modelData[:row], t.model.modelData[row+1:]...)
	//widgetData
	widgetDataChanged := []tableViewModelWidgetData{}
	newWidgetData := []tableViewModelWidgetData{}
	for _, widget := range t.model.widgetData {
		if widget.row < row {
			newWidgetData = append(newWidgetData, widget)
		} else if widget.row == row {
		} else {
			widget.row--
			newWidgetData = append(newWidgetData, widget)
			widgetDataChanged = append(widgetDataChanged, widget)
		}
	}
	t.loadWidgets(widgetDataChanged)
	t.model.widgetData = newWidgetData

	t.model.EndRemoveRows()
}

func (t *TableViewImpl) RemoveRows(from, count int) {
	if len(t.model.modelData) == 0 {
		return
	}
	t.model.BeginRemoveRows(core.NewQModelIndex(), from, from+count-1)
	t.model.modelData = append(t.model.modelData[:from], t.model.modelData[from+count:]...)
	//widgetData
	widgetDataChanged := []tableViewModelWidgetData{}
	newWidgetData := []tableViewModelWidgetData{}
	for _, widget := range t.model.widgetData {
		if widget.row < from {
			newWidgetData = append(newWidgetData, widget)
		} else if widget.row < from+count {
		} else {
			widget.row -= count
			newWidgetData = append(newWidgetData, widget)
			widgetDataChanged = append(widgetDataChanged, widget)
		}
	}
	t.loadWidgets(widgetDataChanged)
	t.model.widgetData = newWidgetData

	t.model.EndRemoveRows()
}

func (t *TableViewImpl) AddRow(row int) {
	t.model.BeginInsertRows(core.NewQModelIndex(), row, row)
	columnCount := t.ColumnCount()
	widgetData := []tableViewModelWidgetData{}
	qvs := []*core.QVariant{}
	for column := 0; column < columnCount; column++ {
		qv, widget := t.createData(row, column)
		qvs = append(qvs, qv)
		widgetData = append(widgetData, widget...)
	}
	t.model.modelData = append(append(t.model.modelData[:row], qvs), t.model.modelData[row:]...)

	//widgetData
	newWidgetData := []tableViewModelWidgetData{}
	newWidgetData = append(newWidgetData, widgetData...)
	for i, widget := range t.model.widgetData {
		if widget.row >= row {
			ws := t.updateData(widget.row+1, widget.column)
			if len(ws) > 0 {
				t.model.widgetData[i] = ws[0]
				widgetData = append(widgetData, ws[0])
				newWidgetData = append(newWidgetData, ws[0])
			}
		} else {
			newWidgetData = append(newWidgetData, widget)
		}
	}
	t.loadWidgets(widgetData)
	t.model.widgetData = newWidgetData

	t.model.EndInsertRows()
}

func (t *TableViewImpl) AddRows(row, count int) {
	t.model.BeginInsertRows(core.NewQModelIndex(), row, row+count-1)
	columnCount := t.ColumnCount()

	widgetData := []tableViewModelWidgetData{}
	qvss := [][]*core.QVariant{}
	for i := 0; i < count; i++ {
		qvs := []*core.QVariant{}
		for column := 0; column < columnCount; column++ {
			qv, widget := t.createData(row+i, column)
			qvs = append(qvs, qv)
			widgetData = append(widgetData, widget...)
		}
		qvss = append(qvss, qvs)
	}

	t.model.modelData = append(append(t.model.modelData[:row], qvss...), t.model.modelData[row:]...)

	//widgetData
	newWidgetData := []tableViewModelWidgetData{}
	newWidgetData = append(newWidgetData, widgetData...)
	for i, widget := range t.model.widgetData {
		if widget.row >= row {
			ws := t.updateData(widget.row+count, widget.column)
			if len(ws) > 0 {
				t.model.widgetData[i] = ws[0]
				widgetData = append(widgetData, ws[0])
				newWidgetData = append(newWidgetData, ws[0])
			}
		} else {
			newWidgetData = append(newWidgetData, widget)
		}
	}
	t.loadWidgets(widgetData)
	t.model.widgetData = newWidgetData

	t.model.EndInsertRows()
}

func (t *TableViewImpl) ShowGrid(b bool) ITableView {
	t.tableView.SetShowGrid(b)
	return t
}

func (t *TableViewImpl) NoHorizontalHeader() ITableView {
	h := widgets.NewQHeaderView(core.Qt__Horizontal, nil)
	h.SetVisible(false)
	t.tableView.SetHorizontalHeader(h)
	return t
}

func (t *TableViewImpl) NoVerticalHeader() ITableView {
	v := widgets.NewQHeaderView(core.Qt__Vertical, nil)
	v.SetVisible(false)
	t.tableView.SetVerticalHeader(v)
	return t
}

func (t *TableViewImpl) NoSelection() ITableView {
	t.tableView.SetSelectionMode(widgets.QAbstractItemView__NoSelection)
	return t
}
