// Package malgo - Mini audio library (miniaudio cgo bindings).
package malgo

/*
#cgo CFLAGS: -std=gnu99 -Wno-unused-result
#cgo ma_debug CFLAGS: -DMA_DEBUG_OUTPUT=1

#cgo linux,!android LDFLAGS: -ldl -lpthread -lm
#cgo linux,arm LDFLAGS: -latomic
#cgo openbsd LDFLAGS: -lpthread -lm
#cgo netbsd LDFLAGS: -ldl -lpthread -lm
#cgo freebsd LDFLAGS: -ldl -lpthread -lm
#cgo android LDFLAGS: -lm
#cgo ios CFLAGS: -x objective-c
#cgo ios LDFLAGS: -framework CoreFoundation -framework AVFAudio -framework CoreAudio -framework AudioToolbox

#cgo !noasm,!arm,!arm64 CFLAGS: -msse2
#cgo !noasm,arm,arm64 CFLAGS: -mfpu=neon -mfloat-abi=hard
#cgo noasm CFLAGS: -DMA_NO_SSE2 -DMA_NO_AVX2 -DMA_NO_AVX512 -DMA_NO_NEON

#include "malgo.h"
*/
import "C"

// SampleSizeInBytes retrieves the size of a sample in bytes for the given format.
func SampleSizeInBytes(format FormatType) int {
	cformat := (C.ma_format)(format)
	ret := C.ma_get_bytes_per_sample(cformat)
	return int(ret)
}

const (
	rawDeviceInfoSize = C.sizeof_ma_device_info
)
