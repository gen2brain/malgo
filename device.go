package malgo

// #include "malgo.h"
import "C"
import (
	"sync"
	"unsafe"
)

// DataProc type.
type DataProc func(pOutputSample, pInputSamples []byte, framecount uint32)

// StopProc type.
type StopProc func()

// DeviceCallbacks contains callbacks for one initialized device.
type DeviceCallbacks struct {
	// Data is called for the full duplex IO.
	Data DataProc
	// Stop is called when the device stopped.
	Stop StopProc
}

// Device represents a streaming instance.
type Device struct {
	ptr *unsafe.Pointer
}

// InitDevice initializes a device.
//
// The device ID can be nil, in which case the default device is used. Otherwise, you
// can retrieve the ID by calling Context.Devices() and use the ID from the returned data.
//
// Set device ID to nil to use the default device. Do _not_ rely on the first device ID returned
// by Context.Devices() to be the default device.
//
// The returned instance has to be cleaned up using Uninit().
func InitDevice(context Context, deviceConfig DeviceConfig, deviceCallbacks DeviceCallbacks) (*Device, error) {
	ptr := C.ma_malloc(C.sizeof_ma_device, nil)
	dev := Device{
		ptr: &ptr,
	}
	if uintptr(*dev.ptr) == 0 {
		return nil, ErrOutOfMemory
	}
	devConfigC, release := deviceConfig.toC()
	defer release()

	rawDevice := dev.cptr()
	C.goSetDeviceConfigCallbacks(&devConfigC)
	result := C.ma_device_init(context.cptr(), &devConfigC, rawDevice)
	if result != 0 {
		dev.free()
		return nil, errorFromResult(result)
	}
	deviceMutex.Lock()
	dataCallbacks[rawDevice] = deviceCallbacks.Data
	stopCallbacks[rawDevice] = deviceCallbacks.Stop
	deviceMutex.Unlock()

	return &dev, nil
}

func (dev Device) cptr() *C.ma_device {
	return (*C.ma_device)(*dev.ptr)
}

func (dev Device) free() {
	if dev.ptr != nil {
		C.ma_free(*dev.ptr, nil)
	}
}

// Type returns device type.
func (dev *Device) Type() DeviceType {
	return DeviceType(dev.cptr()._type)
}

// PlaybackFormat returns device playback format.
func (dev *Device) PlaybackFormat() FormatType {
	return FormatType(dev.cptr().playback.format)
}

// CaptureFormat returns device capture format.
func (dev *Device) CaptureFormat() FormatType {
	return FormatType(dev.cptr().capture.format)
}

// PlaybackChannels returns number of playback channels.
func (dev *Device) PlaybackChannels() uint32 {
	return uint32(dev.cptr().playback.channels)
}

// CaptureChannels returns number of playback channels.
func (dev *Device) CaptureChannels() uint32 {
	return uint32(dev.cptr().capture.channels)
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
	result := C.ma_device_start(dev.cptr())
	return errorFromResult(result)
}

// IsStarted determines whether or not the device is started.
func (dev *Device) IsStarted() bool {
	result := C.ma_device_is_started(dev.cptr())
	return result != 0
}

// Stop puts the device to sleep, but does not uninitialize it. Use Start() to start it up again.
//
// This API needs to wait on the worker thread to stop the backend device properly before returning. It
// also waits on a mutex for thread-safety. In addition, some backends need to wait for the device to
// finish playback/recording of the current fragment which can take some time (usually proportionate to
// the buffer size that was specified at initialization time).
func (dev *Device) Stop() error {
	result := C.ma_device_stop(dev.cptr())
	return errorFromResult(result)
}

// Uninit uninitializes a device.
//
// This will explicitly stop the device. You do not need to call Stop() beforehand, but it's
// harmless if you do.
func (dev *Device) Uninit() {
	rawDevice := dev.cptr()
	deviceMutex.Lock()
	delete(dataCallbacks, rawDevice)
	delete(stopCallbacks, rawDevice)
	deviceMutex.Unlock()

	C.ma_device_uninit(rawDevice)
	dev.free()
}

var deviceMutex sync.Mutex
var dataCallbacks = make(map[*C.ma_device]DataProc)
var stopCallbacks = make(map[*C.ma_device]StopProc)

//export goDataCallback
func goDataCallback(pDevice *C.ma_device, pOutput, pInput unsafe.Pointer, frameCount C.ma_uint32) {
	deviceMutex.Lock()
	callback := dataCallbacks[pDevice]
	deviceMutex.Unlock()

	if callback != nil {
		var inputSamples, outputSamples []byte

		if pOutput != nil {
			sampleCount := uint32(frameCount) * uint32(pDevice.playback.channels)
			sizeInBytes := uint32(C.ma_get_bytes_per_sample(pDevice.playback.format))
			outputSamples = unsafe.Slice((*byte)(pOutput), sampleCount*sizeInBytes)
		}

		if pInput != nil {
			sampleCount := uint32(frameCount) * uint32(pDevice.capture.channels)
			sizeInBytes := uint32(C.ma_get_bytes_per_sample(pDevice.capture.format))
			inputSamples = unsafe.Slice((*byte)(pInput), sampleCount*sizeInBytes)
		}

		callback(outputSamples, inputSamples, uint32(frameCount))
	}
}

//export goStopCallback
func goStopCallback(pDevice *C.ma_device) {
	deviceMutex.Lock()
	callback := stopCallbacks[pDevice]
	deviceMutex.Unlock()

	if callback != nil {
		callback()
	}
}
