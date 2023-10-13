package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("already handler registered")

type EventDispatcher struct {
	handlers map[string][]EventHandleInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandleInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandleInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
