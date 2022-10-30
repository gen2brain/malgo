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
}

// CoreAudioConfig type.
type CoreAudioConfig struct {
	SessionCategory        IOSSessionCategory
	SessionCategoryOptions IOSSessionCategoryOptions
}

// JackContextConfig type.
type JackContextConfig struct {
	PClientName    *byte
	TryStartServer uint32
}

// ContextConfig type.
type ContextConfig struct {
	LogCallback         *[0]byte
	ThreadPriority      ThreadPriority
	PUserData           *byte
	AllocationCallbacks AllocationCallbacks
	Alsa                AlsaContextConfig
	Pulse               PulseContextConfig
	CoreAudio           CoreAudioConfig
	Jack                JackContextConfig
}

func (d *ContextConfig) toC() (C.ma_context_config, error) {
	ctxConfig := C.ma_context_config_init()
	ctxConfig.threadPriority = C.ma_thread_priority(d.ThreadPriority)
	ctxConfig.pUserData = unsafe.Pointer(d.PUserData)
	ctxConfig.allocationCallbacks.pUserData = unsafe.Pointer(d.AllocationCallbacks.PUserData)
	ctxConfig.allocationCallbacks.onMalloc = d.AllocationCallbacks.OnMalloc
	ctxConfig.allocationCallbacks.onRealloc = d.AllocationCallbacks.OnRealloc
	ctxConfig.allocationCallbacks.onFree = d.AllocationCallbacks.OnFree
	ctxConfig.alsa.useVerboseDeviceEnumeration = C.uint(d.Alsa.UseVerboseDeviceEnumeration)
	ctxConfig.pulse.pApplicationName = (*C.char)(unsafe.Pointer((d.Pulse.PApplicationName)))
	ctxConfig.pulse.pServerName = (*C.char)(unsafe.Pointer((d.Pulse.PServerName)))
	ctxConfig.pulse.tryAutoSpawn = C.uint(d.Pulse.TryAutoSpawn)
	ctxConfig.coreaudio.sessionCategory = C.ma_ios_session_category(d.CoreAudio.SessionCategory)
	ctxConfig.coreaudio.sessionCategoryOptions = C.uint(d.CoreAudio.SessionCategoryOptions)
	ctxConfig.jack.pClientName = (*C.char)(unsafe.Pointer((d.Jack.PClientName)))
	ctxConfig.jack.tryStartServer = C.uint(d.Jack.TryStartServer)

	return ctxConfig, nil
}

// AllocationCallbacks types.
type AllocationCallbacks struct {
	PUserData *byte
	OnMalloc  *[0]byte
	OnRealloc *[0]byte
	OnFree    *[0]byte
}

// Context is used for selecting and initializing the relevant backends.
type Context struct {
	ptr *unsafe.Pointer
}

// DefaultContext is an unspecified context. It can be used to initialize a streaming
// function with implicit context defaults.
var DefaultContext Context = Context{}

func (ctx Context) cptr() *C.ma_context {
	return (*C.ma_context)(*ctx.ptr)
}

// Uninit uninitializes a context.
// Results are undefined if you call this while any device created by this context is still active.
func (ctx Context) Uninit() error {
	result := C.ma_context_uninit(ctx.cptr())
	return errorFromResult(result)
}

// Devices retrieves basic information about every active playback or capture device.
func (ctx Context) Devices(kind DeviceType) ([]DeviceInfo, error) {
	contextMutex.Lock()
	defer contextMutex.Unlock()

	var playbackDevices *C.ma_device_info
	var playbackDeviceCount C.ma_uint32
	var captureDevices *C.ma_device_info
	var captureDeviceCount C.ma_uint32

	result := C.ma_context_get_devices(ctx.cptr(),
		&playbackDevices, &playbackDeviceCount,
		&captureDevices, &captureDeviceCount)
	err := errorFromResult(result)
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
	deviceInfoAddr := unsafe.Pointer(devices)
	for i := 0; i < deviceCount; i++ {
		info[i] = deviceInfoFromPointer(deviceInfoAddr)
		deviceInfoAddr = unsafe.Add(deviceInfoAddr, rawDeviceInfoSize)
	}

	return info, nil
}

// DeviceInfo retrieves information about a device of the given type, with the specified ID and share mode.
func (ctx Context) DeviceInfo(kind DeviceType, id DeviceID, mode ShareMode) (DeviceInfo, error) {
	var info C.ma_device_info

	result := C.ma_context_get_device_info(ctx.cptr(), C.ma_device_type(kind), id.cptr(), &info)
	err := errorFromResult(result)
	if err != nil {
		return DeviceInfo{}, err
	}

	return deviceInfoFromPointer(unsafe.Pointer(&info)), nil
}

var contextMutex sync.Mutex
var logProcMap = make(map[*C.ma_context]LogProc)

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
func goLogCallback(pContext *C.ma_context, message *C.char) {
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
	config *C.ma_context_config
}

// InitContext creates and initializes a context.
// When the application no longer needs the context instance, it needs to call Free() .
func InitContext(backends []Backend, config ContextConfig, logProc LogProc) (*AllocatedContext, error) {
	configC, err := config.toC()
	if err != nil {
		return nil, err
	}

	ptr := C.ma_malloc(C.sizeof_ma_context, nil)
	ctx := AllocatedContext{
		Context: Context{ptr: &ptr},
	}
	if uintptr(*ctx.Context.ptr) == 0 {
		ctx.Free()
		return nil, ErrOutOfMemory
	}

	C.goSetContextConfigCallbacks(&configC, ctx.cptr())
	ctx.SetLogProc(logProc)

	var backendsArg *C.ma_backend
	backendCountArg := (C.ma_uint32)(len(backends))
	if backendCountArg > 0 {
		backendsArg = (*C.ma_backend)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&backends)).Data))
	}

	result := C.ma_context_init(backendsArg, backendCountArg, &configC, ctx.cptr())
	if err := errorFromResult(result); err != nil {
		ctx.SetLogProc(nil)
		ctx.Free()
		return nil, err
	}
	return &ctx, nil
}

// Free must be called when the allocated data is no longer used.
// This function must only be called for an uninitialized context.
func (ctx *AllocatedContext) Free() {
	if ctx.Context.ptr == nil || uintptr(*ctx.Context.ptr) == 0 {
		return
	}
	ctx.SetLogProc(nil)
	C.ma_free(unsafe.Pointer(ctx.cptr()), nil)
	ctx.Context.ptr = nil
}
