package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0SigningKey() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_signing_key",
		Description: "Signing keys are used to sign ID tokens, access tokens, SAML assertions, and WS-Fed assertions sent to your application or API.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0SigningKeys,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0SigningKeys,
			KeyColumns: plugin.SingleColumn("kid"),
		},

		Columns: []*plugin.Column{
			{Name: "kid", Description: "The key id of the signing key.", Type: proto.ColumnType_STRING, Transform: transform.FromField("KID")},
			{Name: "cert", Description: "The public certificate of the signing key.", Type: proto.ColumnType_STRING},
			{Name: "pkcs7", Description: "The public certificate of the signing key in pkcs7 format.", Type: proto.ColumnType_STRING, Transform: transform.FromField("PKCS7")},
			{Name: "current", Description: "True if the key is the the current key.", Type: proto.ColumnType_BOOL},
			{Name: "next", Description: "True if the key is the the next key.", Type: proto.ColumnType_BOOL},
			{Name: "previous", Description: "True if the key is the the previous key.", Type: proto.ColumnType_BOOL},
			{Name: "current_since", Description: "The date and time when the key became the current key.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "current_until", Description: "The date and time when the current key was rotated.", Type: proto.ColumnType_TIMESTAMP},
			{Name: "fingerprint", Description: "The cert fingerprint.", Type: proto.ColumnType_STRING},
			{Name: "thumbprint", Description: "The cert thumbprint.", Type: proto.ColumnType_STRING},
			{Name: "revoked", Description: "True if the key is revoked.", Type: proto.ColumnType_BOOL},
			{Name: "revoked_at", Description: "The date and time when the key was revoked.", Type: proto.ColumnType_TIMESTAMP},
		},
	}
}

//// LIST FUNCTION

func listAuth0SigningKeys(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_signing_key.listAuth0SigningKeys", "connect_error", err)
		return nil, err
	}

	signingKeysResponse, err := client.SigningKey.List()
	if err != nil {
		logger.Error("auth0_signing_key.listAuth0SigningKeys", "list_signing_keys_error", err)
		return nil, err
	}
	for _, signingKey := range signingKeysResponse {
		d.StreamListItem(ctx, signingKey)
		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}
	return nil, err
}

//// GET FUNCTION

func getAuth0SigningKeys(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	id := d.EqualsQualString("kid")

	// Empty check for id
	if id == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_signing_key.getAuth0SigningKeys", "connect_error", err)
		return nil, err
	}

	logsResponse, err := client.SigningKey.Read(id)
	if err != nil {
		logger.Error("auth0_signing_key.getAuth0SigningKeys", "get_signing_keys_error", err)
		return nil, err
	}

	return logsResponse, nil
}
