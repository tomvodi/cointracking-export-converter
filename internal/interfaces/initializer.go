package interfaces

//go:generate mockgen -source initializer.go -destination ./mocks/mock_initializer.go

type Initializer interface {
	Init() error
}
