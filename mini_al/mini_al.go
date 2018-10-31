// Package mal - Mini audio library (mini_al cgo bindings).
package mini_al

/*
#cgo CFLAGS: -std=gnu99
#cgo mal_debug CFLAGS: -DMAL_DEBUG_OUTPUT=1

#cgo linux LDFLAGS: -ldl -lpthread -lm
#cgo openbsd LDFLAGS: -lpthread -lm -lossaudio
#cgo netbsd LDFLAGS: -lpthread -lm -lossaudio
#cgo freebsd LDFLAGS: -lpthread -lm
#cgo android LDFLAGS: -lOpenSLES

#cgo !noasm,!arm,!arm64 CFLAGS: -msse2
#cgo !noasm,arm,arm64 CFLAGS: -mfpu=neon -mfloat-abi=hard
#cgo noasm CFLAGS: -DMAL_NO_SSE2 -DMAL_NO_AVX -DMAL_NO_AVX512 -DMAL_NO_NEON

#include "malgo.h"
*/
import "C"
import "unsafe"

// SampleSizeInBytes retrieves the size of a sample in bytes for the given format.
func SampleSizeInBytes(format FormatType) int {
	cformat := (C.mal_format)(format)
	ret := C.mal_get_bytes_per_sample(cformat)
	return int(ret)
}

const (
	rawContextConfigSize = unsafe.Sizeof(C.mal_context_config{})
	rawDeviceInfoSize    = unsafe.Sizeof(C.mal_device_info{})
	rawDeviceConfigSize  = unsafe.Sizeof(C.mal_device_config{})
)
