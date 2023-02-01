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
			Hydrate:    getAuth0Roles,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "A unique ID for the role.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the role.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "A description of the role.", Type: proto.ColumnType_STRING},
			{Name: "permissions", Description: "List of permissions of the role.", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Hydrate: getAuth0RolesPermissions},
			{Name: "users", Description: "List of users assigned to the role.", Type: proto.ColumnType_JSON, Transform: transform.FromValue(), Hydrate: getAuth0RolesUsers},
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

	var pageNumber int
	for {
		rolesResponse, err := client.Role.List(
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_role.listAuth0Roles", "list_roles_error", err)
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

func getAuth0Roles(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role.getAuth0Roles", "connect_error", err)
		return nil, err
	}

	rolesResponse, err := client.Role.Read(id)
	if err != nil {
		logger.Error("auth0_role.getAuth0Roles", "get_roles_error", err)
		return nil, err
	}

	return rolesResponse, nil
}

//// GET ROLES PERMISSIONS FUNCTION

func getAuth0RolesPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	role := h.Item.(*management.Role)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role.getAuth0RolesPermissions", "connect_error", err)
		return nil, err
	}

	var permissions []*management.Permission
	var pageNumber, perPage int
	perPage = 50
	for {
		permissionsResponse, err := client.Role.Permissions(
			*role.ID,
			management.PerPage(perPage),
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_role.getAuth0RolesPermissions", "get_roles_permissions_error", err)
			return nil, err
		}
		permissions = append(permissions, permissionsResponse.Permissions...)

		if len(permissionsResponse.Permissions) == 0 {
			break
		}
		pageNumber = pageNumber + 1
	}
	return permissions, err
}

//// GET ROLES USERS FUNCTION

func getAuth0RolesUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	role := h.Item.(*management.Role)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_role.getAuth0RolesUsers", "connect_error", err)
		return nil, err
	}

	var users []*management.User
	var pageNumber, perPage int
	perPage = 50
	for {
		usersResponse, err := client.Role.Users(
			*role.ID,
			management.PerPage(perPage),
			management.Page(pageNumber),
		)
		if err != nil {
			logger.Error("auth0_role.getAuth0RolesUsers", "get_roles_users_error", err)
			return nil, err
		}
		users = append(users, usersResponse.Users...)

		if len(usersResponse.Users) == 0 {
			break
		}
		pageNumber = pageNumber + 1
	}
	return users, err
}
