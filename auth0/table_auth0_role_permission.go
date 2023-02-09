package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0RolePermission() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_role_permission",
		Description: "",
		List: &plugin.ListConfig{
			Hydrate:    listAuth0RolePermissions,
			KeyColumns: plugin.SingleColumn("role_id"),
		},
		Columns: []*plugin.Column{
			{Name: "role_id", Description: "A unique ID for the role.", Type: proto.ColumnType_STRING, Transform: transform.FromQual("role_id")},
			{Name: "permission_name", Description: "The name of the permission.", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name")},
			{Name: "description", Description: "The description of the permission.", Type: proto.ColumnType_STRING},
			{Name: "resource_server_identifier", Description: "The resource server that the permission is attached to.", Type: proto.ColumnType_STRING},
			{Name: "resource_server_name", Description: "The name of the resource server.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listAuth0RolePermissions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role_permission.listAuth0RolePermissions", "connect_error", err)
		return nil, err
	}

	roleId := d.EqualsQualString("role_id")

	var pageNumber, perPage int
	perPage = 50
	for {
		permissionsResponse, err := client.Role.Permissions(
			roleId,
			management.PerPage(perPage),
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_role_permission.listAuth0RolePermissions", "list_role_permissions_error", err)
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
