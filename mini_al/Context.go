package mini_al

// #include "malgo.h"
import "C"
import (
	"reflect"
	"sync"
	"unsafe"
)

// LogProc type.
type LogProc func(message string)

// AlsaContextConfig type.
type AlsaContextConfig struct {
	UseVerboseDeviceEnumeration uint32
}

// PulseContextConfig type.
type PulseContextConfig struct {
	PApplicationName *byte
	PServerName      *byte
	// Enables autospawning of the PulseAudio daemon if necessary.
	TryAutoSpawn uint32
	// Padding
	_ [4]byte
}

// JackContextConfig type.
type JackContextConfig struct {
	PClientName    *byte
	TryStartServer uint32
	// Padding
	_ [4]byte
}

// ContextConfig type.
type ContextConfig struct {
	onLog          uintptr
	ThreadPriority ThreadPriority
	Alsa           AlsaContextConfig
	Pulse          PulseContextConfig
	Jack           JackContextConfig
}

func (d *ContextConfig) cptr() *C.mal_context_config {
	return (*C.mal_context_config)(unsafe.Pointer(d))
}

func contextConfigFromPointer(ptr unsafe.Pointer) ContextConfig {
	return *(*ContextConfig)(ptr)
}

// Context is used for selecting and initializing the relevant backends.
type Context uintptr

// DefaultContext is an unspecified context. It can be used to initialize a streaming
// function with implicit context defaults.
const DefaultContext Context = 0

func (ctx Context) cptr() *C.mal_context {
	return (*C.mal_context)(unsafe.Pointer(ctx))
}

// Uninit uninitializes a context.
// Results are undefined if you call this while any device created by this context is still active.
func (ctx Context) Uninit() error {
	result := C.mal_context_uninit(ctx.cptr())
	return errorFromResult(Result(result))
}

var contextMutex sync.Mutex
var logProcMap map[*C.mal_context]LogProc

// SetLogProc sets the logging callback for the context.
func (ctx Context) SetLogProc(proc LogProc) {
	contextMutex.Lock()
	defer contextMutex.Unlock()
	ptr := ctx.cptr()
	if proc != nil {
		logProcMap[ptr] = proc
	} else {
		delete(logProcMap, ptr)
	}
}

//export goLogCallback
func goLogCallback(pContext *C.mal_context, pDevice *C.mal_device, message *C.char) {
	proc := logProcMap[pContext]
	if proc != nil {
		proc(C.GoString(message))
	}
}

// AllocatedContext is a Context that has been created by the application.
// It must be freed after use in order to release resources.
type AllocatedContext struct {
	Context
}

// InitContext creates and initializes a context.
// When the application no longer needs the context instance, it needs to call Free() .
func InitContext(backends []Backend, config ContextConfig, logProc LogProc) (*AllocatedContext, error) {
	config.onLog = uintptr(C.goLogCallbackPointer())
	ctx := AllocatedContext{Context: Context(C.mal_malloc(C.goContextSize()))}
	if ctx.Context == 0 {
		return nil, ErrOutOfMemory
	}
	ctx.SetLogProc(logProc)

	backendsArg := (*C.mal_backend)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&backends)).Data))
	backendCountArg := (C.mal_uint32)(len(backends))

	result := C.mal_context_init(backendsArg, backendCountArg, config.cptr(), ctx.cptr())
	err := errorFromResult(Result(result))
	if err != nil {
		ctx.SetLogProc(nil)
		ctx.Free()
		return nil, err
	}
	return &ctx, nil
}

// Free must be called when the allocated data is no longer used.
// This function must only be called for an uninitialized context.
func (ctx *AllocatedContext) Free() {
	if ctx.Context == 0 {
		return
	}
	ctx.SetLogProc(nil)
	C.mal_free(unsafe.Pointer(ctx.cptr()))
	ctx.Context = 0
}
