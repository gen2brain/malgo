package io_api

import (
	"context"

	"github.com/gen2brain/malgo/mini_al"
)

func stream(ctx context.Context, abortChan chan error, deviceType mini_al.DeviceType,
	deviceConfig mini_al.DeviceConfig, deviceCallbacks mini_al.DeviceCallbacks) error {
	device, err := mini_al.InitDevice(mini_al.DefaultContext, deviceType, nil, deviceConfig, deviceCallbacks)
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
