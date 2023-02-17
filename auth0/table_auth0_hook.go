package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Hook() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_hook",
		Description: "Hooks allow you to customize the behavior of Auth0 with Node.js code that is executed in selected extension points.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Hooks,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Hook,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "A unique ID for the hook.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the hook.", Type: proto.ColumnType_STRING},
			{Name: "script", Description: "A script that contains the hook's code.", Type: proto.ColumnType_STRING},
			{Name: "trigger_id", Description: "The extensibility point name. Can currently be any of the following: credentials-exchange, pre-user-registration, post-user-registration, post-change-password.", Type: proto.ColumnType_STRING, Transform: transform.FromField("TriggerID")},
			{Name: "dependencies", Description: "Used to store additional metadata.", Type: proto.ColumnType_JSON},
			{Name: "enabled", Description: "Enabled should be set to true if the hook is enabled, false otherwise.", Type: proto.ColumnType_BOOL},
		},
	}
}

//// LIST FUNCTION

func listAuth0Hooks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_hook.listAuth0Hooks", "connect_error", err)
		return nil, err
	}

	perPage := 50
	// Limit indicates the number of records to return at once.
	// If the query limit is less than the API limit, then make API call limit to match query limit.
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < int64(perPage) {
			perPage = int(*limit)
		}
	}

	var pageNumber int
	for {
		hooksResponse, err := client.Hook.List(
			management.Page(pageNumber),
			management.PerPage(perPage),
		)
		if err != nil {
			logger.Error("auth0_hook.listAuth0Hooks", "query_error", err)
			return nil, err
		}
		for _, hook := range hooksResponse.Hooks {
			d.StreamListItem(ctx, hook)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if hooksResponse.Start+len(hooksResponse.Hooks) >= hooksResponse.Total {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}

//// GET FUNCTION

func getAuth0Hook(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_hook.getAuth0Hook", "connect_error", err)
		return nil, err
	}

	hooksResponse, err := client.Hook.Read(id)
	if err != nil {
		logger.Error("auth0_hook.getAuth0Hook", "query_error", err)
		return nil, err
	}

	return hooksResponse, nil
}
