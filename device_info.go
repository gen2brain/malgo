package malgo

// #include "malgo.h"
import "C"
import (
	"encoding/hex"
	"fmt"
	"unsafe"
)

// DeviceID type.
type DeviceID [C.sizeof_ma_device_id]byte

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
	Name          string
	IsDefault     bool
	// the formats supported by the device
	Formats       []FormatType
	MinChannels   int
	MaxChannels   int
	MinSampleRate int
	MaxSampleRate int
}
// String returns string.
func deviceInfoFromCRepr(cInfo C.ma_device_info) DeviceInfo {
	info := DeviceInfo {
		ID : *(*DeviceID)(unsafe.Pointer(&cInfo.id)),
		Name: goString(&cInfo.name[0]),
		IsDefault: intToBool(cInfo.isDefault),
		Formats: func() []FormatType {

			var formats = make([]FormatType, int(cInfo.formatCount))
			for pos, i := range cInfo.formats[:int(cInfo.formatCount)] {
				formats[pos] = FormatType(i)
			}
			return formats
		}(),
		MinChannels: int(cInfo.minChannels),
		MaxChannels: int(cInfo.maxChannels),
		MinSampleRate: int(cInfo.minSampleRate),
		MaxSampleRate: int(cInfo.maxSampleRate),
	}
	return info
}

// String returns string.
func (d *DeviceInfo) String() string {
	return fmt.Sprintf("{ID: [%v], Name: %s}", d.ID, d.Name)
}