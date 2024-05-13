package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0UserPermission() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_user_permission",
		Description: "Permissions assigned to user, including permissions assigned through users roles.",
		List: &plugin.ListConfig{
			Hydrate:    listAuth0UserPermissions,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "user_id", Description: "A unique ID for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromQual("user_id")},
			{Name: "permission_name", Description: "The name of the permission.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
			{Name: "description", Description: "The description of the permission.", Type: proto.ColumnType_STRING},
			{Name: "resource_server_identifier", Description: "The resource server that the permission is attached to.", Type: proto.ColumnType_STRING},
			{Name: "resource_server_name", Description: "The name of the resource server.", Type: proto.ColumnType_STRING},
		}),
	}
}

//// LIST FUNCTION

func listAuth0UserPermissions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_user_permission.listAuth0UserPermissions", "connect_error", err)
		return nil, err
	}

	userId := d.EqualsQualString("user_id")

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
		permissionsResponse, err := client.User.Permissions(
			userId,
			management.PerPage(perPage),
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_user_permission.listAuth0UserPermissions", "query_error", err)
			return nil, err
		}

		for _, permission := range permissionsResponse.Permissions {
			d.StreamListItem(ctx, permission)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(permissionsResponse.Permissions) == 0 {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}
