#ifndef H_MALGO
#define H_MALGO

#include "miniaudio.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void goLogCallback(ma_context* pContext, char* message);
void goSetContextConfigCallbacks(ma_context_config* pConfig, ma_context* pContext);

extern void goDataCallback(ma_device *pDevice, void *pOutput, void *pInput, ma_uint32 frameCount);
extern void goStopCallback(ma_device* pDevice);
void goSetDeviceConfigCallbacks(ma_device_config* pConfig);

#ifdef __cplusplus
}
#endif

#endif
