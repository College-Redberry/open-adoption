package service

type Storage interface {
	UploadFile(objectName string, data []byte) error
	DownloadFile(objectName string) ([]byte, error)
}
