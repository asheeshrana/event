package event

//ListenerInfo provides the name and callback function that will be invoked
// type ListenerInfo struct {
// 	Name      string
// 	Callback  callbackFunc
// 	EventList map[string]bool
// }

//ListenerInfo provides the name and callback function that will be invoked
type ListenerInfo interface {
	GetName() string
	Callback(event Event)
	GetEventNameMap() map[string]bool
}

//callbackFunc defines the function signature that is invoked when the event is triggered
type callbackFunc func(event Event)

//Event defines the interface of the event
type Event interface {
	GetName() string
	GetSourceName() string
	SetSourceName(sourceName string) Event
	GetData() interface{}
	SetData(interface{}) Event
}

//DefaultEvent defines the default event implementation available
type defaultEvent struct {
	eventName  *string
	sourceName *string
	data       interface{}
}

func (e defaultEvent) GetName() string {
	return *(e.eventName)
}

func (e defaultEvent) GetSourceName() string {
	return *(e.sourceName)
}

func (e defaultEvent) SetSourceName(sourceName string) Event {
	e.sourceName = &sourceName
	return e
}

func (e defaultEvent) GetData() interface{} {
	return e.data
}

func (e defaultEvent) SetData(data interface{}) Event {
	e.data = data
	return e
}

// DefaultListenerInfo provides the name and callback function that will be invoked
type DefaultListenerInfo struct {
	name      string
	callback  callbackFunc
	eventList map[string]bool
}

func (dli *DefaultListenerInfo) GetName() string {
	return dli.name
}

func (dli *DefaultListenerInfo) SetName(name string) *DefaultListenerInfo {
	dli.name = name
	return dli
}

func (dli *DefaultListenerInfo) Callback(event Event) {
	dli.callback(event)
}

func (dli *DefaultListenerInfo) SetCallback(callback callbackFunc) *DefaultListenerInfo {
	dli.callback = callback
	return dli
}

func (dli *DefaultListenerInfo) GetEventNameMap() map[string]bool {
	return dli.eventList
}

func (dli *DefaultListenerInfo) SetEventNameMap(eventNameMap map[string]bool) *DefaultListenerInfo {
	dli.eventList = eventNameMap
	return dli
}
