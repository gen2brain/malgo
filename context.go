package malgo

/*
#include "malgo.h"
static void setGoContextConfigCallbacks(ma_context_config* pConfig, ma_bool8 log) {
	if(log) {
		pConfig->logCallback = goLogCallbackWrapper;
	}
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

// LogProc type.
type LogProc func(message string)

type ContextConfig struct {
	ThreadPriority  int
	ThreadStackSize int
	Alsa struct {
		UseVerboseDeviceEnumeration bool
	}
	Pulse struct {
		ApplicationName string
		ServerName      string
		// Enables autospawning of the PulseAudio daemon if necessary.
		TryAutoSpawn bool
	}
	CoreAudio struct {
		SessionCategory                                  int
		SessionCategoryOptions                           int
		NoAudioSessionActivate, NoAudioSessionDeactivate bool
	}
	Jack struct {
		ClientName     string
		TryStartServer bool
	}
}
type internalContextInfo struct {
	LogProc     LogProc
	memory pointerList
}


// Initializes a new ContextConfig with defaults. You must call this instead of creating a ContextConfig object directly.
func NewContextConfig() ContextConfig {
	cConfig := C.ma_context_config_init()
	config := ContextConfig{}
	config.Alsa.UseVerboseDeviceEnumeration = intToBool(cConfig.alsa.useVerboseDeviceEnumeration)

	config.Pulse.ApplicationName = goString(cConfig.pulse.pApplicationName)
	config.Pulse.ServerName = goString(cConfig.pulse.pServerName)
	config.Pulse.TryAutoSpawn = intToBool(cConfig.pulse.tryAutoSpawn)

	config.CoreAudio.SessionCategory = int(cConfig.coreaudio.sessionCategory)
	config.CoreAudio.SessionCategoryOptions = int(cConfig.coreaudio.sessionCategoryOptions)
	config.CoreAudio.NoAudioSessionActivate = intToBool(cConfig.coreaudio.noAudioSessionActivate)
	config.CoreAudio.NoAudioSessionDeactivate = intToBool(cConfig.coreaudio.noAudioSessionDeactivate)

	config.Jack.ClientName = goString(cConfig.jack.pClientName)
	config.Jack.TryStartServer = intToBool(cConfig.jack.tryStartServer)
	config.ThreadPriority = int(cConfig.threadPriority)
	config.ThreadStackSize = int(cConfig.threadStackSize)
	return config
}

// converts this config to a C.ma_context_config for use in ma_context_init
func (self *ContextConfig) toCRepr() (C.ma_context_config, pointerList) {
	// even if we forget to initialize some fields, telling Miniaudio to reinitialize the config ensures it doesn't break anything
	config := C.ma_context_config_init()
	// a list of pointers that holds any memory allocated by the config like strings and ma_backends ETC. This is returned alongside the new cConfig, is stored in the internalDeviceInfo struct and freed when Context.Free() is called
	var memory pointerList

	config.threadPriority = C.ma_thread_priority(self.ThreadPriority)
	config.threadStackSize = C.size_t(self.ThreadStackSize)
	// use a C function to set the callbacks so the C compiler has a chance to check if the callback types match
	config.alsa.useVerboseDeviceEnumeration = boolToInt(self.Alsa.UseVerboseDeviceEnumeration)

	config.pulse.pApplicationName = memory.cString(self.Pulse.ApplicationName)
	config.pulse.pServerName = memory.cString(self.Pulse.ServerName)
	config.pulse.tryAutoSpawn = boolToInt(self.Pulse.TryAutoSpawn)

	config.coreaudio.sessionCategory = C.ma_ios_session_category(self.CoreAudio.SessionCategory)
	config.coreaudio.sessionCategoryOptions = C.ma_uint32(self.CoreAudio.SessionCategoryOptions)
	config.coreaudio.noAudioSessionActivate = boolToInt(self.CoreAudio.NoAudioSessionActivate)
	config.coreaudio.noAudioSessionDeactivate = boolToInt(self.CoreAudio.NoAudioSessionDeactivate)

	config.jack.pClientName = memory.cString(self.Jack.ClientName)
	config.jack.tryStartServer = boolToInt(self.Jack.TryStartServer)
	return config, memory
}
// Context is used for selecting and initializing the relevant backends.
type Context uintptr

// DefaultContext is an unspecified context. It can be used to initialize a streaming
// function with implicit context defaults.
const DefaultContext Context = 0

func (ctx Context) cptr() *C.ma_context {
	return (*C.ma_context)(unsafe.Pointer(ctx))
}

// Uninit uninitializes a context.
// Results are undefined if you call this while any device created by this context is still active.
func (ctx Context) Uninit() error {
	result := C.ma_context_uninit(ctx.cptr())
	return errorFromResult(Result(result))
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
		info[i] = deviceInfoFromCRepr(*(*C.ma_device_info)(unsafe.Pointer(deviceInfoAddr)))
		deviceInfoAddr += rawDeviceInfoSize
	}

	return info, nil
}

// DeviceInfo retrieves information about a device of the given type, with the specified ID and share mode.
func (ctx Context) DeviceInfo(kind DeviceType, id DeviceID, mode ShareMode) (DeviceInfo, error) {
	var info C.ma_device_info

	result := C.ma_context_get_device_info(ctx.cptr(), C.ma_device_type(kind), id.cptr(), C.ma_share_mode(mode), &info)
	err := errorFromResult(Result(result))
	if err != nil {
		return DeviceInfo{}, err
	}

	return deviceInfoFromCRepr(info), nil
}
// for making some context functions like getDevice threadsafe
var contextMutex sync.Mutex


var contextInfosMutex sync.Mutex
var contextInfos = make(map[*C.ma_context]*internalContextInfo)

//export goLogCallback
func goLogCallback(pContext *C.ma_context, pDevice *C.ma_device, message *C.char) {
	contextInfosMutex.Lock()
	callback := contextInfos[pContext].LogProc
	contextInfosMutex.Unlock()


	callback(goString(message))
}

// AllocatedContext is a Context that has been created by the application.
// It must be freed after use in order to release resources.
type AllocatedContext struct {
	Context
}

// InitContext creates and initializes a context.
// When the application no longer needs the context instance, it needs to call Free() .
func InitContext(backends []Backend, config ContextConfig, logProc LogProc) (*AllocatedContext, error) {
	ctx := AllocatedContext{
		Context: Context(C.ma_malloc(C.sizeof_ma_context, nil)),
	}
	if ctx.Context == 0 {
		return nil, ErrOutOfMemory
	}

	var backendsArg *C.ma_backend
	backendCountArg := (C.ma_uint32)(len(backends))
	cConfig, memory := config.toCRepr()
	C.setGoContextConfigCallbacks(&cConfig, boolToInt8(logProc != nil))
	if backendCountArg > 0 {
		// allocate the backends array in C land and add it to the pointer list
		backendsArg = (*C.ma_backend)(C.malloc(C.size_t(len(backends) * C.sizeof_ma_backend)))
		for i := 0;  i < len(backends); i++ {
			offset := uintptr(i) * C.sizeof_ma_backend
			ptr := (*C.ma_backend)(unsafe.Pointer(uintptr(unsafe.Pointer(backendsArg)) + offset))
		*ptr = C.ma_backend(backends[i])
		}
		memory.addPointer(unsafe.Pointer(backendsArg))
	}
	// we need to store the internalContextInfo before initializing the context because the LogProc can be called during context initialization.
	contextInfosMutex.Lock()
	contextInfos[ctx.cptr()] = &internalContextInfo{LogProc: logProc, memory: memory}
	result := C.ma_context_init(backendsArg, backendCountArg, &cConfig, ctx.cptr())
	err := errorFromResult(Result(result))
	if err != nil {
		C.ma_free(unsafe.Pointer(ctx.cptr()), nil)
		delete(contextInfos, ctx.cptr())
		return nil, err
	}
	contextInfosMutex.Unlock()
	return &ctx, nil
}

// Free must be called when the allocated data is no longer used.
// This function must only be called for an uninitialized context.
func (ctx *AllocatedContext) Free() {
	if ctx.Context == 0 {
		return
	}

	C.ma_free(unsafe.Pointer(ctx.cptr()), nil)
	contextInfosMutex.Lock()
	if info, exists := contextInfos[ctx.cptr()]; exists {
		delete(contextInfos, ctx.cptr())
		info.memory.free()
	}
contextInfosMutex.Unlock()
	ctx.Context = 0
}
