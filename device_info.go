package malgo

// #include "malgo.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

// DeviceID type.
type DeviceID [unsafe.Sizeof(C.ma_device_id{})]byte

// String returns the string representation of the identifier.
// It is the hexadecimal form of the underlying bytes of a minimum length of 2 digits, with trailing zeroes removed.
func (d DeviceID) String() string {
	displayLen := len(d)
	for (displayLen > 1) && (d[displayLen-1] == 0) {
		displayLen--
	}
	return hex.EncodeToString(d[:displayLen])
}

func (d *DeviceID) cptr() *C.ma_device_id {
	return (*C.ma_device_id)(unsafe.Pointer(d))
}

// DeviceInfo type.
type DeviceInfo struct {
	ID            DeviceID
	name          [256]byte
	FormatCount   uint32
	Formats       [6]uint32
	MinChannels   uint32
	MaxChannels   uint32
	MinSampleRate uint32
	MaxSampleRate uint32
}

// Name returns the name of the device.
func (d *DeviceInfo) Name() string {
	return string(d.name[:])
}

// String returns string.
func (d *DeviceInfo) String() string {
	return fmt.Sprintf("{ID: [%v], Name: %s}", d.ID, d.Name())
}

func deviceInfoFromPointer(ptr unsafe.Pointer) DeviceInfo {
	return *(*DeviceInfo)(ptr)
}
