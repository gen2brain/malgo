package mini_al

// #include "malgo.h"
import "C"
import (
	"sync"
	"unsafe"
)

// RecvProc type.
type RecvProc func(framecount uint32, psamples []byte)

// SendProc type.
type SendProc func(framecount uint32, psamples []byte) uint32

// StopProc type.
type StopProc func()

// DeviceCallbacks contains callbacks for one initialized device.
type DeviceCallbacks struct {
	// Recv is called for capturing devices.
	Recv RecvProc
	// Send is called for playback devices.
	Send SendProc
	// Stop is called when the device stopped.
	Stop StopProc
}

// Device represents a streaming instance.
type Device uintptr

// InitDevice initializes a device.
//
// The device ID can be nil, in which case the default device is used. Otherwise, you
// can retrieve the ID by calling Context.Devices() and use the ID from the returned data.
//
// Set device ID to nil to use the default device. Do _not_ rely on the first device ID returned
// by Context.Devices() to be the default device.
//
// The returned instance has to be cleaned up using Uninit().
func InitDevice(context Context, deviceType DeviceType, deviceID *DeviceID,
	deviceConfig DeviceConfig, deviceCallbacks DeviceCallbacks) (*Device, error) {
	dev := Device(C.mal_aligned_malloc(C.size_t(unsafe.Sizeof(C.mal_device{})), simdAlignment))
	if dev == 0 {
		return nil, ErrOutOfMemory
	}

	rawDevice := dev.cptr()
	C.goSetDeviceConfigCallbacks(deviceConfig.cptr())
	result := C.mal_device_init(context.cptr(), C.mal_device_type(deviceType), deviceID.cptr(),
		deviceConfig.cptr(), nil, rawDevice)
	if result != 0 {
		dev.free()
		return nil, errorFromResult(Result(result))
	}
	deviceMutex.Lock()
	recvCallbacks[rawDevice] = deviceCallbacks.Recv
	sendCallbacks[rawDevice] = deviceCallbacks.Send
	stopCallbacks[rawDevice] = deviceCallbacks.Stop
	deviceMutex.Unlock()

	return &dev, nil
}

func (dev Device) cptr() *C.mal_device {
	return (*C.mal_device)(unsafe.Pointer(dev))
}

func (dev Device) free() {
	C.mal_aligned_free(unsafe.Pointer(dev))
}

// Type returns device type.
func (dev *Device) Type() DeviceType {
	return DeviceType(dev.cptr()._type)
}

// Format returns device format.
func (dev *Device) Format() FormatType {
	return FormatType(dev.cptr().format)
}

// Channels returns number of channels.
func (dev *Device) Channels() uint32 {
	return uint32(dev.cptr().channels)
}

// SampleRate returns sample rate.
func (dev *Device) SampleRate() uint32 {
	return uint32(dev.cptr().sampleRate)
}

// Start activates the device.
// For playback devices this begins playback. For capture devices it begins recording.
//
// For a playback device, this will retrieve an initial chunk of audio data from the client before
// returning. The reason for this is to ensure there is valid audio data in the buffer, which needs
// to be done _before_ the device begins playback.
//
// This API waits until the backend device has been started for real by the worker thread. It also
// waits on a mutex for thread-safety.
func (dev *Device) Start() error {
	result := C.mal_device_start(dev.cptr())
	return errorFromResult(Result(result))
}

// IsStarted determines whether or not the device is started.
func (dev *Device) IsStarted() bool {
	result := C.mal_device_is_started(dev.cptr())
	return result != 0
}

// Stop puts the device to sleep, but does not uninitialize it. Use Start() to start it up again.
//
// This API needs to wait on the worker thread to stop the backend device properly before returning. It
// also waits on a mutex for thread-safety. In addition, some backends need to wait for the device to
// finish playback/recording of the current fragment which can take some time (usually proportionate to
// the buffer size that was specified at initialization time).
func (dev *Device) Stop() error {
	result := C.mal_device_stop(dev.cptr())
	return errorFromResult(Result(result))
}

// Uninit uninitializes a device.
//
// This will explicitly stop the device. You do not need to call Stop() beforehand, but it's
// harmless if you do.
func (dev *Device) Uninit() {
	rawDevice := dev.cptr()
	deviceMutex.Lock()
	delete(recvCallbacks, rawDevice)
	delete(sendCallbacks, rawDevice)
	delete(stopCallbacks, rawDevice)
	deviceMutex.Unlock()

	C.mal_device_uninit(rawDevice)
	dev.free()
}

var deviceMutex sync.Mutex
var recvCallbacks = make(map[*C.mal_device]RecvProc)
var sendCallbacks = make(map[*C.mal_device]SendProc)
var stopCallbacks = make(map[*C.mal_device]StopProc)

//export goRecvCallback
func goRecvCallback(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) {
	deviceMutex.Lock()
	callback := recvCallbacks[pDevice]
	deviceMutex.Unlock()

	if callback != nil {
		samples := extractSlice(pDevice, frameCount, pSamples)
		callback(uint32(frameCount), samples)
	}
}

//export goSendCallback
func goSendCallback(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) (r C.mal_uint32) {
	deviceMutex.Lock()
	callback := sendCallbacks[pDevice]
	deviceMutex.Unlock()

	if callback != nil {
		samples := extractSlice(pDevice, frameCount, pSamples)
		r = C.mal_uint32(callback(uint32(frameCount), samples))
	}
	return r
}

//export goStopCallback
func goStopCallback(pDevice *C.mal_device) {
	deviceMutex.Lock()
	callback := stopCallbacks[pDevice]
	deviceMutex.Unlock()

	if callback != nil {
		callback()
	}
}

func extractSlice(pDevice *C.mal_device, frameCount C.mal_uint32, pSamples unsafe.Pointer) []byte {
	sampleCount := uint32(frameCount) * uint32(pDevice.channels)
	sizeInBytes := uint32(C.mal_get_bytes_per_sample(pDevice.format))
	psamples := (*[1 << 30]byte)(pSamples)[0 : sampleCount*sizeInBytes]
	return psamples
}
