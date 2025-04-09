package events

import "event-notifier/events/bernabeu"

type EventManager struct {
	FootballManager *bernabeu.FootballManager
	// Add more as you please
}

func NewEventManager(football *bernabeu.FootballManager) *EventManager {
	return &EventManager{FootballManager: football}
}
