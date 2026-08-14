package config

var defaultConfig = map[string]any{
	"auth.refresh_subject": RefreshTokenSubject,
	"auth.access_subject":  AccessTokenSubject,
}
