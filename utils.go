package malgo
import (
)
/*
*/
import "C"

func GoString(ptr *C.char) string {
	if ptr == nil {
		return ""
	} else {
		return C.GoString(ptr)
	}
}
func CString(value string) *C.char {
	if len(value) == 0 {
		return nil
	} else {
		return C.CString(value)
	}
}