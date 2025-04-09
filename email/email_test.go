package email

import (
	"event-notifier/events/bernabeu"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
)

func TestSendEmail(t *testing.T) {
	emailer, err := initEnvForTesting()
	if err != nil {
		t.Fatal(err)
	}

	if err := emailer.Send([]string{gofakeit.Email()}, gofakeit.BuzzWord(), gofakeit.Phrase()); err != nil {
		t.Fatal(err)
	}
}

func TestSendAlert(t *testing.T) {
	emailer, err := initEnvForTesting()
	if err != nil {
		t.Fatal(err)
	}

	match := bernabeu.Match{
		Description: struct {
			Plaintext string "json:\"plaintext\""
		}{
			"Real Madrid vs Osasuna",
		},
		DateTime: time.Now().UTC(),
		Competition: struct {
			OptaID        string "json:\"optaId\""
			OptaLegacyID  string "json:\"optaLegacyId\""
			CompetitionID string "json:\"competitionId\""
			Slug          string "json:\"slug\""
			Logo          struct {
				PublishURL string "json:\"_publishUrl\""
				DmS7URL    string "json:\"_dmS7Url\""
			} "json:\"logo\""
			Name string "json:\"name\""
		}{
			Name: "Champions League",
		},
	}

	if err := emailer.SendFootballMatchTodayAlert(match); err != nil {
		t.Fatal(err)
	}

}

func initEnvForTesting() (*Emailer, error) {
	pathToEnv := "../.env" // change your directory accordingly

	if err := godotenv.Load(pathToEnv); err != nil {
		return nil, err
	}

	smtpConfig, err := LoadDefaultConfig()
	if err != nil {
		return nil, err
	}

	emailer := NewEmailer(NewDialer(smtpConfig), smtpConfig)

	return emailer, nil
}
