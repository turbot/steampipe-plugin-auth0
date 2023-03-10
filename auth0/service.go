package auth0

import (
	"context"
	"fmt"
	"os"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*management.Management, error) {
	conn, err := connectCached(ctx, d, nil)
	if err != nil {
		return nil, err
	}
	return conn.(*management.Management), nil
}

var connectCached = plugin.HydrateFunc(connectUncached).Memoize()

func connectUncached(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (any, error) {
	auth0Config := GetConfig(d.Connection)

	var domain, clientId, clientSecret, apiToken string

	domain = os.Getenv("AUTH0_DOMAIN")
	if auth0Config.Domain != nil {
		domain = *auth0Config.Domain
	} 
	clientId = os.Getenv("AUTH0_CLIENT_ID")
	if auth0Config.ClientId != nil {
		clientId = *auth0Config.ClientId
	} 
	clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	if auth0Config.ClientSecret != nil {
		clientSecret = *auth0Config.ClientSecret
	} 
	apiToken = os.Getenv("AUTH0_API_TOKEN")
	if auth0Config.ApiToken != nil {
		apiToken = *auth0Config.ApiToken
	} 

	// No creds
	if domain == "" {
		return nil, fmt.Errorf("domain must be configured")
	}

	if apiToken == "" {
		if clientId == "" {
			return nil, fmt.Errorf("either api_token or client_id must be configured")
		}
		if clientSecret == "" {
			return nil, fmt.Errorf("client_secret must be configured")
		}
	}

	if apiToken != "" {
		m, err := management.New(domain, management.WithStaticToken(apiToken))
		if err != nil {
			return nil, err
		}
		return m, nil
	}

	m, err := management.New(domain, management.WithClientCredentials(clientId, clientSecret))
	if err != nil {
		return nil, err
	}
	return m, nil
}
