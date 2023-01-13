package auth0

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableAuth0Client() *plugin.Table {
	return &plugin.Table{
		Name:        "auth0_client",
		Description: "Client is an application or a sso integration.",
		List: &plugin.ListConfig{
			Hydrate: listAuth0Clients,
		},
		Get: &plugin.GetConfig{
			Hydrate:    getAuth0Clients,
			KeyColumns: plugin.SingleColumn("client_id"),
		},

		Columns: []*plugin.Column{
			{Name: "client_id", Description: "The ID of the client.", Type: proto.ColumnType_STRING, Transform: transform.FromField("ClientID")},
			{Name: "name", Description: "The name of the client.", Type: proto.ColumnType_STRING},
			{Name: "description", Description: "Free text description of the purpose of the Client.", Type: proto.ColumnType_STRING},
			{Name: "client_secret", Description: "The client secret, it must not be public.", Type: proto.ColumnType_STRING},
			{Name: "app_type", Description: "The type of application this client represents.", Type: proto.ColumnType_STRING},
			{Name: "logo_uri", Description: "The URL of the client logo (recommended size: 150x150).", Type: proto.ColumnType_STRING, Transform: transform.FromField("LogoURI")},
			{Name: "is_first_party", Description: "Whether this client a first party client or not.", Type: proto.ColumnType_BOOL},
			{Name: "is_token_endpoint_ip_header_trusted", Description: "Set header `auth0-forwarded-for` as trusted to be used as source of end user ip for brute-force-protection on token endpoint.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("IsTokenEndpointIPHeaderTrusted")},
			{Name: "oidc_conformant", Description: "Whether this client will conform to strict OIDC specifications.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("IsTokenEndpointIPHeaderTrusted")},
			{Name: "callbacks", Description: "The URLs that Auth0 can use to as a callback for the client.", Type: proto.ColumnType_JSON},
			{Name: "allowed_origins", Description: "The allowed origin URLs.", Type: proto.ColumnType_JSON},
			{Name: "web_origins", Description: "A set of URLs that represents valid web origins for use with web message response mode.", Type: proto.ColumnType_JSON},
			{Name: "client_aliases", Description: "Client aliases.", Type: proto.ColumnType_JSON},
			{Name: "allowed_clients", Description: "Allowed clients.", Type: proto.ColumnType_JSON},
			{Name: "allowed_logout_urls", Description: "A set of URLs that are valid to redirect to after logout from Auth0.", Type: proto.ColumnType_JSON, Transform: transform.FromField("AllowedLogoutURLs")},
			{Name: "jwt_configuration", Description: "JSON web token configuration.", Type: proto.ColumnType_JSON, Transform: transform.FromField("JWTConfiguration")},
			{Name: "signing_keys", Description: "Client signing keys.", Type: proto.ColumnType_JSON},
			{Name: "encryption_key", Description: "Client encryption key.", Type: proto.ColumnType_JSON},
			{Name: "sso", Description: "Client single sign-on.", Type: proto.ColumnType_BOOL, Transform: transform.FromField("SSO")},
			{Name: "sso_disabled", Description: "True to disable Single Sign On, false otherwise (default: false).", Type: proto.ColumnType_BOOL, Transform: transform.FromField("SSODisabled")},
			{Name: "cross_origin_auth", Description: "True if this client can be used to make cross-origin authentication requests, false otherwise (default: false).", Type: proto.ColumnType_BOOL},
			{Name: "grant_types", Description: "List of acceptable Grant Types for this Client.", Type: proto.ColumnType_JSON},
			{Name: "cross_origin_location", Description: "URL for the location in your site where the cross origin verification takes place for the cross-origin auth flow when performing Auth in your own domain instead of Auth0 hosted login page.", Type: proto.ColumnType_STRING},
			{Name: "CustomLoginPageOn", Description: "True if the custom login page is to be used, false otherwise. Defaults to true.", Type: proto.ColumnType_BOOL},
			{Name: "CustomLoginPage", Description: "The custom login page to be used.", Type: proto.ColumnType_STRING},
			{Name: "CustomLoginPagePreview", Description: "The custom login page preview to be used.", Type: proto.ColumnType_STRING},
			{Name: "FormTemplate", Description: "The form template to be used.", Type: proto.ColumnType_STRING},
			{Name: "Addons", Description: "Addons for our client.", Type: proto.ColumnType_JSON},
			{Name: "TokenEndpointAuthMethod", Description: "Defines the requested authentication method for the token endpoint.", Type: proto.ColumnType_STRING},
			{Name: "ClientMetadata", Description: "Metadata associated with the client.", Type: proto.ColumnType_JSON},
			{Name: "Mobile", Description: "Mobile app settings.", Type: proto.ColumnType_JSON},
			{Name: "InitiateLoginURI", Description: "Initiate login uri, must be https and cannot contain a fragment.", Type: proto.ColumnType_STRING, Transform: transform.FromField("InitiateLoginURI")},
			{Name: "NativeSocialLogin", Description: "Native Social Login settings.", Type: proto.ColumnType_JSON},
			{Name: "RefreshToken", Description: "Refresh Token settings for our Client.", Type: proto.ColumnType_JSON},
			{Name: "OrganizationUsage", Description: "Organization Usage.", Type: proto.ColumnType_STRING},
			{Name: "OrganizationRequireBehavior", Description: "Organization Require Behavior.", Type: proto.ColumnType_STRING},
		},
	}
}

//// LIST FUNCTION

func listAuth0Clients(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_client.listAuth0Clients", "connect_error", err)
		return nil, err
	}

	clientsResponse, err := client.Client.List()
	if err != nil {
		logger.Error("auth0_client.listAuth0Clients", "list_clients_error", err)
		return nil, err
	}
	for _, client := range clientsResponse.Clients {
		d.StreamListItem(ctx, client)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, err
}

//// GET FUNCTION

func getAuth0Clients(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)
	clientId := d.EqualsQualString("client_id")

	// Empty check for clientId
	if clientId == "" {
		return nil, nil
	}

	client, err := Connect(ctx, d)
	if err != nil {
		logger.Error("auth0_client.getAuth0Clients", "connect_error", err)
		return nil, err
	}

	clientsResponse, err := client.Client.Read(clientId)
	if err != nil {
		logger.Error("auth0_client.getAuth0Clients", "get_clients_error", err)
		return nil, err
	}

	return clientsResponse, nil
}
