

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QTimerEvent>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class SplashCtxObj6592ff: public QObject
{
Q_OBJECT
Q_PROPERTY(QString status READ status WRITE setStatus NOTIFY statusChanged)
public:
	SplashCtxObj6592ff(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType();SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaTypes();callbackSplashCtxObj6592ff_Constructor(this);};
	QString status() { return ({ Moc_PackedString tempVal = callbackSplashCtxObj6592ff_Status(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setStatus(QString status) { QByteArray t48a366 = status.toUtf8(); Moc_PackedString statusPacked = { const_cast<char*>(t48a366.prepend("WHITESPACE").constData()+10), t48a366.size()-10 };callbackSplashCtxObj6592ff_SetStatus(this, statusPacked); };
	void Signal_StatusChanged(QString status) { QByteArray t48a366 = status.toUtf8(); Moc_PackedString statusPacked = { const_cast<char*>(t48a366.prepend("WHITESPACE").constData()+10), t48a366.size()-10 };callbackSplashCtxObj6592ff_StatusChanged(this, statusPacked); };
	 ~SplashCtxObj6592ff() { callbackSplashCtxObj6592ff_DestroySplashCtxObj(this); };
	void childEvent(QChildEvent * event) { callbackSplashCtxObj6592ff_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackSplashCtxObj6592ff_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackSplashCtxObj6592ff_CustomEvent(this, event); };
	void deleteLater() { callbackSplashCtxObj6592ff_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackSplashCtxObj6592ff_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackSplashCtxObj6592ff_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackSplashCtxObj6592ff_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackSplashCtxObj6592ff_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackSplashCtxObj6592ff_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackSplashCtxObj6592ff_TimerEvent(this, event); };
	QString statusDefault() { return _status; };
	void setStatusDefault(QString p) { if (p != _status) { _status = p; statusChanged(_status); } };
signals:
	void statusChanged(QString status);
public slots:
private:
	QString _status;
};

Q_DECLARE_METATYPE(SplashCtxObj6592ff*)


void SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

struct Moc_PackedString SplashCtxObj6592ff_Status(void* ptr)
{
	return ({ QByteArray t2a30d7 = static_cast<SplashCtxObj6592ff*>(ptr)->status().toUtf8(); Moc_PackedString { const_cast<char*>(t2a30d7.prepend("WHITESPACE").constData()+10), t2a30d7.size()-10 }; });
}

struct Moc_PackedString SplashCtxObj6592ff_StatusDefault(void* ptr)
{
	return ({ QByteArray t540991 = static_cast<SplashCtxObj6592ff*>(ptr)->statusDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t540991.prepend("WHITESPACE").constData()+10), t540991.size()-10 }; });
}

void SplashCtxObj6592ff_SetStatus(void* ptr, struct Moc_PackedString status)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->setStatus(QString::fromUtf8(status.data, status.len));
}

void SplashCtxObj6592ff_SetStatusDefault(void* ptr, struct Moc_PackedString status)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->setStatusDefault(QString::fromUtf8(status.data, status.len));
}

void SplashCtxObj6592ff_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<SplashCtxObj6592ff*>(ptr), static_cast<void (SplashCtxObj6592ff::*)(QString)>(&SplashCtxObj6592ff::statusChanged), static_cast<SplashCtxObj6592ff*>(ptr), static_cast<void (SplashCtxObj6592ff::*)(QString)>(&SplashCtxObj6592ff::Signal_StatusChanged));
}

void SplashCtxObj6592ff_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<SplashCtxObj6592ff*>(ptr), static_cast<void (SplashCtxObj6592ff::*)(QString)>(&SplashCtxObj6592ff::statusChanged), static_cast<SplashCtxObj6592ff*>(ptr), static_cast<void (SplashCtxObj6592ff::*)(QString)>(&SplashCtxObj6592ff::Signal_StatusChanged));
}

void SplashCtxObj6592ff_StatusChanged(void* ptr, struct Moc_PackedString status)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->statusChanged(QString::fromUtf8(status.data, status.len));
}

int SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType()
{
	return qRegisterMetaType<SplashCtxObj6592ff*>();
}

int SplashCtxObj6592ff_SplashCtxObj6592ff_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<SplashCtxObj6592ff*>(const_cast<const char*>(typeName));
}

int SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SplashCtxObj6592ff>();
#else
	return 0;
#endif
}

int SplashCtxObj6592ff_SplashCtxObj6592ff_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<SplashCtxObj6592ff>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* SplashCtxObj6592ff___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SplashCtxObj6592ff___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SplashCtxObj6592ff___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* SplashCtxObj6592ff___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void SplashCtxObj6592ff___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* SplashCtxObj6592ff___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* SplashCtxObj6592ff___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SplashCtxObj6592ff___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SplashCtxObj6592ff___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SplashCtxObj6592ff___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SplashCtxObj6592ff___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SplashCtxObj6592ff___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SplashCtxObj6592ff___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void SplashCtxObj6592ff___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* SplashCtxObj6592ff___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* SplashCtxObj6592ff_NewSplashCtxObj(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new SplashCtxObj6592ff(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new SplashCtxObj6592ff(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new SplashCtxObj6592ff(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new SplashCtxObj6592ff(static_cast<QWindow*>(parent));
	} else {
		return new SplashCtxObj6592ff(static_cast<QObject*>(parent));
	}
}

void SplashCtxObj6592ff_DestroySplashCtxObj(void* ptr)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->~SplashCtxObj6592ff();
}

void SplashCtxObj6592ff_DestroySplashCtxObjDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void SplashCtxObj6592ff_ChildEventDefault(void* ptr, void* event)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void SplashCtxObj6592ff_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void SplashCtxObj6592ff_CustomEventDefault(void* ptr, void* event)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void SplashCtxObj6592ff_DeleteLaterDefault(void* ptr)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->QObject::deleteLater();
}

void SplashCtxObj6592ff_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char SplashCtxObj6592ff_EventDefault(void* ptr, void* e)
{
	return static_cast<SplashCtxObj6592ff*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char SplashCtxObj6592ff_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<SplashCtxObj6592ff*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void SplashCtxObj6592ff_TimerEventDefault(void* ptr, void* event)
{
	static_cast<SplashCtxObj6592ff*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
