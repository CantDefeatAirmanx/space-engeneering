package config_payment_client

type PaymentClientConfigData struct {
	Url string `env:"url,required"`
}
