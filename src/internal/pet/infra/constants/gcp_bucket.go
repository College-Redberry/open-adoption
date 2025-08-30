package constants

import "os"

var (
	GCP_BUCKET_NAME          string = os.Getenv("GCP_BUCKET_NAME")
	GCP_BUCKET_IMAGES_FOLDER string = os.Getenv("GCP_BUCKET_IMAGES_FOLDER")
)
