package email

import (
	"os"
	"strconv"
)

type Config struct {
	Provider              string
	Port                  int
	Username              string
	Password              string
	TlsInsecureSkipVerify bool
}

func NewConfig(
	provider string,
	port int,
	username string,
	password string,
	tlsVerify bool,
) Config {
	return Config{
		Provider:              provider,
		Port:                  port,
		Username:              username,
		Password:              password,
		TlsInsecureSkipVerify: tlsVerify,
	}
}

func LoadDefaultConfig() (Config, error) {
	port, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		return Config{}, err
	}

	tlsVerify, err := strconv.ParseBool(os.Getenv("EMAIL_TLS_SKIP_VERIFY"))
	if err != nil {
		return Config{}, err
	}

	return Config{
		Provider:              os.Getenv("EMAIL_PROVIDER"),     // gmail, outlook, etc
		Port:                  port,                            // most cases 587
		Username:              os.Getenv("EMAIL_USERNAME"),     // your email: example@example.com
		Password:              os.Getenv("EMAIL_APP_PASSWORD"), // this is your EMAIL APPLICATION PASSWORD, different from the one you use to login. Check here how to set it up: https://support.google.com/mail/answer/185833
		TlsInsecureSkipVerify: tlsVerify,                       // for ease of use, in most cases true
	}, nil
}
