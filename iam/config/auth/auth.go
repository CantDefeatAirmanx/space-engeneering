package config_auth

type AuthConfigData struct {
	SessionTTLHours int `env:"sessionTTLHours,required"`
}
