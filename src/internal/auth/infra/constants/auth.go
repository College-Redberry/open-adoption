package constants

import (
	"os"
	"strconv"
)

var (
	JWT_SECRET_KEY          = os.Getenv("JWT_SECRET_KEY")
	ISSUER                  = os.Getenv("ISSUER")
	EXPIRATION_TIME         = must("EXPIRATION_TIME")
	REFRESH_EXPIRATION_TIME = must("REFRESH_EXPIRATION_TIME")
)

func must(envName string) int {
	result, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		panic("Environment variable " + envName + " is not set or is not a valid integer" + err.Error())
	}

	return result
}
