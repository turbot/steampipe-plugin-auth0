package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Action() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_action",
		Description: "Actions are secure, tenant-specific, versioned functions written in Node.js that execute at certain points within the Auth0 platform.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Actions,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Action,
			KeyColumns: plugin.SingleColumn("id"),
		},

		Columns: []*plugin.Column{
			{Name: "id", Description: "A unique ID for the action.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ID")},
			{Name: "name", Description: "The name of the action.", Type: proto.ColumnType_STRING},
			{Name: "supported_triggers", Description: "List of triggers that this action supports. At this time, an action can only target a single trigger at a time.", Type: proto.ColumnType_JSON},
			{Name: "code", Description: "The source code of the action.", Type: proto.ColumnType_STRING},
			{Name: "dependencies", Description: "List of third party npm modules, and their versions, that this action depends on.", Type: proto.ColumnType_JSON},
			{Name: "runtime", Description: "The Node runtime. For example `node16`, defaults to `node12`.", Type: proto.ColumnType_STRING},
			{Name: "secrets", Description: "List of secrets that are included in an action or a version of an action.", Type: proto.ColumnType_JSON},
			{Name: "deployed_version", Description: "Version of the action that is currently deployed.", Type: proto.ColumnType_JSON},
			{Name: "status", Description: "The build status of this action.", Type: proto.ColumnType_STRING},
			{Name: "all_changes_deployed", Description: "True if all of an Action's contents have been deployed.", Type: proto.ColumnType_BOOL},
			{Name: "built_at", Description: "The time when this action was built successfully.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "created_at", Description: "The time when this action was created.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "updated_at", Description: "The time when this action was updated.", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

//// LIST FUNCTION

func listAuth0Actions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_action.listAuth0Actions", "connect_error", err)
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
		actionsResponse, err := client.Action.List(
			management.Page(pageNumber),
			management.PerPage(perPage),
		)
		if err != nil {
			logger.Error("auth0_action.listAuth0Actions", "query_error", err)
			return nil, err
		}
		for _, action := range actionsResponse.Actions {
			d.StreamListItem(ctx, action)

			// Context can be cancelled due to manual cancellation or the limit has been hit
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		if actionsResponse.Start+len(actionsResponse.Actions) >= actionsResponse.Total {
			break
		}
		pageNumber = pageNumber + 1
	}
	return nil, err
}

//// GET FUNCTION

func getAuth0Action(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("id")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_action.getAuth0Action", "connect_error", err)
		return nil, err
	}

	actionsResponse, err := client.Action.Read(id)
	if err != nil {
		logger.Error("auth0_action.getAuth0Action", "query_error", err)
		return nil, err
	}

	return actionsResponse, nil
}
