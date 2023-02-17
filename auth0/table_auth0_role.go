package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Role() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_role",
		Description: "Role is a collection of permissions that you can apply to user.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Roles,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Role,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "A unique ID for the role.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the role.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the role.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listAuth0Roles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role.listAuth0Roles", "connect_error", err)
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
		rolesResponse, err := client.Role.List(
			management.Page(pageNumber),
			management.PerPage(perPage),
		)
		if err != nil {
			logger.Error("auth0_role.listAuth0Roles", "query_error", err)
			return nil, err
		}
		for _, role := range rolesResponse.Roles {
			d.StreamListItem(ctx, role)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if rolesResponse.Start+len(rolesResponse.Roles) >= rolesResponse.Total {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}

//// GET FUNCTION

func getAuth0Role(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role.getAuth0Role", "connect_error", err)
		return nil, err
	}

	rolesResponse, err := client.Role.Read(id)
	if err != nil {
		logger.Error("auth0_role.getAuth0Role", "query_error", err)
		return nil, err
	}

	return rolesResponse, nil
}
