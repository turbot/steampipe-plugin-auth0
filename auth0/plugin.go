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
		Name:             pluginName,
		DefaultTransform: transform.FromCamel(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"auth0_client":       tableAuth0Client(),
			"auth0_connection":   tableAuth0Connection(),
			"auth0_hook":         tableAuth0Hook(),
			"auth0_log":          tableAuth0Log(),
			"auth0_organization": tableAuth0Organization(),
			"auth0_role":         tableAuth0Role(),
			"auth0_user":         tableAuth0User(),
		},
	}

	return p
}
