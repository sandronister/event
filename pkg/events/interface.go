package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandleInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandleInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandleInterface) error
	Has(eventName string, handler EventHandleInterface) bool
	Clear() error
}
