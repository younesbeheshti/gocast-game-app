package config

import "time"

var defaultConfig = map[string]any{
	"auth.refresh_subject":                  RefreshTokenSubject,
	"auth.access_subject":                   AccessTokenSubject,
	"auth.refresh_duration_time":            RefreshTokenExpireDuration,
	"auth.access_duration_time":             AccessTokenExpireDuration,
	"application.graceful_shutdown_timeout": time.Second * 5,
}
