package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Connection() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_connection",
		Description: "Connection is the relationship between Auth0 and a source of users.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Connections,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Connection,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: commonColumns([]*plugin.Column{
			{Name: "id", Description: "A generated string identifying the connection.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the connection.", Type: proto.ColumnType_STRING},
			{Name: "display_name", Description: "The display name of the connection.", Type: proto.ColumnType_STRING},
			{Name: "strategy", Description: "The identity provider identifier for the connection.", Type: proto.ColumnType_STRING},
			{Name: "is_domain_connection", Description: "True if the connection is domain level.", Type: proto.ColumnType_BOOL},
			{Name: "options", Description: "Options for validation.", Type: proto.ColumnType_JSON},
			{Name: "enabled_clients", Description: "The identifiers of the clients for which the connection is to be enabled.", Type: proto.ColumnType_JSON},
			{Name: "realms", Description: "Defines the realms for which the connection will be used (ie: email domains).", Type: proto.ColumnType_JSON},
			{Name: "metadata", Description: "Metadata of the connection.", Type: proto.ColumnType_JSON},
			{Name: "provisioning_ticket_url", Description: "Provisioning Ticket URL is Ticket URL for Active Directory/LDAP, etc.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ProvisioningTicketURL")},
			{Name: "show_as_button", Description: "Display connection as a button.", Type: proto.ColumnType_BOOL},
		}),
	}
}

//// LIST FUNCTION

func listAuth0Connections(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_connection.listAuth0Connections", "connect_error", err)
		return nil, err
	}

	perPage := 50
	// Limit indicates the number of records to return at once.
	// If the query limit is less than the API limit, then make API call limit to match query limit.
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(perPage) {
			perPage = int(*limit)
		}
	}

	var pageNumber int
	for {
		connectionsResponse, err := client.Connection.List(
			management.Page(pageNumber),
			management.PerPage(perPage),
		)
		if err != nil {
			logger.Error("auth0_connection.listAuth0Connections", "query_error", err)
			return nil, err
		}
		for _, connection := range connectionsResponse.Connections {
			d.StreamListItem(ctx, connection)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if connectionsResponse.Start+len(connectionsResponse.Connections) >= connectionsResponse.Total {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}

//// GET FUNCTION

func getAuth0Connection(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_connection.getAuth0Connection", "connect_error", err)
		return nil, err
	}

	connectionsResponse, err := client.Connection.Read(id)
	if err != nil {
		logger.Error("auth0_connection.getAuth0Connection", "query_error", err)
		return nil, err
	}

	return connectionsResponse, nil
}
