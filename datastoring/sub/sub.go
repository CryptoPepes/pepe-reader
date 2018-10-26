package sub

import (
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"log"
	"cryptopepe.io/cryptopepe-reader/datastoring/sub/event"
)

func HandleAll(hub *triggers.TriggerHub, context *event.EventContext) {

	// define all the handlers
	var presets = map[string]event.EventHandlerPreset{
		"pepe update": {Broadcast: hub.Pepe, HandleFn: PepeUpdate},
		"user update": {Broadcast: hub.User, HandleFn: UserUpdate},
	}

	// create an instance of all presets, and start them.
	for key, preset := range presets {
		handler := preset.Make()
		go func(name string, evHandler *event.EventHandler) {
			log.Printf("Handler '%s' started!\n", name)
			evHandler.Start(context)
		}(key, handler)
	}
}