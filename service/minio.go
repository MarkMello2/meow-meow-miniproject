package service

type MinioService interface {
	getUrlImagePath(string) (string, error)
}
