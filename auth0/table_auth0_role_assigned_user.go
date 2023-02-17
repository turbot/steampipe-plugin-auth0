package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0RoleAssignedUser() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_role_assigned_user",
		Description: "List of users assigned to role.",
		List: &plugin.ListConfig{
			Hydrate:    listAuth0RoleAssignedUsers,
			KeyColumns: plugin.SingleColumn("role_id"),
		},
		Columns: []*plugin.Column{
			{Name: "role_id", Description: "A unique ID for the role.", Type: proto.ColumnType_STRING, Transform: transform.FromQual("role_id")},
			{Name: "user_id", Description: "A unique ID for the user.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the user.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "The users' email.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listAuth0RoleAssignedUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role_assigned_user.listAuth0RoleAssignedUsers", "connect_error", err)
		return nil, err
	}

	roleId := d.EqualsQualString("role_id")

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
		usersResponse, err := client.Role.Users(
			roleId,
			management.PerPage(perPage),
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_role_assigned_user.listAuth0RoleAssignedUsers", "query_error", err)
			return nil, err
		}

		for _, role := range usersResponse.Users {
			d.StreamListItem(ctx, role)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if len(usersResponse.Users) == 0 {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}
