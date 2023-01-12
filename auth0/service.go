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

	var domain, clientId, secret string

	if auth0Config.Domain != nil {
		domain = *auth0Config.Domain
	} else {
		// TODO review correct env var names
		domain = os.Getenv("AUTH0_DOMAIN")
	}
	if auth0Config.ClientId != nil {
		clientId = *auth0Config.ClientId
	} else {
		// TODO review correct env var names
		clientId = os.Getenv("AUTH0_CLIENT_ID")
	}
	if auth0Config.Secret != nil {
		secret = *auth0Config.Secret
	} else {
		// TODO review correct env var names
		secret = os.Getenv("AUTH0_SECRET")
	}

	// No creds
	if domain == "" {
		return nil, fmt.Errorf("domain must be configured")
	}
	if clientId == "" {
		return nil, fmt.Errorf("client_id must be configured")
	}
	if secret == "" {
		return nil, fmt.Errorf("secret must be configured")
	}

	m, err := management.New(domain, management.WithClientCredentials(clientId, secret))
	if err != nil {
		return nil, err
	}

	// Save session into cache
	d.ConnectionManager.Cache.Set(sessionCacheKey, m)

	return m, nil

}
