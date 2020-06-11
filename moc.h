

#pragma once

#ifndef GO_MOC_fdc0e4_H
#define GO_MOC_fdc0e4_H

#include <stdint.h>

#ifdef __cplusplus
class qmlBridgefdc0e4;
void qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; void* ptr; };
struct Moc_PackedList { void* data; long long len; };
void qmlBridgefdc0e4_ConnectSendToQml(void* ptr, long long t);
void qmlBridgefdc0e4_DisconnectSendToQml(void* ptr);
void qmlBridgefdc0e4_SendToQml(void* ptr, uintptr_t fn);
int qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType();
int qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType2(char* typeName);
int qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType();
int qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* qmlBridgefdc0e4___children_atList(void* ptr, int i);
void qmlBridgefdc0e4___children_setList(void* ptr, void* i);
void* qmlBridgefdc0e4___children_newList(void* ptr);
void* qmlBridgefdc0e4___dynamicPropertyNames_atList(void* ptr, int i);
void qmlBridgefdc0e4___dynamicPropertyNames_setList(void* ptr, void* i);
void* qmlBridgefdc0e4___dynamicPropertyNames_newList(void* ptr);
void* qmlBridgefdc0e4___findChildren_atList(void* ptr, int i);
void qmlBridgefdc0e4___findChildren_setList(void* ptr, void* i);
void* qmlBridgefdc0e4___findChildren_newList(void* ptr);
void* qmlBridgefdc0e4___findChildren_atList3(void* ptr, int i);
void qmlBridgefdc0e4___findChildren_setList3(void* ptr, void* i);
void* qmlBridgefdc0e4___findChildren_newList3(void* ptr);
void* qmlBridgefdc0e4_NewQmlBridge(void* parent);
void qmlBridgefdc0e4_DestroyQmlBridge(void* ptr);
void qmlBridgefdc0e4_DestroyQmlBridgeDefault(void* ptr);
void qmlBridgefdc0e4_ChildEventDefault(void* ptr, void* event);
void qmlBridgefdc0e4_ConnectNotifyDefault(void* ptr, void* sign);
void qmlBridgefdc0e4_CustomEventDefault(void* ptr, void* event);
void qmlBridgefdc0e4_DeleteLaterDefault(void* ptr);
void qmlBridgefdc0e4_DisconnectNotifyDefault(void* ptr, void* sign);
char qmlBridgefdc0e4_EventDefault(void* ptr, void* e);
char qmlBridgefdc0e4_EventFilterDefault(void* ptr, void* watched, void* event);
;
void qmlBridgefdc0e4_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif