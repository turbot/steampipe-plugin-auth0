package auth0

import (
	"context"
	"fmt"
	"os"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Connect(ctx context.Context, d *plugin.QueryData) (*management.Management, error) {
	// have we already created and cached the session?
	sessionCacheKey := "Auth0Session"
	if cachedData, ok := d.ConnectionManager.Cache.Get(sessionCacheKey); ok {
		return cachedData.(*management.Management), nil
	}

	auth0Config := GetConfig(d.Connection)

	var domain, clientId, clientSecret, apiToken string

	if auth0Config.Domain != nil {
		domain = *auth0Config.Domain
	} else {
		domain = os.Getenv("AUTH0_DOMAIN")
	}
	if auth0Config.ClientId != nil {
		clientId = *auth0Config.ClientId
	} else {
		clientId = os.Getenv("AUTH0_CLIENT_ID")
	}
	if auth0Config.ClientSecret != nil {
		clientSecret = *auth0Config.ClientSecret
	} else {
		clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	}
	if auth0Config.ApiToken != nil {
		apiToken = *auth0Config.ApiToken
	} else {
		apiToken = os.Getenv("AUTH0_API_TOKEN")
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
		// Save session into cache
		d.ConnectionManager.Cache.Set(sessionCacheKey, m)
		return m, nil
	}

	m, err := management.New(domain, management.WithClientCredentials(clientId, clientSecret))
	if err != nil {
		return nil, err
	}
	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, m)
	return m, nil
}
