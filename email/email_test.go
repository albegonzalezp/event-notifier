package email

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
)

func TestSendEmail(t *testing.T) {
	pathToEnv := "../.env" // change your directory accordingly

	if err := godotenv.Load(pathToEnv); err != nil {
		t.Fatal(err)
	}

	smtpConfig, err := LoadDefaultConfig()
	if err != nil {
		t.Fatal(err)
	}

	emailer := NewEmailer(NewDialer(smtpConfig), smtpConfig)

	if err := emailer.Send([]string{gofakeit.Email()}, gofakeit.BuzzWord(), gofakeit.Phrase()); err != nil {
		t.Fatal(err)
	}
}
