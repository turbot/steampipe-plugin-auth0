package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const pluginName = "steampipe-plugin-auth0"

// Plugin creates this (auth0) plugin
func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:               pluginName,
		DefaultTransform:   transform.FromCamel(),
		DefaultRetryConfig: &plugin.RetryConfig{ShouldRetryErrorFunc: shouldRetryError([]string{"429"})},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"auth0_action":             tableAuth0Action(),
			"auth0_client":             tableAuth0Client(),
			"auth0_connection":         tableAuth0Connection(),
			"auth0_hook":               tableAuth0Hook(),
			"auth0_log":                tableAuth0Log(),
			"auth0_organization":       tableAuth0Organization(),
			"auth0_role_assigned_user": tableAuth0RoleAssignedUser(),
			"auth0_role_permission":    tableAuth0RolePermission(),
			"auth0_role":               tableAuth0Role(),
			"auth0_signing_key":        tableAuth0SigningKey(),
			"auth0_stats_daily":        tableAuth0StatsDaily(),
			"auth0_tenant_settings":    tableAuth0TenantSettings(),
			"auth0_user_assigned_role": tableAuth0UserAssignedRole(),
			"auth0_user_permission":    tableAuth0UserPermission(),
			"auth0_user":               tableAuth0User(),
		},
	}

	return p
}
