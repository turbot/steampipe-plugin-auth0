package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Log() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_log",
		Description: "Logs are all the events that occur in your tenants including user authentication and administrative actions such as adding and updating applications, connections, and rules.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Logs,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Log,
			KeyColumns: plugin.SingleColumn("log_id"),
		},

		Columns: []*plugin.Column{
			{Name: "log_id", Description: "Log identifier", Type: proto.ColumnType_STRING, Transform: transform.FromField("LogID")},
			{Name: "date", Description: "The date when the log event was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "type", Description: "The log event type.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "The log event description.", Type: proto.ColumnType_STRING},
			{Name: "connection", Description: "Name of the connection the log event relates to.", Type: proto.ColumnType_STRING},
			{Name: "connection_id", Description: "ID of the connection the log event relates to.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ConnectionID")},
			{Name: "organization_id", Description: "ID of the organization the log event relates to.", Type: proto.ColumnType_STRING, Transform: transform.FromField("OrganizationID")},
			{Name: "organization_name", Description: "Name of the organization the log event relates to.", Type: proto.ColumnType_STRING},
			{Name: "client_id", Description: "The ID of the client (application).", Type: proto.ColumnType_STRING, Transform: transform.FromField("ClientID")},
			{Name: "client_name", Description: "The name of the client (application).", Type: proto.ColumnType_STRING},
			{Name: "ip", Description: "The IP address of the log event source.", Type: proto.ColumnType_STRING, Transform: transform.FromField("IP")},
			{Name: "hostname", Description: "Hostname the log event applies to.", Type: proto.ColumnType_STRING},
			{Name: "details", Description: "Additional useful details about this event (structure is dependent upon event type).", Type: proto.ColumnType_JSON},
			{Name: "user_id", Description: "ID of the user involved in the log event.", Type: proto.ColumnType_STRING, Transform: transform.FromField("UserID")},
			{Name: "user_name", Description: "Name of the user involved in the log event.", Type: proto.ColumnType_STRING},
			{Name: "user_agent", Description: "User agent string from the client device that caused the event.", Type: proto.ColumnType_STRING},
			{Name: "audience", Description: "API audience the event applies to.", Type: proto.ColumnType_STRING},
			{Name: "scope", Description: "Scope permissions applied to the event.", Type: proto.ColumnType_STRING},
			{Name: "strategy", Description: "Name of the strategy involved in the event.", Type: proto.ColumnType_STRING},
			{Name: "strategy_type", Description: "Type of strategy involved in the event.", Type: proto.ColumnType_STRING},
			{Name: "is_mobile", Description: "Whether the client was a mobile device (true) or desktop/laptop/server (false).", Type: proto.ColumnType_BOOL},
			{Name: "location_info", Description: "Information about the location that triggered this event based on the `IP`.", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listAuth0Logs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_log.listAuth0Logs", "connect_error", err)
		return nil, err
	}

	var pageNumber int
	for {
		logsResponse, err := client.Log.List(
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_log.listAuth0Logs", "query_error", err)
			return nil, err
		}
		for _, log := range logsResponse {
			d.StreamListItem(ctx, log)
			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
		if len(logsResponse) == 0 {
			break
		}
		pageNumber = pageNumber + 1
	}

	return nil, err
}

//// GET FUNCTION

func getAuth0Log(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("log_id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_log.getAuth0Log", "connect_error", err)
		return nil, err
	}

	logsResponse, err := client.Log.Read(id)
	if err != nil {
		logger.Error("auth0_log.getAuth0Log", "query_error", err)
		return nil, err
	}

	return logsResponse, nil
}
