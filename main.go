package main

import (
	"fmt"

	"github.com/asheeshrana/event/event"
)

func main() {
	eventService := event.GetInstance()
	startEvent := eventService.CreateEvent("StartEvent")
	endEvent := eventService.CreateEvent("EndEvent")
	listenerInfo := event.ListenerInfo{
		Name:     "GenericListener",
		Callback: eventCallBack,
		EventList: map[string]bool{
			"StartEvent": true,
			"EndEvent":   true,
		},
	}
	eventService.RegisterListener(listenerInfo)
	eventService.TriggerEventSync(startEvent)
	fmt.Println("I am in the middle doing processing...")
	eventService.TriggerEventSync(endEvent)

}

func eventCallBack(e event.Event) {
	fmt.Println("Received Event " + e.GetName())
}
