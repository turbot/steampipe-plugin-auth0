package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0User() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_user",
		Description: "User represents an Auth0 user resource.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Users,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Users,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "The users' identifier.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "connection", Description: "The connection the user belongs to.", Type: proto.ColumnType_STRING},
			{Name: "email", Description: "The users' email.", Type: proto.ColumnType_STRING},
			{Name: "name", Description: "The users' name.", Type: proto.ColumnType_STRING},
			{Name: "given_name", Description: "The users' given name.", Type: proto.ColumnType_STRING},
			{Name: "family_name", Description: "The users' family name.", Type: proto.ColumnType_STRING},
			{Name: "username", Description: "The users' username. Only valid if the connection requires a username.", Type: proto.ColumnType_STRING},
			{Name: "nickname", Description: "The users' nickname.", Type: proto.ColumnType_STRING},
			{Name: "screen_name", Description: "The screen name, handle, or alias that this user identifies themselves with.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "The user-defined UTF-8 string describing their account.", Type: proto.ColumnType_STRING},
			{Name: "location", Description: "The user-defined location for this account’s profile.", Type: proto.ColumnType_STRING},
			{Name: "password", Description: "The users' password (mandatory for non SMS connections)", Type: proto.ColumnType_STRING},
			{Name: "phone_number", Description: "The users' phone number (following the E.164 recommendation).", Type: proto.ColumnType_STRING},
			{Name: "created_at", Description: "The time the user was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "updated_at", Description: "The last time the user was updated.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "last_login", Description: "The last time the user has logged in.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "last_password_reset", Description: "The last time the user had their password reset.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "user_metadata", Description: "UserMetadata holds data that the user has read/write access to.", Type: proto.ColumnType_JSON},
			{Name: "identities", Description: "Identities is a list of user identities for when accounts are linked.", Type: proto.ColumnType_JSON},
			{Name: "email_verified", Description: "True if the user's email is verified, false otherwise.", Type: proto.ColumnType_BOOL},
			{Name: "verify_email", Description: "If true, the user will receive a verification email after creation, even if created with email_verified set to true.", Type: proto.ColumnType_BOOL},
			{Name: "phone_verified", Description: "True if the user's phone number is verified, false otherwise.", Type: proto.ColumnType_BOOL},
			{Name: "app_metadata", Description: "Holds data that the user has read-only access to.", Type: proto.ColumnType_JSON},
			{Name: "picture", Description: "The user's picture url.", Type: proto.ColumnType_STRING},
			{Name: "url", Description: "A URL provided by the user in association with their profile.", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL")},
			{Name: "blocked", Description: "True if the user is blocked from the application, false if the user is enabled.", Type: proto.ColumnType_BOOL},
			{Name: "last_ip", Description: "Last IP address from which this user logged in. Read only, cannot be modified.", Type: proto.ColumnType_STRING, Transform: transform.FromField("LastIP")},
			{Name: "logins_count", Description: "Total number of logins this user has performed. Read only, cannot be modified.", Type: proto.ColumnType_INT},
			{Name: "multifactor", Description: "List of multi-factor authentication providers with which this user has enrolled.", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listAuth0Users(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_user.listAuth0Users", "connect_error", err)
		return nil, err
	}

	usersResponse, err := client.User.List()
	if err != nil {
		logger.Error("auth0_user.listAuth0Users", "list_users_error", err)
		return nil, err
	}
	for _, user := range usersResponse.Users {
		d.StreamListItem(ctx, user)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// GET FUNCTION

func getAuth0Users(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_user.getAuth0Users", "connect_error", err)
		return nil, err
	}

	usersResponse, err := client.User.Read(id)
	if err != nil {
		logger.Error("auth0_user.getAuth0Users", "get_users_error", err)
		return nil, err
	}

	return usersResponse, nil
}
