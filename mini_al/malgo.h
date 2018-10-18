#ifndef H_MALGO
#define H_MALGO

#include "mini_al.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void goLogCallback(mal_context* pContext, mal_device* pDevice, char* message);
void goSetContextConfigCallbacks(mal_context_config* pConfig);

extern void goRecvCallback(mal_device* pDevice, mal_uint32 frameCount, void* pSamples);
extern mal_uint32 goSendCallback(mal_device* pDevice, mal_uint32 frameCount, void* pSamples);
extern void goStopCallback(mal_device* pDevice);
void goSetDeviceConfigCallbacks(mal_device_config* pConfig);

#ifdef __cplusplus
}
#endif

#endif
