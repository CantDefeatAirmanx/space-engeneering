package config_auth_client

type AuthClientConfigData struct {
	Url string `env:"url,required"`
}
