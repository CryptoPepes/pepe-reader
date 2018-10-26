package event

import (
	"cryptopepe.io/cryptopepe-reader/reader"
	"cryptopepe.io/cryptopepe-reader/datastoring/triggers"
	"cryptopepe.io/cryptopepe-reader/datastoring/data"
	"cryptopepe.io/cryptopepe-reader/bio-gen"
)


type EventHandlerFn func(context *EventContext, trig *triggers.Trigger) error

type EventHandler struct {
	input triggers.TriggerListener
	handleFn EventHandlerFn
}

func (handler *EventHandler) Start(context *EventContext) {
	//range over input channel, will close when input closes.
	for ev := range handler.input {
		handler.handleFn(context, ev)
	}
}

type EventHandlerPreset struct {
	Broadcast *triggers.TriggerBroadcast
	HandleFn EventHandlerFn
}

func (preset *EventHandlerPreset) Make() *EventHandler {
	ch := make(chan *triggers.Trigger)
	preset.Broadcast.RegisterListener(ch)
	return &EventHandler{input: ch, handleFn: preset.HandleFn}
}

type EventContext struct {
	Reader reader.Reader
	EntityBuf *data.EntityBuffer
	BioGenerator *bio_gen.BioGenerator
}