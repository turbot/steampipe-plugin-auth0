package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

//// TABLE DEFINITION

func tableAuth0TenantSettings() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_tenant_settings",
		Description: "Tenant settings definitions.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0TenantSettings,
		},

		Columns: []*plugin.Column{
			{Name: "change_password", Description: "Change password page settings.", Type: proto.ColumnType_JSON},
			{Name: "guardian_mfa_page", Description: "Guardian MFA page settings.", Type: proto.ColumnType_JSON},
			{Name: "default_audience", Description: "Default audience for API Authorization.", Type: proto.ColumnType_STRING},
			{Name: "default_directory", Description: "Name of the connection that will be used for password grants at the token endpoint.", Type: proto.ColumnType_STRING},
			{Name: "error_page", Description: "Settings for the error page.", Type: proto.ColumnType_JSON},
			{Name: "device_flow", Description: "Settings for device flow.", Type: proto.ColumnType_JSON},
			{Name: "flags", Description: "Tenant flags.", Type: proto.ColumnType_JSON},
			{Name: "friendly_name", Description: "The friendly name of the tenant.", Type: proto.ColumnType_STRING},
			{Name: "picture_url", Description: "The URL of the tenant logo.", Type: proto.ColumnType_STRING},
			{Name: "support_email", Description: "User support email.", Type: proto.ColumnType_STRING},
			{Name: "support_url", Description: "User support URL.", Type: proto.ColumnType_STRING},
			{Name: "universal_login", Description: "Used to store additional metadata.", Type: proto.ColumnType_JSON},
			{Name: "allowed_logout_urls", Description: "A set of URLs that are valid to redirect to after logout from Auth0.", Type: proto.ColumnType_JSON},
			{Name: "session_lifetime", Description: "Login session lifetime, how long the session will stay valid (hours).", Type: proto.ColumnType_DOUBLE},
			{Name: "idle_session_lifetime", Description: "Force a user to login after they have been inactive for the specified number (hours).", Type: proto.ColumnType_DOUBLE},
			{Name: "sandbox_version", Description: "The selected sandbox version to be used for the extensibility environment.", Type: proto.ColumnType_STRING},
			{Name: "sandbox_versions_available", Description: "A set of available sandbox versions for the extensibility environment.", Type: proto.ColumnType_JSON},
			{Name: "default_redirection_uri", Description: "The default absolute redirection uri, must be https and cannot contain a fragment.", Type: proto.ColumnType_STRING},
			{Name: "enabled_locales", Description: "Supported locales for the UI.", Type: proto.ColumnType_JSON},
			{Name: "session_cookie", Description: "Session cookie settings.", Type: proto.ColumnType_JSON},
		},
	}
}

//// LIST FUNCTION

func listAuth0TenantSettings(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_tenant_settings.listAuth0TenantSettings", "connect_error", err)
		return nil, err
	}

	tenantSettings, err := client.Tenant.Read()
	if err != nil {
		logger.Error("auth0_tenant_settings.listAuth0TenantSettings", "query_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, tenantSettings)

	return nil, err
}
