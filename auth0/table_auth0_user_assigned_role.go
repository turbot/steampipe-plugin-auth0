package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0UserAssignedRole() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_user_assigned_role",
		Description: "List of roles assigned to user.",
		List: &plugin.ListConfig{
			Hydrate:    listAuth0UserAssignedRoles,
			KeyColumns: plugin.SingleColumn("user_id"),
		},
		Columns: []*plugin.Column{
			{Name: "user_id", Description: "A unique ID for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromQual("user_id")},
			{Name: "role_id", Description: "A unique ID for the role.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the role.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the role.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listAuth0UserAssignedRoles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_user_assigned_role.listAuth0UserAssignedRoles", "connect_error", err)
		return nil, err
	}

	userId := d.EqualsQualString("user_id")

	var pageNumber, perPage int
	perPage = 50
	for {
		rolesResponse, err := client.User.Roles(
			userId,
			management.PerPage(perPage),
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_user_assigned_role.listAuth0UserAssignedRoles", "query_error", err)
			return nil, err
		}

		for _, role := range rolesResponse.Roles {
			d.StreamListItem(ctx, role)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(rolesResponse.Roles) == 0 {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}
