package internal

type ServiceDataLoad interface {
	Load(path string) error
}
