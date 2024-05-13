package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "domain_name",
			Description: "The name of the domain.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getDomainName,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getDomainNameMemoized = plugin.HydrateFunc(getDomainNameUncached).Memoize(memoize.WithCacheKeyFunction(getDomainNameCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getDomainName(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getDomainNameMemoized(ctx, d, h)
}

// Build a cache key for the call to getDomainNameCacheKey.
func getDomainNameCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getDomainName"
	return key, nil
}

func getDomainNameUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	auth0Config := GetConfig(d.Connection)

	return auth0Config.Domain, nil
}
