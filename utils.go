// a bunch of things to help with working with CGO and CGO types

package malgo
import (
	"unsafe"
)
/*
#include <stdlib.h>
#include "miniaudio.h"
*/
import "C"

func goString(ptr *C.char) string {
	if ptr == nil {
		return ""
	} else {
		return C.GoString(ptr)
	}
}
func cString(value string) *C.char {
	if len(value) == 0 {
		return nil
	} else {
		return C.CString(value)
	}
}
func freeCString(ptr *C.char) {
	if ptr == nil {
		return
	} else {
		C.free(unsafe.Pointer(ptr))
	}
}
func intToBool(value C.ma_bool32) bool {
	return value != 0
}
func boolToInt(value bool) C.ma_bool32 {
	switch value {
		case false: return 0
		case true: return 1
	}
	return 0
}
func int8ToBool(value C.ma_bool8) bool {
	return value != 0
}
func boolToInt8(value bool) C.ma_bool8 {
	switch value {
		case false: return 0
		default: return 1
	}
}
// a list of pointers to C allocated memory that makes it easy to associate a bunch of pointers with something like a context and then free them all at once
type pointerList []unsafe.Pointer
// allocates a new cString and adds it to self
func (self *pointerList) cString(value string) *C.char {
	ptr := cString(value)
	self.addPointer(unsafe.Pointer(ptr))
	return ptr
}
func (self *pointerList) addPointer(ptr unsafe.Pointer ) {
	if ptr != nil {
		*self = append(*self, ptr)
	}
}
// allocates value into C bytes, adds the ptr to self and returns ptr
func (self *pointerList) cbytes(value []byte) unsafe.Pointer {
	ptr := C.CBytes(value)
	self.addPointer(ptr)
	return ptr
}
// empties self and frees all the associated pointers
// self is still useable after calling free
func (self *pointerList) free() {
	for _, i := range *self {
		C.free(i)
	}
	*self = (*self)[:0] // self is now empty
}