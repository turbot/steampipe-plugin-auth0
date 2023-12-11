---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/auth0.svg"
brand_color: "#635DFF"
display_name: "Auth0"
name: "auth0"
description: "Use SQL to query users, clients, connections, keys and more from Auth0."
og_description: "Query Auth0 with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/auth0-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Auth0 + Steampipe

[Auth0](https://www.auth0.com/) is a flexible, drop-in solution to add authentication and authorization services to your applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  email,
  id,
  updated_at
from
  auth0_user
where
  not email_verified;
```

```
+--------------------------------+--------------------------------+---------------------------+
| email                          | id                             | updated_at                |
+--------------------------------+--------------------------------+---------------------------+
| select-joey@mail.com           | auth0|63c1e7732fbd7515d3fbe622 | 2023-01-19T12:56:32-04:00 |
| brief-ocelot@coffeetech.com.br | auth0|63d1a8b93addd17774ec0302 | 2022-07-28T00:28:13-04:00 |
| enough-snake@mail.com          | auth0|63c1e75e20e950c9f84d9ca5 | 2020-10-23T20:35:35-04:00 |
+--------------------------------+--------------------------------+---------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/auth0/tables)**

## Get started

### Install

Download and install the latest Auth0 plugin:

```bash
steampipe plugin install auth0
```

### Credentials

| Item        | Description                                                                                                                                                             |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | Auth0 requires an [API token or a Client ID and Client Secret](https://auth0.com/docs/secure/tokens/access-tokens/management-api-access-tokens) for all requests.       |
| Permissions | API tokens have the same permissions as the user who creates them, and if the user permissions change, the API token permissions also change.                           |
| Radius      | Each connection represents a single Auth0 Installation.                                                                                                                 |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/auth0.spc`)<br />2. Credentials specified in environment variables, e.g., `AUTH0_TOKEN`. |

### Configuration

Installing the latest auth0 plugin will create a config file (~/.steampipe/config/auth0.spc) with a single connection named `auth0`:

```hcl
connection "auth0" {
  plugin = "auth0"

  # `domain` (required) - Your Auth0 domain name.
  # This can also be set via the `AUTH0_DOMAIN` environment variable.
  # domain = "<your_auth0_domain>.<region>.auth0.com"

  # Either api_token or client_id + client_secret is also required

  # Get your API token from Auth0 https://auth0.com/docs/secure/tokens/access-tokens/management-api-access-tokens
  # This can also be set via the `AUTH0_API_TOKEN` environment variable.
  api_token = "fyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InZYbjJoN251dFRFbS1ZSndPSEdFdiJ9.eyJpc3MiOiJodHRwczovL2Rldi1zdGVhZHktbGFyay51cy5hdXRoMC5jb20vIiwic3ViIjoickNRWTF6UndJOEFOTzM4RkF3NE5nRFg2dzJIVHJOUWZAY1xp"

  # The below client_id and client_secret can be used instead of api_token. If both are specified, api_token will be used over client_id + client_secret.
  # This can also be set via the `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET` environment variables.
  # client_id = "rCQY1zRwI8ANO38FAw4NgDX6w2HTrNQh"
  # client_secret = "p8vxBHRRLiYDRNAQ9sk37sh2-6k_9XY25YgC2YY-mYcw715hvAl9olXg2Iqqpa2o"
}
```

### Credentials from Environment Variables

The Auth0 plugin will use the standard Auth0 environment variables to obtain credentials **only if other arguments (`domain`, `api_token`, `client_id`, `client_secret`) are not specified** in the connection:

#### API Token

```sh
export AUTH0_DOMAIN=<your_auth0_domain>.<region>.auth0.com
export AUTH0_API_TOKEN="fyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InZYbjJoN251dFRFbS1ZSndPSEdFdiJ9.eyJpc3MiOiJodHRwczovL2Rldi1zdGVhZHktbGFyay51cy5hdXRoMC5jb20vIiwic3ViIjoickNRWTF6UndJOEFOTzM4RkF3NE5nRFg2dzJIVHJOUWZAY2xp"
```

#### Client App

```sh
export AUTH0_DOMAIN=<your_auth0_domain>.<region>.auth0.com
export AUTH0_CLIENT_ID="rCQY1zRwI8ANO38FAw4NgDX6w2HTrNQf"
export AUTH0_CLIENT_SECRET="p8vxBHRRLiYDRNAQ9sk37sh2-6k_9XY25YgC2YY-mYcw715hvAl9olXg2Iqqpa7o"
```


