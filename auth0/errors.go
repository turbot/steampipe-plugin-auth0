package auth0

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func shouldRetryError(retryErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {

		if strings.Contains(err.Error(), "429") {
			plugin.Logger(ctx).Debug("auth0_errors.shouldRetryError", "rate_limit_error", err)
			return true
		}
		return false
	}
}
