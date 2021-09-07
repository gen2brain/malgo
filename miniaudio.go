// Package malgo - Mini audio library (miniaudio cgo bindings).
package malgo
import (
	"time"
)


/*
#cgo CFLAGS: -std=gnu99
#cgo ma_debug CFLAGS: -DMA_DEBUG_OUTPUT=1

#cgo linux,!android LDFLAGS: -ldl -lpthread -lm
#cgo openbsd LDFLAGS: -ldl -lpthread -lm
#cgo netbsd LDFLAGS: -ldl -lpthread -lm
#cgo freebsd LDFLAGS: -ldl -lpthread -lm
#cgo android LDFLAGS: -lm

#cgo !noasm,!arm,!arm64 CFLAGS: -msse2
#cgo !noasm,arm,arm64 CFLAGS: -mfpu=neon -mfloat-abi=hard
#cgo noasm CFLAGS: -DMA_NO_SSE2 -DMA_NO_AVX2 -DMA_NO_AVX512 -DMA_NO_NEON

#include "malgo.h"
*/
import "C"

// SampleSizeInBytes retrieves the size of a sample in bytes for the given format.
func SampleSizeInBytes(format Format) int {
	cformat := (C.ma_format)(format)
	ret := C.ma_get_bytes_per_sample(cformat)
	return int(ret)
}
// returns the time needed to play the specified amount of frames, with the provided sample rate
func FramesToDuration(frameCount int, sampleRate int) time.Duration {
	// a frame is 2 samples, sampleRate is how many frames should be played per second and a sample is a primitive type (I.E. float32, int16) ETC. Knowing this, we can calculate bytesPerFrame, then figure out how many frames are contained in length bytes of data with length / bytesPerFrame. Then divide the result by sampleRate and this returns a float64 representing seconds.
	seconds := float64(frameCount) / float64(sampleRate)
	return time.Duration(float64(time.Second) * seconds)
}
// returns the time needed to play the specified amount of bytes, with the provided sample rate, channels and format
func BytesToDuration(byteCount int, sampleRate int, channels int, format Format) time.Duration {
	frames := BytesToFrames(byteCount, sampleRate, channels, format)
	seconds := float64(frames) / float64(sampleRate)
	return time.Duration(float64(time.Second) * seconds)
}
func DurationToFrames(d time.Duration, sampleRate int) int {
	return int(d.Seconds() * float64(sampleRate))
}
func BytesToFrames(byteCount int, sampleRate int, channels int, format Format) int {
	// the number of frames in bytesAmount is bytesAmount / sizeof_frame, and sizeof_frame is ma_get_bytes_per_sample(format) * channels
	bytesPerFrame := SampleSizeInBytes(format) * channels
	return byteCount / bytesPerFrame
}
func FramesToBytes(frameCount int, sampleRate int, channels int, format Format) int {
	bytesPerFrame := SampleSizeInBytes(format) * channels
	return frameCount * bytesPerFrame
}

const (
	rawContextConfigSize = C.sizeof_ma_context_config
	rawDeviceInfoSize    = C.sizeof_ma_device_info
	rawDeviceConfigSize  = C.sizeof_ma_device_config
)
