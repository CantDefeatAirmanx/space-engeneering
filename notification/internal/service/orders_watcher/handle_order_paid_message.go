package service_orders_watcher

import (
	"bytes"
	"context"
	"embed"
	"text/template"

	kafka_events_order "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/order/v1"
)

//go:embed templates/order_paid_alert.tmpl
var templateFs embed.FS

var orderPaidTemplate = template.Must(template.ParseFS(
	templateFs, "templates/order_paid_alert.tmpl",
))

func (o *OrdersWatcherServiceImpl) handleOrderPaidMessage(
	ctx context.Context,
	message kafka_events_order.OrderPaidEvent,
) error {
	tmplData := orderPaidTmplData{
		OrderUUID:     message.OrderUUID,
		UserUUID:      message.UserUUID,
		PaymentMethod: string(message.PaymentMethod),
	}

	bytes := bytes.Buffer{}
	if err := orderPaidTemplate.Execute(&bytes, tmplData); err != nil {
		return err
	}

	err := o.notificationSender.SendNotification(
		ctx,
		bytes.String(),
	)
	return err
}
