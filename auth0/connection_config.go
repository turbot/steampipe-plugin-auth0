package auth0

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type auth0Config struct {
	Domain       *string `hcl:"domain"`
	ClientId     *string `hcl:"client_id"`
	ClientSecret *string `hcl:"client_secret"`
	ApiToken     *string `hcl:"api_token"`
}

func ConfigInstance() interface{} {
	return &auth0Config{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) auth0Config {
	if connection == nil || connection.Config == nil {
		return auth0Config{}
	}
	config, _ := connection.Config.(auth0Config)
	return config
}
