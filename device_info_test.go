package malgo_test

import (
	"github.com/gen2brain/malgo"
	"testing"
)

func TestDeviceIDString(t *testing.T) {
	// Sample data should be a device ID with the hex value 0102030405 and the string representation should be "0102030405"
	sampleData := malgo.DeviceID{0x01, 0x02, 0x03, 0x04, 0x05} // remaining bytes are zero
	if len(sampleData) != 256 {
		t.Errorf("Expected length of 256, got %d", len(sampleData))
	}

	t.Run("can output string representation of bytes", func(t *testing.T) {
		expected := "0102030405"
		actual := sampleData.String()
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("can parse string into DeviceID", func(t *testing.T) {
		expected := sampleData
		actual, err := malgo.NewDeviceIDFromString("0102030405")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if actual != expected {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})
}
