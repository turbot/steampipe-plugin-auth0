package main

import (
	"github.com/turbot/steampipe-plugin-auth0/auth0"

	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: auth0.Plugin})
}
