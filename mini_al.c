#include "_cgo_export.h"

#define MAL_IMPLEMENTATION
#include "mini_al.h"

void goSetRecvCallback(mal_device* pDevice) {
	mal_device_set_recv_callback(pDevice, goRecvCallback);
}

void goSetSendCallback(mal_device* pDevice) {
	mal_device_set_send_callback(pDevice, goSendCallback);
}

void goSetStopCallback(mal_device* pDevice) {
	mal_device_set_stop_callback(pDevice, goStopCallback);
}

mal_device_config goConfigInit(mal_format format, mal_uint32 channels, mal_uint32 sampleRate) {
    return mal_device_config_init(format, channels, sampleRate, goRecvCallback, goSendCallback);
}

mal_device_config goConfigInitCapture(mal_format format, mal_uint32 channels, mal_uint32 sampleRate) {
    return mal_device_config_init(format, channels, sampleRate, goRecvCallback, NULL);
}

mal_device_config goConfigInitPlayback(mal_format format, mal_uint32 channels, mal_uint32 sampleRate) {
    return mal_device_config_init(format, channels, sampleRate, NULL, goSendCallback);
}
