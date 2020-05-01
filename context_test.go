package malgo_test

import (
	"flag"
	"testing"

	"github.com/gen2brain/malgo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testWithHardware = flag.Bool("malgo.hardware", false, "run tests with expecting hardware")

func TestContextLifecycle(t *testing.T) {
	config := malgo.ContextConfig{ThreadPriority: malgo.ThreadPriorityNormal}
	backends := []malgo.Backend{malgo.BackendNull}

	ctx, err := malgo.InitContext(backends, config, nil)
	require.Nil(t, err, "No error expected initializing context")
	require.NotNil(t, ctx, "Context instance expected")
	assert.NotEqual(t, malgo.Context(0), ctx.Context, "Context value expected")

	err = ctx.Uninit()
	assert.Nil(t, err, "No error expected uninitializing")

	ctx.Free()
	assert.Equal(t, malgo.Context(0), ctx.Context, "Expected context value to be reset")
}

func TestContextDeviceEnumeration(t *testing.T) {
	if *testWithHardware {
		t.Log("Running test expecting devices\n")
	}

	ctx, err := malgo.InitContext(nil, malgo.ContextConfig{}, nil)
	require.Nil(t, err, "No error expected initializing context")
	defer func() {
		err := ctx.Uninit()
		assert.Nil(t, err, "No error expected uninitializing")
		ctx.Free()
	}()

	playbackDevices, err := ctx.Devices(malgo.Playback)
	assert.Nil(t, err, "No error expected querying playback devices")
	if *testWithHardware {
		assert.True(t, len(playbackDevices) > 0, "No playback devices found")
	}

	captureDevices, err := ctx.Devices(malgo.Capture)
	assert.Nil(t, err, "No error expected querying capture devices")
	if *testWithHardware {
		assert.True(t, len(captureDevices) > 0, "No capture devices found")
	}
}
