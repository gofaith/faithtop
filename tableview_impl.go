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
		widgetData []tableViewModelData
		locker     sync.Mutex
	}
	tableViewModelData struct {
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
		v.createData()

		v.tableView.SetModel(v.model)

		v.loadWidgets()
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

func (t *TableViewImpl) createData() {
	headerData := []*core.QVariant{}
	columnCount := t.columnCount()
	for i := 0; i < columnCount; i++ {
		headerData = append(headerData, core.NewQVariant1(t.headerData(i)))
	}

	modelData := [][]*core.QVariant{}
	widgetData := []tableViewModelData{}
	rowCount := t.rowCount()
	for i := 0; i < rowCount; i++ {
		row := []*core.QVariant{}
		for j := 0; j < columnCount; j++ {
			v := t.data(i, j)

			var variant *core.QVariant
			if str, ok := v.(string); ok {
				variant = core.NewQVariant1(str)
			} else if iwidget, ok := v.(IWidget); ok {
				widgetItem := tableViewModelData{
					row:    i,
					column: j,
					widget: iwidget.getWidget().(*WidgetImpl).widget,
				}
				widgetData = append(widgetData, widgetItem)
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
	t.model.widgetData = widgetData
}

func (t *TableViewImpl) loadWidgets() {
	for _, item := range t.model.widgetData {
		t.tableView.SetIndexWidget(t.model.Index(item.row, item.column, nil), item.widget)
	}
}
