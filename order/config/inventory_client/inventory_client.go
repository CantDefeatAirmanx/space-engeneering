package config_inventory_client

type InventoryClientConfigData struct {
	Url string `env:"url,required"`
}
