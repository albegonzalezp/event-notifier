package main

import (
	"event-notifier/email"
	"event-notifier/events"
)

type App struct {
	Emailer       *email.Emailer
	EventsManager *events.EventManager
}

func NewApp(emailer *email.Emailer, eventsManager *events.EventManager) *App {
	return &App{
		Emailer:       emailer,
		EventsManager: eventsManager,
	}
}
