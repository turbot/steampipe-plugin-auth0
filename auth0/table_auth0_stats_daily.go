package auth0

import (
	"context"

	"github.com/auth0/go-auth0/management"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

const (
	YYYYMMDD = "20060102"
)

//// TABLE DEFINITION

func tableAuth0StatsDaily() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_stats_daily",
		Description: "Daily stats for an Auth0 Tenant.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0StatsDaily,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "date", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional},
			},
		},

		Columns: []*plugin.Column{
			{Name: "date", Description: "Date of the stats.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "logins", Description: "Number of logins.", Type: proto.ColumnType_INT},
			{Name: "sign_ups", Description: "Number of sign ups.", Type: proto.ColumnType_INT, Transform: transform.FromField("Signups")},
			{Name: "leaked_passwords", Description: "Number of leaked passwords.", Type: proto.ColumnType_INT},
			{Name: "created_at", Description: "Timestamp of when the stat started to count.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "updated_at", Description: "Timestamp of the last update to stat of the day.", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

//// LIST FUNCTION

func listAuth0StatsDaily(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_stats_daily.listAuth0StatsDaily", "connect_error", err)
		return nil, err
	}

	var opts []management.RequestOption
	quals := d.Quals
	if quals["date"] != nil {
		for _, q := range quals["date"].Quals {
			qDate := q.Value.GetTimestampValue().AsTime().Format(YYYYMMDD)
			switch q.Operator {
			case "=":
				opts = append(opts, management.From(qDate))
				opts = append(opts, management.Parameter("to", qDate))
			case ">=", ">":
				opts = append(opts, management.From(qDate))
			case "<=", "<":
				opts = append(opts, management.Parameter("to", qDate))
			}
		}
	}

	stats, err := client.Stat.Daily(opts...)
	if err != nil {
		logger.Error("auth0_stats_daily.listAuth0StatsDaily", "query_error", err)
		return nil, err
	}
	for _, stat := range stats {
		d.StreamListItem(ctx, stat)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}
