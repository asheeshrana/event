package event

type manager struct {

}

type ManagerInterface interface {
	RegisterEventType() EventInterface
	GetEventInstanceByNane(eventId string) EventInterface
	GetEventInstanceById(eventId int) EventInterface
}