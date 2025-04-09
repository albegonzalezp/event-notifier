package main

import (
	"event-notifier/email"
	"event-notifier/events"
	"event-notifier/events/bernabeu"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if os.Getenv("ENV") != "prod" {
		log.Println("Loaded envs.")
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Load smtp config.")
	smtpConfig, err := email.LoadDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting app.")
	app := NewApp(
		email.NewEmailer(email.NewDialer(smtpConfig), smtpConfig), // Initialize Emailer
		events.NewEventManager(bernabeu.NewFootballManager()),     // Initiliaze Events Manager
	)

	log.Println("Scaning for today football games...")
	if err := app.EventsManager.FootballManager.GetMatchesToday(); err != nil {
		log.Fatal(err)
	}

	if len(app.EventsManager.FootballManager.MatchesToday) < 1 {
		log.Println("No matches found today.")
		os.Exit(0)
	}

	log.Println("Sending email(s) to notify about the game today.")
	if err := app.Emailer.SendFootballMatchTodayAlert(app.EventsManager.FootballManager.MatchesToday[0]); err != nil {
		log.Fatal(err)
	}

}
