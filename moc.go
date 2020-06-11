package faithtop

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoFreePacked(ptr unsafe.Pointer) { std_core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_Moc_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type qmlBridge_ITF interface {
	std_core.QObject_ITF
	qmlBridge_PTR() *qmlBridge
}

func (ptr *qmlBridge) qmlBridge_PTR() *qmlBridge {
	return ptr
}

func (ptr *qmlBridge) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *qmlBridge) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQmlBridge(ptr qmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.qmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) (n *qmlBridge) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(qmlBridge)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *qmlBridge:
			n = deduced

		case *std_core.QObject:
			n = &qmlBridge{QObject: *deduced}

		default:
			n = new(qmlBridge)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackqmlBridgefdc0e4_Constructor
func callbackqmlBridgefdc0e4_Constructor(ptr unsafe.Pointer) {
	this := NewQmlBridgeFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackqmlBridgefdc0e4_SendToQml
func callbackqmlBridgefdc0e4_SendToQml(ptr unsafe.Pointer, fn C.uintptr_t) {
	var fnD func()
	if fnI, ok := qt.ReceiveTemp(unsafe.Pointer(uintptr(fn))); ok {
		qt.UnregisterTemp(unsafe.Pointer(uintptr(fn)))
		fnD = (*(*func())(fnI))
	}
	if signal := qt.GetSignal(ptr, "sendToQml"); signal != nil {
		(*(*func(func()))(signal))(fnD)
	}

}

func (ptr *qmlBridge) ConnectSendToQml(f func(fn func())) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "sendToQml") {
			C.qmlBridgefdc0e4_ConnectSendToQml(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "sendToQml")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "sendToQml"); signal != nil {
			f := func(fn func()) {
				(*(*func(func()))(signal))(fn)
				f(fn)
			}
			qt.ConnectSignal(ptr.Pointer(), "sendToQml", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "sendToQml", unsafe.Pointer(&f))
		}
	}
}

func (ptr *qmlBridge) DisconnectSendToQml() {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4_DisconnectSendToQml(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "sendToQml")
	}
}

func (ptr *qmlBridge) SendToQml(fn func()) {
	if ptr.Pointer() != nil {
		qt.RegisterTemp(unsafe.Pointer(&fn), unsafe.Pointer(&fn))
		C.qmlBridgefdc0e4_SendToQml(ptr.Pointer(), C.uintptr_t(uintptr(unsafe.Pointer(&fn))))
	}
}

func qmlBridge_QRegisterMetaType() int {
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType()))
}

func (ptr *qmlBridge) QRegisterMetaType() int {
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType()))
}

func qmlBridge_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType2(typeNameC)))
}

func (ptr *qmlBridge) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType2(typeNameC)))
}

func qmlBridge_QmlRegisterType() int {
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType()))
}

func (ptr *qmlBridge) QmlRegisterType() int {
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType()))
}

func qmlBridge_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *qmlBridge) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *qmlBridge) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.qmlBridgefdc0e4___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *qmlBridge) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *qmlBridge) __children_newList() unsafe.Pointer {
	return C.qmlBridgefdc0e4___children_newList(ptr.Pointer())
}

func (ptr *qmlBridge) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.qmlBridgefdc0e4___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *qmlBridge) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *qmlBridge) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.qmlBridgefdc0e4___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *qmlBridge) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.qmlBridgefdc0e4___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *qmlBridge) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *qmlBridge) __findChildren_newList() unsafe.Pointer {
	return C.qmlBridgefdc0e4___findChildren_newList(ptr.Pointer())
}

func (ptr *qmlBridge) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.qmlBridgefdc0e4___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *qmlBridge) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *qmlBridge) __findChildren_newList3() unsafe.Pointer {
	return C.qmlBridgefdc0e4___findChildren_newList3(ptr.Pointer())
}

func NewQmlBridge(parent std_core.QObject_ITF) *qmlBridge {
	qmlBridge_QRegisterMetaType()
	tmpValue := NewQmlBridgeFromPointer(C.qmlBridgefdc0e4_NewQmlBridge(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackqmlBridgefdc0e4_DestroyQmlBridge
func callbackqmlBridgefdc0e4_DestroyQmlBridge(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~qmlBridge"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DestroyQmlBridgeDefault()
	}
}

func (ptr *qmlBridge) ConnectDestroyQmlBridge(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~qmlBridge"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~qmlBridge", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~qmlBridge", unsafe.Pointer(&f))
		}
	}
}

func (ptr *qmlBridge) DisconnectDestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~qmlBridge")
	}
}

func (ptr *qmlBridge) DestroyQmlBridge() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.qmlBridgefdc0e4_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *qmlBridge) DestroyQmlBridgeDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.qmlBridgefdc0e4_DestroyQmlBridgeDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackqmlBridgefdc0e4_ChildEvent
func callbackqmlBridgefdc0e4_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *qmlBridge) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackqmlBridgefdc0e4_ConnectNotify
func callbackqmlBridgefdc0e4_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *qmlBridge) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackqmlBridgefdc0e4_CustomEvent
func callbackqmlBridgefdc0e4_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *qmlBridge) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackqmlBridgefdc0e4_DeleteLater
func callbackqmlBridgefdc0e4_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *qmlBridge) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.qmlBridgefdc0e4_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackqmlBridgefdc0e4_Destroyed
func callbackqmlBridgefdc0e4_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackqmlBridgefdc0e4_DisconnectNotify
func callbackqmlBridgefdc0e4_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *qmlBridge) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackqmlBridgefdc0e4_Event
func callbackqmlBridgefdc0e4_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *qmlBridge) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.qmlBridgefdc0e4_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackqmlBridgefdc0e4_EventFilter
func callbackqmlBridgefdc0e4_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *qmlBridge) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.qmlBridgefdc0e4_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackqmlBridgefdc0e4_ObjectNameChanged
func callbackqmlBridgefdc0e4_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackqmlBridgefdc0e4_TimerEvent
func callbackqmlBridgefdc0e4_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *qmlBridge) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.qmlBridgefdc0e4_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

func init() {
}
