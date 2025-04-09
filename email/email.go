package email

import (
	"crypto/tls"
	"event-notifier/events/bernabeu"
	"fmt"

	"gopkg.in/gomail.v2"
)

type Emailer struct {
	Dialer *gomail.Dialer
	Config Config
}

func NewEmailer(dialer *gomail.Dialer, config Config) *Emailer {
	return &Emailer{Dialer: dialer, Config: config}
}

func NewDialer(config Config) *gomail.Dialer {

	d := gomail.NewDialer(
		fmt.Sprintf("smtp.%s.com", config.Provider),
		config.Port,
		config.Username,
		config.Password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: config.TlsInsecureSkipVerify}

	return d
}

func (em *Emailer) Send(to []string, subject, body string) error {
	message := gomail.NewMessage()

	message.SetHeader("From", em.Config.Username)
	message.SetHeader("To", to...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	if err := em.Dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}

func (em *Emailer) SendFootballMatchTodayAlert(match bernabeu.Match) error {
	to := []string{"example1@gmail.com", "example2@gmail.com"}
	subject := "ATENCIÓN: Hoy hay partido en el Bernabéu"
	body := fmt.Sprintf("Este es una notificación para recordarte que hoy juega %s, a las %s, por %s", match.Description.Plaintext, match.DateTime, match.Competition.Name)

	if err := em.Send(to, subject, body); err != nil {
		return err
	}

	return nil
}
