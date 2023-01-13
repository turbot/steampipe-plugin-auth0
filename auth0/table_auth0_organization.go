package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Organization() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_organization",
		Description: "Organization is used to allow B2B customers to better manage their partners and customers, and to customize the ways that end-users access their applications.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Organizations,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Organizations,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "Organization identifier", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "Name of this organization.", Type: proto.ColumnType_STRING},
			{Name: "display_name", Description: "DisplayName of this organization.", Type: proto.ColumnType_STRING},
			{Name: "branding", Description: "Branding defines how to style the login pages", Type: proto.ColumnType_STRING},
			{Name: "metadata", Description: "Metadata associated with the organization.", Type: proto.ColumnType_STRING},
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

	var nextPage string
	for {
		organizationsResponse, err := client.Organization.List(
			management.From(nextPage),
		)
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
		if organizationsResponse.Next == "" {
			break
		}
		nextPage = organizationsResponse.Next
	}

	return nil, err
}

//// GET FUNCTION

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

	return organizationsResponse, nil
}
