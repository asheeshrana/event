package event

import "fmt"

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
	eventName  string
	sourceName string
	data       interface{}
}

func (e *defaultEvent) GetName() string {
	return e.eventName
}

func (e *defaultEvent) GetSourceName() string {
	return e.sourceName
}

func (e *defaultEvent) SetSourceName(sourceName string) Event {
	e.sourceName = sourceName
	return e
}

func (e *defaultEvent) GetData() interface{} {
	return e.data
}

func (e *defaultEvent) SetData(data interface{}) Event {
	e.data = data
	return e
}

// DefaultListenerInfo provides the name and callback function that will be invoked
type DefaultListenerInfo struct {
	name      string
	callback  callbackFunc
	eventList map[string]bool
}

//GetName returns the name of the listener. It should be a unique identifier
func (dli *DefaultListenerInfo) GetName() string {
	return dli.name
}

//SetName sets the name of the listener. It should be a unique identifier
func (dli *DefaultListenerInfo) SetName(name string) *DefaultListenerInfo {
	dli.name = name
	return dli
}

//Callback is invoked whenever an event that listener subscribed to is triggered
func (dli *DefaultListenerInfo) Callback(event Event) {
	dli.callback(event)
}

//SetCallback is used to set the callback funnction
func (dli *DefaultListenerInfo) SetCallback(callback callbackFunc) *DefaultListenerInfo {
	dli.callback = callback
	return dli
}

//GetEventNameMap returns the map of events to which the listener is subscribed to
func (dli *DefaultListenerInfo) GetEventNameMap() map[string]bool {
	return dli.eventList
}

//SetEventNameMap sets the map with event names as keys. Callbacks will be invoked on the listeners only for these events
func (dli *DefaultListenerInfo) SetEventNameMap(eventNameMap map[string]bool) *DefaultListenerInfo {
	dli.eventList = eventNameMap
	return dli
}

//DefaultCallback is the default implementation of the callback function
func (dli *DefaultListenerInfo) DefaultCallback(event Event) {
	fmt.Println("This is default callback. EventName: " + event.GetName() + " SourceName: " + event.GetSourceName())
}
