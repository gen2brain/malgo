package malgo_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/gen2brain/malgo"
)

var testWithHardware = flag.Bool("malgo.hardware", false, "run tests with expecting hardware")

func TestContextLifecycle(t *testing.T) {
	config := malgo.ContextConfig{ThreadPriority: malgo.ThreadPriorityNormal}

	ctx, err := malgo.InitContext(nil, config, nil)
	assertNil(t, err, "No error expected initializing context")
	assertNotNil(t, ctx, "Context instance expected")
	assertNotEqual(t, malgo.Context{}, ctx.Context, "Context value expected")

	err = ctx.Uninit()
	assertNil(t, err, "No error expected uninitializing")

	ctx.Free()
	assertEqual(t, malgo.Context{}, ctx.Context, "Expected context value to be reset")
}

func TestContextDeviceEnumeration(t *testing.T) {
	if *testWithHardware {
		t.Log("Running test expecting devices\n")
	}

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	assertNil(t, err, "No error expected initializing context")
	defer func() {
		err := ctx.Uninit()
		assertNil(t, err, "No error expected uninitializing")
		ctx.Free()
	}()

	playbackDevices, err := ctx.Devices(malgo.Playback)
	assertNil(t, err, "No error expected querying playback devices")
	if *testWithHardware {
		assertTrue(t, len(playbackDevices) > 0, "No playback devices found")
	}

	captureDevices, err := ctx.Devices(malgo.Capture)
	assertNil(t, err, "No error expected querying capture devices")
	if *testWithHardware {
		assertTrue(t, len(captureDevices) > 0, "No capture devices found")
	}
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

func assertNotEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v == %v", a, b)
	}
	t.Fatal(message)
}

func assertNil(t *testing.T, v interface{}, message string) {
	if v == nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("expected nil, got %#v", v)
	}
	t.Fatal(message)
}

func assertNotNil(t *testing.T, v interface{}, message string) {
	if v != nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("expected value not to be nil")
	}
	t.Fatal(message)
}

func assertTrue(t *testing.T, v bool, message string) {
	if v {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("should be true")
	}
	t.Fatal(message)
}
