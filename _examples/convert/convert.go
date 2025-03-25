package main

import (
	"github.com/gen2brain/malgo"
	"io"
)

type PCMFormat struct {
	Type       malgo.FormatType
	Channels   int
	SampleRate int
}

func convert(reader io.Reader, inputFormat PCMFormat, writer io.Writer, outputFormat PCMFormat) error {
	config := malgo.ConverterConfig{
		FormatIn:      inputFormat.Type,
		FormatOut:     outputFormat.Type,
		ChannelsIn:    inputFormat.Channels,
		ChannelsOut:   outputFormat.Channels,
		SampleRateIn:  inputFormat.SampleRate,
		SampleRateOut: outputFormat.SampleRate,
		Resampling: malgo.ResampleConfig{
			Algorithm: malgo.ResampleAlgorithmLinear,
		},
		DitherMode:     malgo.DitherModeTriangle,
		ChannelMixMode: malgo.ChannelMixModeSimple,
	}
	converter, err := malgo.InitConverter(config)
	if err != nil {
		return err
	}
	defer converter.Uninit()

	inFrameSize := malgo.FrameSizeInBytes(config.FormatIn, config.ChannelsIn)
	outFrameSize := malgo.FrameSizeInBytes(config.FormatOut, config.ChannelsOut)

	inputFrames := 1000
	expectFrames, _ := converter.ExpectOutputFrameCount(inputFrames)
	inBuffer := make([]byte, inFrameSize*inputFrames)
	outBuffer := make([]byte, outFrameSize*expectFrames)

	for {
		n, err := reader.Read(inBuffer)
		if err != nil {
			return err
		}

		readFrameCount := n / inFrameSize
		_, outFrameCount, err := converter.ProcessFrames(inBuffer, readFrameCount, outBuffer, expectFrames)
		if err != nil {
			return err
		}

		_, err = writer.Write(outBuffer[:outFrameCount*outFrameSize])
		if err != nil {
			return err
		}
	}
}
