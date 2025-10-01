package service_assemblies_watcher

import (
	"bytes"
	"context"
	"embed"
	"html/template"

	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
)

//go:embed templates/assembly_completed_alert.tmpl
var templateFs embed.FS

var assemblyCompletedTemplate = template.Must(template.ParseFS(
	templateFs, "templates/assembly_completed_alert.tmpl",
))

func (a *AssembliesWatcherServiceImpl) handleAssemblyCompleted(
	ctx context.Context,
	message kafka_events_ship_assembly.ShipAssembledEvent,
) error {
	tmplData := assemblyCompletedTmplData{
		AssemblyUUID: message.AssemblyUUID,
		OrderUUID:    message.OrderUUID,
		BuildTimeSec: message.BuildTimeSec,
	}

	bytes := bytes.Buffer{}
	if err := assemblyCompletedTemplate.Execute(&bytes, tmplData); err != nil {
		return err
	}

	err := a.notificationSender.SendNotification(
		ctx,
		bytes.String(),
	)
	if err != nil {
		return nil
	}

	return nil
}
