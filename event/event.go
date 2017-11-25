package event

//EventInterface defines the interface of the event
type EventInterface interface {
	GetEventId() int
	GetEventName() string	
	GetDescription() string
	GetSourceId() int
	GetSourceName() string
	GetData() interface {}
	SetData(interface{}) EventInterface
}

//DefaultEvent defines the default event implementation available
type defaultEvent struct {
	eventId *int
	eventName *string
	description *string
	sourceId *int
	sourceName *string
	data interface {}
}

func (e defaultEvent) GetEventId() int {
	return *(e.eventId)
}

func (e defaultEvent) GetEventName() string {
	return *(e.eventName)
}

func (e defaultEvent) GetDescription() string {
	return *(e.description)
}

func (e defaultEvent) GetSourceId() int {
	return *(e.sourceId)
}

func (e defaultEvent) GetSourceName() string {
	return *(e.sourceName)
}

func (e defaultEvent) GetData() interface{} {
	return e.data
}

func (e defaultEvent) SetData(data interface{}) EventInterface {
	e.data = data
	return e
}