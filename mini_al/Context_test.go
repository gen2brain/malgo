package mini_al_test

import (
	"testing"

	"github.com/gen2brain/malgo/mini_al"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContextLifecycle(t *testing.T) {
	config := mini_al.ContextConfig{ThreadPriority: mini_al.ThreadPriorityNormal}
	backends := []mini_al.Backend{mini_al.BackendNull}

	ctx, err := mini_al.InitContext(backends, config, nil)
	require.Nil(t, err, "No error expected initializing context")
	require.NotNil(t, ctx, "Context instance expected")
	assert.NotEqual(t, mini_al.Context(0), ctx.Context, "Context value expected")

	err = ctx.Uninit()
	assert.Nil(t, err, "No error expected uninitializing")

	ctx.Free()
	assert.Equal(t, mini_al.Context(0), ctx.Context, "Expected context value to be reset")
}
