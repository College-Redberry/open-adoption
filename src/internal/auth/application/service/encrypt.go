package service

type EncryptService interface {
	Decrypt(ciphertext []byte) ([]byte, error)
}
