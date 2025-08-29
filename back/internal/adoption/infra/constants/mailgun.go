package constants

import "os"

var (
	MAILGUN_DOMAIN  string = os.Getenv("MAILGUN_DOMAIN")
	MAILGUN_API_KEY string = os.Getenv("MAILGUN_API_KEY")
	MAILGUN_SENDER  string = os.Getenv("MAILGUN_SENDER")
	EMAIL_TO        string = os.Getenv("EMAIL_TO")
)
