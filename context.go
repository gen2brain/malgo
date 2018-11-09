package malgo

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
	_              uintptr
	ThreadPriority ThreadPriority
	Alsa           AlsaContextConfig
	Pulse          PulseContextConfig
	Jack           JackContextConfig
}

func (d *ContextConfig) cptr() *C.mal_context_config {
	return (*C.mal_context_config)(unsafe.Pointer(d))
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

// Devices retrieves basic information about every active playback or capture device.
func (ctx Context) Devices(kind DeviceType) ([]DeviceInfo, error) {
	contextMutex.Lock()
	defer contextMutex.Unlock()

	var playbackDevices *C.mal_device_info
	var playbackDeviceCount C.uint
	var captureDevices *C.mal_device_info
	var captureDeviceCount C.uint

	result := C.mal_context_get_devices(ctx.cptr(),
		&playbackDevices, &playbackDeviceCount,
		&captureDevices, &captureDeviceCount)
	err := errorFromResult(Result(result))
	if err != nil {
		return nil, err
	}
	devices := playbackDevices
	deviceCount := int(playbackDeviceCount)
	if kind == Capture {
		devices = captureDevices
		deviceCount = int(captureDeviceCount)
	}
	info := make([]DeviceInfo, deviceCount)
	deviceInfoAddr := uintptr(unsafe.Pointer(devices))
	for i := 0; i < deviceCount; i++ {
		info[i] = deviceInfoFromPointer(unsafe.Pointer(deviceInfoAddr))
		deviceInfoAddr += rawDeviceInfoSize
	}

	return info, nil
}

var contextMutex sync.Mutex
var logProcMap = make(map[*C.mal_context]LogProc)

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
	contextMutex.Lock()
	callback := logProcMap[pContext]
	contextMutex.Unlock()

	if callback != nil {
		callback(C.GoString(message))
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
	C.goSetContextConfigCallbacks(config.cptr())
	ctx := AllocatedContext{Context: Context(C.mal_malloc(C.size_t(unsafe.Sizeof(C.mal_context{}))))}
	if ctx.Context == 0 {
		return nil, ErrOutOfMemory
	}
	ctx.SetLogProc(logProc)

	var backendsArg *C.mal_backend
	backendCountArg := (C.mal_uint32)(len(backends))
	if backendCountArg > 0 {
		backendsArg = (*C.mal_backend)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&backends)).Data))
	}

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
