// Package mal - Mini audio library (mini_al cgo bindings).
package malgo

/*
#cgo CFLAGS: -std=gnu99
#cgo ma_debug CFLAGS: -DMAL_DEBUG_OUTPUT=1

#cgo linux LDFLAGS: -ldl -lpthread -lm
#cgo openbsd LDFLAGS: -ldl -lpthread -lm
#cgo netbsd LDFLAGS: -ldl -lpthread -lm
#cgo freebsd LDFLAGS: -ldl -lpthread -lm
#cgo android LDFLAGS: -lOpenSLES

#cgo !noasm,!arm,!arm64 CFLAGS: -msse2
#cgo !noasm,arm,arm64 CFLAGS: -mfpu=neon -mfloat-abi=hard
#cgo noasm CFLAGS: -DMA_NO_SSE2 -DMA_NO_AVX2 -DMA_NO_AVX512 -DMA_NO_NEON

#include "malgo.h"
*/
import "C"
import "unsafe"

// SampleSizeInBytes retrieves the size of a sample in bytes for the given format.
func SampleSizeInBytes(format FormatType) int {
	cformat := (C.ma_format)(format)
	ret := C.ma_get_bytes_per_sample(cformat)
	return int(ret)
}

const (
	rawContextConfigSize = unsafe.Sizeof(C.ma_context_config{})
	rawDeviceInfoSize    = unsafe.Sizeof(C.ma_device_info{})
	rawDeviceConfigSize  = unsafe.Sizeof(C.ma_device_config{})
)
