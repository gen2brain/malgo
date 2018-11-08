package io_api

import (
	"context"

	"github.com/gen2brain/malgo"
)

func stream(ctx context.Context, abortChan chan error, deviceType malgo.DeviceType,
	deviceConfig malgo.DeviceConfig, deviceCallbacks malgo.DeviceCallbacks) error {
	device, err := malgo.InitDevice(malgo.DefaultContext, deviceType, nil, deviceConfig, deviceCallbacks)
	if err != nil {
		return err
	}
	defer device.Uninit()

	err = device.Start()
	if err != nil {
		return err
	}

	ctxChan := ctx.Done()
	if ctxChan != nil {
		select {
		case <-ctxChan:
			err = ctx.Err()
		case err = <-abortChan:
		}
	} else {
		err = <-abortChan
	}

	return err
}
