package events

import (
	"errors"
	"sync"
)

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
	if ed.Has(eventName, handler) {
		return ErrHandlerAlreadyRegistered
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	wg := &sync.WaitGroup{}
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandleInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandleInterface)
}
