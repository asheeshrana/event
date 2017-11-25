package event

//EventInterface defines the interface of the event
type EventInterface interface {
	GetEventId() int
	GetDescription() string
	GetSourceId() int
	GetSourceName() string
	GetData() interface {}
	SetData(interface{})
}

//DefaultEvent defines the default event implementation available
type DefaultEvent struct {
	eventId int
	description string
	data interface {}
}

