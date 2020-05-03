package malgo

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestRawSizes(t *testing.T) {
	assert.EqualValues(t, rawContextConfigSize, unsafe.Sizeof(ContextConfig{}), "ContextConfig size mismatch")
	assert.EqualValues(t, rawDeviceConfigSize, unsafe.Sizeof(DeviceConfig{}), "DeviceConfig size mismatch")
	assert.EqualValues(t, rawDeviceInfoSize, unsafe.Sizeof(DeviceInfo{}), "DeviceInfo size mismatch")
}
