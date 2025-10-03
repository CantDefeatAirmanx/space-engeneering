package model_notification_method

type NotificationMethod struct {
	UUID         string
	ProviderName ProviderName
	Target       string
}

type NotificationMethodFilter struct {
	UserUUID string
}

type ProviderName string

const (
	ProviderNameTelegram ProviderName = "telegram"
)
