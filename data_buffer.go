package malgo
import (
	"fmt"
	"unsafe"
)

/*
#include "miniaudio.h"
*/
import "C"

// a DataBuffer is a type that encapsulates a pInput or pInput buffer in a device data callback and lets you easily convert it to different audio formats
// any slices returned by calling methods on this type point to the Miniaudio-provided buffer and can be used to read or overwrite the contained data. Never append to any of these slices as that will cause reallocation, you should always just read or replace the existing data
// all the conversion methods except for Bytes will panic if the format this callbacks associated device is using is not the specified type. This can be achieved by setting the format in the DeviceConfig, either playback or capture depending on weather this is the output or input buffer.
type DataBuffer struct {
	data unsafe.Pointer
	length int
	format Format
	sampleSize int
}
// returns a slice of the raw underlying bytes of this DataBuffer, no matter the format
// if you just want to move data between DataBuffers such as in duplex mode, this is for you
func (self *DataBuffer) Bytes() []byte {
	return (*[1 << 30]byte)(self.data)[0 : uintptr(self.length)]
}
// returns this buffers format type
func (self *DataBuffer) Format() Format {
	return self.format
}
func (self *DataBuffer) ensureFormat(format Format) {
	if self.format != format {
		panic(fmt.Errorf("Malgo: attempt to convert %v data buffer into %v", self.format, format))
	}
}


// returns a slice of uint8 pointing to the data in this buffer. It panics if this buffers format is not FormatU8
func (self *DataBuffer) Uint8() []uint8 {
	self.ensureFormat(FormatU8)
	return (*[1 << 30]uint8)(self.data)[0 : uintptr(self.length)]
}
func (self *DataBuffer) divideLength(size uintptr) int {
	// if if length % size != 0, something is very wrong
	if self.length % int(size) != 0 {
		panic(fmt.Errorf("Malgo: %v is not a multiple of the size of %v, %v", self.length % int(size), size, self.format))
	}
	return self.length / int(size)
}
// returns a slice of int16 pointing to the data in this buffer. It panics if this buffers format is not FormatS16
func (self *DataBuffer) Int16() []int16 {
	self.ensureFormat(FormatS16)
	return (*[1 << 30]int16)(self.data)[0 : self.divideLength(unsafe.Sizeof(float32(0)))]
}
// returns a slice of int32 pointing to the data in this buffer. It panics if this buffers format is not FormatS32
func (self *DataBuffer) Int32() []int32 {
	self.ensureFormat(FormatS32)
	return (*[1 << 30]int32)(self.data)[0 : self.divideLength(unsafe.Sizeof(int32(0)))]
}
// returns a slice of float32 pointing to the data in this buffer. It panics if this buffers format is not FormatF32
func (self *DataBuffer) Float32() []float32 {
	self.ensureFormat(FormatF32)
	return (*[1 << 30]float32)(self.data)[0 : self.divideLength(unsafe.Sizeof(float32(0)))]
}



