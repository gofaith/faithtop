

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class qmlBridgefdc0e4: public QObject
{
Q_OBJECT
public:
	qmlBridgefdc0e4(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType();qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaTypes();callbackqmlBridgefdc0e4_Constructor(this);};
	void Signal_SendToQml(quintptr fn) { callbackqmlBridgefdc0e4_SendToQml(this, fn); };
	 ~qmlBridgefdc0e4() { callbackqmlBridgefdc0e4_DestroyQmlBridge(this); };
	void childEvent(QChildEvent * event) { callbackqmlBridgefdc0e4_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackqmlBridgefdc0e4_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackqmlBridgefdc0e4_CustomEvent(this, event); };
	void deleteLater() { callbackqmlBridgefdc0e4_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackqmlBridgefdc0e4_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackqmlBridgefdc0e4_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackqmlBridgefdc0e4_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackqmlBridgefdc0e4_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackqmlBridgefdc0e4_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackqmlBridgefdc0e4_TimerEvent(this, event); };
signals:
	void sendToQml(quintptr fn);
public slots:
private:
};

Q_DECLARE_METATYPE(qmlBridgefdc0e4*)


void qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaTypes() {
}

void qmlBridgefdc0e4_ConnectSendToQml(void* ptr, long long t)
{
	QObject::connect(static_cast<qmlBridgefdc0e4*>(ptr), static_cast<void (qmlBridgefdc0e4::*)(quintptr)>(&qmlBridgefdc0e4::sendToQml), static_cast<qmlBridgefdc0e4*>(ptr), static_cast<void (qmlBridgefdc0e4::*)(quintptr)>(&qmlBridgefdc0e4::Signal_SendToQml), static_cast<Qt::ConnectionType>(t));
}

void qmlBridgefdc0e4_DisconnectSendToQml(void* ptr)
{
	QObject::disconnect(static_cast<qmlBridgefdc0e4*>(ptr), static_cast<void (qmlBridgefdc0e4::*)(quintptr)>(&qmlBridgefdc0e4::sendToQml), static_cast<qmlBridgefdc0e4*>(ptr), static_cast<void (qmlBridgefdc0e4::*)(quintptr)>(&qmlBridgefdc0e4::Signal_SendToQml));
}

void qmlBridgefdc0e4_SendToQml(void* ptr, uintptr_t fn)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->sendToQml(fn);
}

int qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType()
{
	return qRegisterMetaType<qmlBridgefdc0e4*>();
}

int qmlBridgefdc0e4_qmlBridgefdc0e4_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<qmlBridgefdc0e4*>(const_cast<const char*>(typeName));
}

int qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<qmlBridgefdc0e4>();
#else
	return 0;
#endif
}

int qmlBridgefdc0e4_qmlBridgefdc0e4_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<qmlBridgefdc0e4>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* qmlBridgefdc0e4___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void qmlBridgefdc0e4___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* qmlBridgefdc0e4___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* qmlBridgefdc0e4___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void qmlBridgefdc0e4___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* qmlBridgefdc0e4___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* qmlBridgefdc0e4___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void qmlBridgefdc0e4___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* qmlBridgefdc0e4___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* qmlBridgefdc0e4___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void qmlBridgefdc0e4___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* qmlBridgefdc0e4___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* qmlBridgefdc0e4_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new qmlBridgefdc0e4(static_cast<QWindow*>(parent));
	} else {
		return new qmlBridgefdc0e4(static_cast<QObject*>(parent));
	}
}

void qmlBridgefdc0e4_DestroyQmlBridge(void* ptr)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->~qmlBridgefdc0e4();
}

void qmlBridgefdc0e4_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void qmlBridgefdc0e4_ChildEventDefault(void* ptr, void* event)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void qmlBridgefdc0e4_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void qmlBridgefdc0e4_CustomEventDefault(void* ptr, void* event)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void qmlBridgefdc0e4_DeleteLaterDefault(void* ptr)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->QObject::deleteLater();
}

void qmlBridgefdc0e4_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char qmlBridgefdc0e4_EventDefault(void* ptr, void* e)
{
	return static_cast<qmlBridgefdc0e4*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char qmlBridgefdc0e4_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<qmlBridgefdc0e4*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void qmlBridgefdc0e4_TimerEventDefault(void* ptr, void* event)
{
	static_cast<qmlBridgefdc0e4*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
