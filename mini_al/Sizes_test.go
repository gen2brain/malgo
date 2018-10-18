package mini_al

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestRawSizes(t *testing.T) {
	assert.Equal(t, unsafe.Sizeof(ContextConfig{}), rawContextConfigSize, "ContextConfig size mismatch")
	assert.Equal(t, unsafe.Sizeof(DeviceConfig{}), rawDeviceConfigSize, "DeviceConfig size mismatch")
	assert.Equal(t, unsafe.Sizeof(DeviceInfo{}), rawDeviceInfoSize, "DeviceInfo size mismatch")
}
