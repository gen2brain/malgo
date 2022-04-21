package malgo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestRawSizes(t *testing.T) {
	assertEqual(t, rawContextConfigSize, unsafe.Sizeof(ContextConfig{}), "ContextConfig size mismatch")
	assertEqual(t, rawDeviceConfigSize, unsafe.Sizeof(DeviceConfig{}), "DeviceConfig size mismatch")
	assertEqual(t, rawDeviceInfoSize, unsafe.Sizeof(DeviceInfo{}), "DeviceInfo size mismatch")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
