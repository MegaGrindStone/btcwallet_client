

#pragma once

#ifndef GO_MOC_6592ff_H
#define GO_MOC_6592ff_H

#include <stdint.h>

#ifdef __cplusplus
class SplashCtxObj6592ff;
void SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
struct Moc_PackedString SplashCtxObj6592ff_Status(void* ptr);
struct Moc_PackedString SplashCtxObj6592ff_StatusDefault(void* ptr);
void SplashCtxObj6592ff_SetStatus(void* ptr, struct Moc_PackedString status);
void SplashCtxObj6592ff_SetStatusDefault(void* ptr, struct Moc_PackedString status);
void SplashCtxObj6592ff_ConnectStatusChanged(void* ptr);
void SplashCtxObj6592ff_DisconnectStatusChanged(void* ptr);
void SplashCtxObj6592ff_StatusChanged(void* ptr, struct Moc_PackedString status);
int SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType();
int SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType2(char* typeName);
int SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType();
int SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* SplashCtxObj6592ff___children_atList(void* ptr, int i);
void SplashCtxObj6592ff___children_setList(void* ptr, void* i);
void* SplashCtxObj6592ff___children_newList(void* ptr);
void* SplashCtxObj6592ff___dynamicPropertyNames_atList(void* ptr, int i);
void SplashCtxObj6592ff___dynamicPropertyNames_setList(void* ptr, void* i);
void* SplashCtxObj6592ff___dynamicPropertyNames_newList(void* ptr);
void* SplashCtxObj6592ff___findChildren_atList(void* ptr, int i);
void SplashCtxObj6592ff___findChildren_setList(void* ptr, void* i);
void* SplashCtxObj6592ff___findChildren_newList(void* ptr);
void* SplashCtxObj6592ff___findChildren_atList3(void* ptr, int i);
void SplashCtxObj6592ff___findChildren_setList3(void* ptr, void* i);
void* SplashCtxObj6592ff___findChildren_newList3(void* ptr);
void* SplashCtxObj6592ff___qFindChildren_atList2(void* ptr, int i);
void SplashCtxObj6592ff___qFindChildren_setList2(void* ptr, void* i);
void* SplashCtxObj6592ff___qFindChildren_newList2(void* ptr);
void* SplashCtxObj6592ff_NewSplashCtxObj(void* parent);
void SplashCtxObj6592ff_DestroySplashCtxObj(void* ptr);
void SplashCtxObj6592ff_DestroySplashCtxObjDefault(void* ptr);
void SplashCtxObj6592ff_ChildEventDefault(void* ptr, void* event);
void SplashCtxObj6592ff_ConnectNotifyDefault(void* ptr, void* sign);
void SplashCtxObj6592ff_CustomEventDefault(void* ptr, void* event);
void SplashCtxObj6592ff_DeleteLaterDefault(void* ptr);
void SplashCtxObj6592ff_DisconnectNotifyDefault(void* ptr, void* sign);
char SplashCtxObj6592ff_EventDefault(void* ptr, void* e);
char SplashCtxObj6592ff_EventFilterDefault(void* ptr, void* watched, void* event);
;
void SplashCtxObj6592ff_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif