package qt

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
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

type SplashCtxObj_ITF interface {
	std_core.QObject_ITF
	SplashCtxObj_PTR() *SplashCtxObj
}

func (ptr *SplashCtxObj) SplashCtxObj_PTR() *SplashCtxObj {
	return ptr
}

func (ptr *SplashCtxObj) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *SplashCtxObj) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromSplashCtxObj(ptr SplashCtxObj_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.SplashCtxObj_PTR().Pointer()
	}
	return nil
}

func NewSplashCtxObjFromPointer(ptr unsafe.Pointer) (n *SplashCtxObj) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(SplashCtxObj)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *SplashCtxObj:
			n = deduced

		case *std_core.QObject:
			n = &SplashCtxObj{QObject: *deduced}

		default:
			n = new(SplashCtxObj)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *SplashCtxObj) Init() { this.init() }

//export callbackSplashCtxObj6592ff_Constructor
func callbackSplashCtxObj6592ff_Constructor(ptr unsafe.Pointer) {
	this := NewSplashCtxObjFromPointer(ptr)
	qt.Register(ptr, this)
	this.init()
}

//export callbackSplashCtxObj6592ff_Status
func callbackSplashCtxObj6592ff_Status(ptr unsafe.Pointer) C.struct_Moc_PackedString {
	if signal := qt.GetSignal(ptr, "status"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := NewSplashCtxObjFromPointer(ptr).StatusDefault()
	return C.struct_Moc_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *SplashCtxObj) ConnectStatus(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "status"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "status", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "status", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SplashCtxObj) DisconnectStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "status")
	}
}

func (ptr *SplashCtxObj) Status() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.SplashCtxObj6592ff_Status(ptr.Pointer()))
	}
	return ""
}

func (ptr *SplashCtxObj) StatusDefault() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.SplashCtxObj6592ff_StatusDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackSplashCtxObj6592ff_SetStatus
func callbackSplashCtxObj6592ff_SetStatus(ptr unsafe.Pointer, status C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "setStatus"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(status))
	} else {
		NewSplashCtxObjFromPointer(ptr).SetStatusDefault(cGoUnpackString(status))
	}
}

func (ptr *SplashCtxObj) ConnectSetStatus(f func(status string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setStatus"); signal != nil {
			f := func(status string) {
				(*(*func(string))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "setStatus", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setStatus", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SplashCtxObj) DisconnectSetStatus() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setStatus")
	}
}

func (ptr *SplashCtxObj) SetStatus(status string) {
	if ptr.Pointer() != nil {
		var statusC *C.char
		if status != "" {
			statusC = C.CString(status)
			defer C.free(unsafe.Pointer(statusC))
		}
		C.SplashCtxObj6592ff_SetStatus(ptr.Pointer(), C.struct_Moc_PackedString{data: statusC, len: C.longlong(len(status))})
	}
}

func (ptr *SplashCtxObj) SetStatusDefault(status string) {
	if ptr.Pointer() != nil {
		var statusC *C.char
		if status != "" {
			statusC = C.CString(status)
			defer C.free(unsafe.Pointer(statusC))
		}
		C.SplashCtxObj6592ff_SetStatusDefault(ptr.Pointer(), C.struct_Moc_PackedString{data: statusC, len: C.longlong(len(status))})
	}
}

//export callbackSplashCtxObj6592ff_StatusChanged
func callbackSplashCtxObj6592ff_StatusChanged(ptr unsafe.Pointer, status C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "statusChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(status))
	}

}

func (ptr *SplashCtxObj) ConnectStatusChanged(f func(status string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "statusChanged") {
			C.SplashCtxObj6592ff_ConnectStatusChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "statusChanged"); signal != nil {
			f := func(status string) {
				(*(*func(string))(signal))(status)
				f(status)
			}
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "statusChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SplashCtxObj) DisconnectStatusChanged() {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "statusChanged")
	}
}

func (ptr *SplashCtxObj) StatusChanged(status string) {
	if ptr.Pointer() != nil {
		var statusC *C.char
		if status != "" {
			statusC = C.CString(status)
			defer C.free(unsafe.Pointer(statusC))
		}
		C.SplashCtxObj6592ff_StatusChanged(ptr.Pointer(), C.struct_Moc_PackedString{data: statusC, len: C.longlong(len(status))})
	}
}

func SplashCtxObj_QRegisterMetaType() int {
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType()))
}

func (ptr *SplashCtxObj) QRegisterMetaType() int {
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType()))
}

func SplashCtxObj_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType2(typeNameC)))
}

func (ptr *SplashCtxObj) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType2(typeNameC)))
}

func SplashCtxObj_QmlRegisterType() int {
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType()))
}

func (ptr *SplashCtxObj) QmlRegisterType() int {
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType()))
}

func SplashCtxObj_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SplashCtxObj) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
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
	return int(int32(C.SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *SplashCtxObj) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SplashCtxObj6592ff___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SplashCtxObj) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SplashCtxObj) __children_newList() unsafe.Pointer {
	return C.SplashCtxObj6592ff___children_newList(ptr.Pointer())
}

func (ptr *SplashCtxObj) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.SplashCtxObj6592ff___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *SplashCtxObj) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *SplashCtxObj) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.SplashCtxObj6592ff___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *SplashCtxObj) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SplashCtxObj6592ff___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SplashCtxObj) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SplashCtxObj) __findChildren_newList() unsafe.Pointer {
	return C.SplashCtxObj6592ff___findChildren_newList(ptr.Pointer())
}

func (ptr *SplashCtxObj) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SplashCtxObj6592ff___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SplashCtxObj) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SplashCtxObj) __findChildren_newList3() unsafe.Pointer {
	return C.SplashCtxObj6592ff___findChildren_newList3(ptr.Pointer())
}

func (ptr *SplashCtxObj) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.SplashCtxObj6592ff___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *SplashCtxObj) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *SplashCtxObj) __qFindChildren_newList2() unsafe.Pointer {
	return C.SplashCtxObj6592ff___qFindChildren_newList2(ptr.Pointer())
}

func NewSplashCtxObj(parent std_core.QObject_ITF) *SplashCtxObj {
	tmpValue := NewSplashCtxObjFromPointer(C.SplashCtxObj6592ff_NewSplashCtxObj(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackSplashCtxObj6592ff_DestroySplashCtxObj
func callbackSplashCtxObj6592ff_DestroySplashCtxObj(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~SplashCtxObj"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSplashCtxObjFromPointer(ptr).DestroySplashCtxObjDefault()
	}
}

func (ptr *SplashCtxObj) ConnectDestroySplashCtxObj(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~SplashCtxObj"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~SplashCtxObj", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~SplashCtxObj", unsafe.Pointer(&f))
		}
	}
}

func (ptr *SplashCtxObj) DisconnectDestroySplashCtxObj() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~SplashCtxObj")
	}
}

func (ptr *SplashCtxObj) DestroySplashCtxObj() {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_DestroySplashCtxObj(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *SplashCtxObj) DestroySplashCtxObjDefault() {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_DestroySplashCtxObjDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSplashCtxObj6592ff_ChildEvent
func callbackSplashCtxObj6592ff_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewSplashCtxObjFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *SplashCtxObj) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackSplashCtxObj6592ff_ConnectNotify
func callbackSplashCtxObj6592ff_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSplashCtxObjFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SplashCtxObj) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSplashCtxObj6592ff_CustomEvent
func callbackSplashCtxObj6592ff_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewSplashCtxObjFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *SplashCtxObj) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackSplashCtxObj6592ff_DeleteLater
func callbackSplashCtxObj6592ff_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewSplashCtxObjFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *SplashCtxObj) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackSplashCtxObj6592ff_Destroyed
func callbackSplashCtxObj6592ff_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackSplashCtxObj6592ff_DisconnectNotify
func callbackSplashCtxObj6592ff_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewSplashCtxObjFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *SplashCtxObj) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackSplashCtxObj6592ff_Event
func callbackSplashCtxObj6592ff_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSplashCtxObjFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *SplashCtxObj) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SplashCtxObj6592ff_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackSplashCtxObj6592ff_EventFilter
func callbackSplashCtxObj6592ff_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewSplashCtxObjFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *SplashCtxObj) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.SplashCtxObj6592ff_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackSplashCtxObj6592ff_ObjectNameChanged
func callbackSplashCtxObj6592ff_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackSplashCtxObj6592ff_TimerEvent
func callbackSplashCtxObj6592ff_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewSplashCtxObjFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *SplashCtxObj) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.SplashCtxObj6592ff_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
