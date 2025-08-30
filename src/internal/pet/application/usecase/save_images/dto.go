package request

import "mime/multipart"

type Input struct {
	ID     string
	Images []*multipart.FileHeader
}
