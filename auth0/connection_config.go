package auth0

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type auth0Config struct {
	Domain   *string `cty:"domain"`
	ClientId *string `cty:"client_id"`
	Secret   *string `cty:"secret"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"domain": {
		Type: schema.TypeString,
	},
	"client_id": {
		Type: schema.TypeString,
	},
	"secret": {
		Type: schema.TypeString,
	},
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
