package event

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
