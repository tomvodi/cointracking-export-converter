package interfaces

//go:generate mockgen -source app_config.go -destination ./mocks/mock_app_config.go

type AppConfig interface {
	SetTimezone(tz string) error
	Timezone() string
}
