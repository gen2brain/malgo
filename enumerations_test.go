package malgo_test

import (
	"github.com/gen2brain/malgo"
	"testing"
)

func TestDataFormatFlag(t *testing.T) {
	t.Run("Set", func(t *testing.T) {
		var flag malgo.DataFormatFlag
		flag.Set(malgo.DataFormatFlagExclusiveMode)
		if !flag.Has(malgo.DataFormatFlagExclusiveMode) {
			t.Errorf("Expected flag to have DataFormatFlagExclusiveMode set")
		}
	})

	t.Run("Clear", func(t *testing.T) {
		var flag malgo.DataFormatFlag
		flag.Set(malgo.DataFormatFlagExclusiveMode)
		flag.Clear(malgo.DataFormatFlagExclusiveMode)
		if flag.Has(malgo.DataFormatFlagExclusiveMode) {
			t.Errorf("Expected flag to have DataFormatFlagExclusiveMode cleared")
		}
	})

	t.Run("Has", func(t *testing.T) {
		var flag malgo.DataFormatFlag
		if flag.Has(malgo.DataFormatFlagExclusiveMode) {
			t.Errorf("Expected flag to not have DataFormatFlagExclusiveMode set")
		}
		flag.Set(malgo.DataFormatFlagExclusiveMode)
		if !flag.Has(malgo.DataFormatFlagExclusiveMode) {
			t.Errorf("Expected flag to have DataFormatFlagExclusiveMode set")
		}
	})

	t.Run("List", func(t *testing.T) {
		var flag malgo.DataFormatFlag
		flag.Set(malgo.DataFormatFlagExclusiveMode)
		list := flag.List()
		if len(list) != 1 || list[0] != "ExclusiveMode" {
			t.Errorf("Expected list to contain 'ExclusiveMode'")
		}
	})
}
