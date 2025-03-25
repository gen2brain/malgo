package malgo

// #include "malgo.h"
import "C"
import (
	"unsafe"
)

type ConverterConfig struct {
	FormatIn       FormatType
	FormatOut      FormatType
	ChannelsIn     int
	ChannelsOut    int
	SampleRateIn   int
	SampleRateOut  int
	DitherMode     DitherModeType
	ChannelMixMode ChannelMixModeType
	Resampling     ResampleConfig

	// Unexposed: pChannelMapIn, pChannelMapOut, calculateLFEFromSpatialChannels, ppChannelWeights, allowDynamicSampleRate
}

type Converter struct {
	ptr *unsafe.Pointer
}

// InitConverter initializes a converter.
//
// Converter can be used to wrap sample format conversion, channel conversion and
// resampling into one operation. This is what miniaudio uses internally to convert between the format
// requested when the device was initialized and the format of the backend's native device.
//
// It is very similar to the resampling API.
//
// The returned instance has to be cleaned up using Uninit().
func InitConverter(config ConverterConfig) (*Converter, error) {
	ptr := C.ma_malloc(C.sizeof_ma_data_converter, nil)
	converter := Converter{
		ptr: &ptr,
	}
	if uintptr(*converter.ptr) == 0 {
		return nil, ErrOutOfMemory
	}

	configC := C.ma_data_converter_config_init_default()
	configC.formatIn = C.ma_format(config.FormatIn)
	configC.formatOut = C.ma_format(config.FormatOut)
	configC.channelsIn = C.ma_uint32(config.ChannelsIn)
	configC.channelsOut = C.ma_uint32(config.ChannelsOut)
	configC.sampleRateIn = C.ma_uint32(config.SampleRateIn)
	configC.sampleRateOut = C.ma_uint32(config.SampleRateOut)
	configC.resampling.algorithm = C.ma_resample_algorithm(config.Resampling.Algorithm)
	configC.resampling.linear.lpfOrder = C.uint(config.Resampling.Linear.LpfOrder)

	result := C.ma_data_converter_init(&configC, nil, converter.cptr())
	if result != 0 {
		C.ma_free(ptr, nil)
		return nil, errorFromResult(result)
	}

	return &converter, nil
}

// Uninit cleans up the ma_data_converter object.
func (c *Converter) Uninit() {
	C.ma_data_converter_uninit(c.cptr(), nil)
	c.free()
}

func (c Converter) free() {
	if c.ptr != nil {
		C.ma_free(*c.ptr, nil)
	}
}

func (c Converter) cptr() *C.ma_data_converter {
	return (*C.ma_data_converter)(*c.ptr)
}

// RequiredInputFrameCount returns how many input frames you need to provide in order to output a specific number of output frames.
func (c *Converter) RequiredInputFrameCount(outputFrameCount int) (int, error) {
	var cInputFrameCount C.ma_uint64
	var cOutputFrameCount C.ma_uint64 = C.ma_uint64(outputFrameCount)

	result := C.ma_data_converter_get_required_input_frame_count(c.cptr(), cOutputFrameCount, &cInputFrameCount)
	if result != 0 {
		return 0, errorFromResult(result)
	}

	return int(cInputFrameCount), nil
}

// ExpectOutputFrameCount returns how many output frames you can expect to get from a specific number of input frames.
func (c *Converter) ExpectOutputFrameCount(inputFrameCount int) (int, error) {
	var cInputFrameCount C.ma_uint64 = C.ma_uint64(inputFrameCount)
	var cOutputFrameCount C.ma_uint64

	result := C.ma_data_converter_get_expected_output_frame_count(c.cptr(), cInputFrameCount, &cOutputFrameCount)
	if result != 0 {
		return 0, errorFromResult(result)
	}

	return int(cOutputFrameCount), nil
}

// ProcessFrames processes PCM frames using the data converter.
//
// Processing always happens on a per PCM frame basis and always assumes interleaved input and output.
// De-interleaved processing is not supported. On input, this function takes the number of output frames
// you can fit in the output buffer and the number of input frames contained in the input buffer. On
// output these variables contain the number of output frames that were written to the output buffer
// and the number of input frames that were consumed in the process.
//
// You can pass in nil for the input buffer in which case it will be treated as an infinitely large
// buffer of zeros. The output buffer can also be nil, in which case the processing will be treated
// as seek.
func (c *Converter) ProcessFrames(pFramesIn []byte, frameCountIn int, pFramesOut []byte, frameCountOut int) (int, int, error) {
	var cFramesIn unsafe.Pointer
	if len(pFramesIn) == 0 || pFramesIn == nil {
		cFramesIn = unsafe.Pointer(nil)
	} else {
		cFramesIn = unsafe.Pointer(&pFramesIn[0])
	}

	var cFramesOut unsafe.Pointer
	if len(pFramesOut) == 0 || pFramesOut == nil {
		cFramesOut = unsafe.Pointer(nil)
	} else {
		cFramesOut = unsafe.Pointer(&pFramesOut[0])
	}

	var cFrameCountIn C.ma_uint64 = C.ma_uint64(frameCountIn)
	var cFrameCountOut C.ma_uint64 = C.ma_uint64(frameCountOut)

	result := C.ma_data_converter_process_pcm_frames(c.cptr(), cFramesIn, &cFrameCountIn, cFramesOut, &cFrameCountOut)
	if result != 0 {
		return 0, 0, errorFromResult(result)
	}

	return int(cFrameCountIn), int(cFrameCountOut), nil
}
