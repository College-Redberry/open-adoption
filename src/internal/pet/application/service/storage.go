package service

type Storage interface {
	UploadFile(objectName string, data []byte) error
	GenerateUrl(objectName string) (string, error)
}
