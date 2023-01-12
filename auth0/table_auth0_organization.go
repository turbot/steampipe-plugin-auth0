package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Organization() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_organization",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Organizations,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Organizations,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "", Type: proto.ColumnType_STRING},
			{Name: "display_name", Description: "", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listAuth0Organizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_organization.listAuth0Organizations", "connect_error", err)
		return nil, err
	}

	organizationsResponse, err := client.Organization.List()
	if err != nil {
		logger.Error("auth0_organization.listAuth0Organizations", "list_organizations_error", err)
		return nil, err
	}
	for _, organization := range organizationsResponse.Organizations {
		d.StreamListItem(ctx, organization)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// LIST FUNCTION

func getAuth0Organizations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_organization.getAuth0Organizations", "connect_error", err)
		return nil, err
	}

	organizationsResponse, err := client.Organization.Read(id)
	if err != nil {
		logger.Error("auth0_organization.getAuth0Organizations", "get_organizations_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, organizationsResponse)
	return nil, err
}
