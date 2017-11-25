package event

import (
	"fmt"
	"sync"
)

type defaultService struct {
	eventListenerMap map[string][]ListenerInfo
	listenerMap      map[string]ListenerInfo
	mutex            *sync.Mutex
}

var service Service

//Service defines the interface that will be exposed to the user
type Service interface {
	CreateEvent(eventName string) Event
	TriggerEventSync(event Event) bool
	TriggerEventAsync(event Event) *sync.WaitGroup
	RegisterListener(listenerInfo ListenerInfo) bool
	UnRegisterListener(listenerName string) bool
	UnRegisterListenerFromEvents(listenerName string, eventNames []string) bool
}

//GetInstance returns the intance of the event service
func GetInstance() Service {
	if service == nil {
		service = &defaultService{
			eventListenerMap: make(map[string][]ListenerInfo),
			listenerMap:      make(map[string]ListenerInfo),
			mutex:            &sync.Mutex{},
		}
	}
	return service
}

func (d defaultService) CreateEvent(eventName string) Event {
	return &defaultEvent{eventName: &eventName}
}

func (d defaultService) TriggerEventSync(event Event) bool {
	if listenerInfoList, ok := d.eventListenerMap[event.GetName()]; ok {
		for _, listenerInfo := range listenerInfoList {
			listenerInfo.Callback(event)
		}
		return true
	}
	return false
}

func (d defaultService) TriggerEventAsync(event Event) *sync.WaitGroup {
	var wg sync.WaitGroup
	if listenerInfoList, ok := d.eventListenerMap[event.GetName()]; ok {
		wg.Add(len(listenerInfoList))
		for _, listenerInfo := range listenerInfoList {
			go listenerInfo.Callback(event)
		}
	}
	return &wg
}

func (d defaultService) RegisterListener(listenerInfo ListenerInfo) bool {
	fmt.Println("Name of the event listener = " + listenerInfo.GetName())
	if registeredListenerInfo, ok := d.listenerMap[listenerInfo.GetName()]; ok {
		for eventName := range listenerInfo.GetEventNameMap() {
			if _, ok := registeredListenerInfo.GetEventNameMap()[eventName]; !ok {
				registeredListenerInfo.GetEventNameMap()[eventName] = true
				registeredListenerInfoList := d.eventListenerMap[eventName]
				d.eventListenerMap[eventName] = append(registeredListenerInfoList, registeredListenerInfo)
			}
		}
	} else {
		d.listenerMap[listenerInfo.GetName()] = listenerInfo
		for eventName := range listenerInfo.GetEventNameMap() {
			if registeredListenerInfoList, ok := d.eventListenerMap[eventName]; ok {
				d.eventListenerMap[eventName] = append(registeredListenerInfoList, listenerInfo)
			} else {
				d.eventListenerMap[eventName] = []ListenerInfo{listenerInfo}
			}
		}
	}
	return true
}

func (d defaultService) UnRegisterListener(listenerName string) bool {
	if listenerInfo, ok := d.listenerMap[listenerName]; ok {
		for eventName := range listenerInfo.GetEventNameMap() {
			registeredListenerInfoList := d.eventListenerMap[eventName]
			if len(registeredListenerInfoList) == 1 {
				delete(d.eventListenerMap, eventName)
			} else {
				for index, registeredListenerInfo := range registeredListenerInfoList {
					if listenerInfo == registeredListenerInfo {
						d.eventListenerMap[eventName] = append(registeredListenerInfoList[:index], registeredListenerInfoList[index+1:]...)
						break
					}
				}
			}
		}
		delete(d.listenerMap, listenerName)
	}
	return true
}

func (d defaultService) UnRegisterListenerFromEvents(listenerName string, eventNames []string) bool {
	if listenerInfo, ok := d.listenerMap[listenerName]; ok {
		for _, eventName := range eventNames {
			registeredListenerInfoList := d.eventListenerMap[eventName]
			if len(registeredListenerInfoList) == 1 {
				delete(d.eventListenerMap, eventName)
			} else {
				for index, registeredListenerInfo := range registeredListenerInfoList {
					if listenerInfo == registeredListenerInfo {
						d.eventListenerMap[eventName] = append(registeredListenerInfoList[:index], registeredListenerInfoList[index+1:]...)
						break
					}
				}
			}
			if len(listenerInfo.GetEventNameMap()) == 0 {
				//Delete the listenerInfo as there there is no event associated with the listener
				delete(d.listenerMap, eventName)
			}
		}
	}
	return true
}
