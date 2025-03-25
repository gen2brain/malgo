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

func NewDeviceIDFromString(id string) (DeviceID, error) {
	var deviceID DeviceID
	decoded, err := hex.DecodeString(id)
	if err != nil {
		return deviceID, err
	}
	copy(deviceID[:], decoded)
	return deviceID, nil
}

// String returns the string representation of the identifier.
// It is the hexadecimal form of the underlying bytes of a minimum length of 2 digits, with trailing zeroes removed.
func (d *DeviceID) String() string {
	displayLen := len(d)
	for (displayLen > 1) && (d[displayLen-1] == 0) {
		displayLen--
	}
	return hex.EncodeToString(d[:displayLen])
}

func (d *DeviceID) Pointer() unsafe.Pointer {
	return C.CBytes(d[:])
}

func (d *DeviceID) cptr() *C.ma_device_id {
	return (*C.ma_device_id)(unsafe.Pointer(d))
}

// DeviceInfo type.
type DeviceInfo struct {
	ID          DeviceID
	Name        string
	IsDefault   bool
	FormatCount int
	Formats     []DataFormat
}

// String returns string.
func (d *DeviceInfo) String() string {
	return fmt.Sprintf("{ID: [%v], Name: %s}", d.ID, d.Name)
}

func deviceInfoFromPointer(ptr unsafe.Pointer) DeviceInfo {
	device := (*C.ma_device_info)(ptr)
	var newDevice DeviceInfo
	newDevice.ID = DeviceID(device.id)
	newDevice.Name = C.GoString(&device.name[0])
	newDevice.IsDefault = device.isDefault == True
	newDevice.FormatCount = int(device.nativeDataFormatCount)
	newDevice.Formats = make([]DataFormat, newDevice.FormatCount)
	for i := 0; i < int(newDevice.FormatCount); i++ {
		newDevice.Formats[i] = DataFormat{
			Format:     FormatType(device.nativeDataFormats[i].format),
			Channels:   int(device.nativeDataFormats[i].channels),
			SampleRate: int(device.nativeDataFormats[i].sampleRate),
			Flags:      DataFormatFlag(device.nativeDataFormats[i].flags),
		}
	}
	return newDevice
}

type DataFormat struct {
	Format     FormatType
	Channels   int
	SampleRate int
	Flags      DataFormatFlag
}
