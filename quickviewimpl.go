//go:build impl

package faithtop

import (
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type (
	QuickViewImpl struct {
		*WidgetImpl
		quickView *quick.QQuickView
		bridge    *QuickBridge
	}

	QuickBridge struct {
		quick.QQuickItem
		_        func()                               `constructor:"init"`
		_        func(funcName string, args []string) `signal:"callGo,auto"`
		_        func(funcName string, args []string) `signal:"callQml"`
		handlers map[string]reflect.Value
	}
)

func init() {
	newQuickViewImpl = func(win IWindow) IQuickView {
		if win == nil {
			panic("the parent window cannot be nil")
		}
		v := &QuickViewImpl{
			quickView: quick.NewQQuickView(nil),
			bridge:    NewQuickBridge(nil),
		}
		var widget *widgets.QWidget
		widget = widget.CreateWindowContainer(v.quickView, win.(*WindowImpl).window, 0)
		v.WidgetImpl = widgetImplFrom(widget)

		v.quickView.RootContext().SetContextProperty("bridge", v.bridge)

		return v
	}
	QuickBridge_QmlRegisterType2("QuickBridge", 1, 0, "QuickBridge")
}

func (b *QuickBridge) init() {
	b.handlers = make(map[string]reflect.Value)
}
func (b *QuickBridge) callGo(funcName string, args []string) {
	if b.handlers == nil {
		return
	}
	if handler, ok := b.handlers[funcName]; ok {
		t := handler.Type()

		if t.NumIn() > len(args) {
			panic("not enough arguments to call Go function: [" + funcName + "], have " + strconv.Itoa(len(args)) + " but want " + strconv.Itoa(t.NumIn()))
		}

		arguments := []reflect.Value{}
		for i := 0; i < t.NumIn(); i++ {
			argument, e := convertToGoType(args[i], t.In(i))
			if e != nil {
				panic(fmt.Errorf("error when calling Go function [%s]: %w", funcName, e))
			}
			arguments = append(arguments, argument)
		}

		handler.Call(arguments)

	} else {
		panic("funcName [" + funcName + "] is not registered on QuickView")
	}
}

func convertToGoType(s string, target reflect.Type) (reflect.Value, error) {
	switch target.Kind() {
	case reflect.String:
		return reflect.ValueOf(s), nil
	default:
		value := reflect.New(target)
		e := json.Unmarshal([]byte(s), value.Interface())
		if e != nil {
			return value, e
		}
		return value.Elem(), nil
	}
}

func convertToQmlArgument(arg interface{}) string {
	if v, ok := arg.(string); ok {
		return v
	}
	b, _ := json.Marshal(arg)
	return string(b)
}

func (q *QuickViewImpl) Assign(v *IQuickView) IQuickView {
	*v = q
	return q
}

func (q *QuickViewImpl) Source(qmlFile string) IQuickView {
	q.quickView.SetSource(core.NewQUrl3(qmlFile, 0))
	return q
}

func (q *QuickViewImpl) Call(funcName string, args ...interface{}) {
	strs := []string{}
	for _, arg := range args {
		strs = append(strs, convertToQmlArgument(arg))
	}
	q.bridge.CallQml(funcName, strs)
}

func (q *QuickViewImpl) BindFunction(fn interface{}) IQuickView {
	funcName := getFunctionName(fn)
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		panic("BindFunction: the input value is not a function, but [" + t.Kind().String() + "]")
	}
	if _, ok := q.bridge.handlers[funcName]; ok {
		panic("BindFunction: function with '" + funcName + "' duplicated")
	}
	q.bridge.handlers[funcName] = reflect.ValueOf(fn)
	return q
}

func getFunctionName(v interface{}) string {
	t := reflect.ValueOf(v)
	fn := runtime.FuncForPC(t.Pointer()).Name()
	n := strings.TrimPrefix(filepath.Ext(fn), ".")
	return n
}
