package constants

import "os"

var (
	DB_URL string = os.Getenv("DB_URL")
)
