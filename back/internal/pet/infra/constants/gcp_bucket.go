package constants

import "os"

var (
	GCP_BUCKET_NAME string = os.Getenv("GCP_BUCKET_NAME")
)
