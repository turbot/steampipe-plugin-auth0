---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/auth0.svg"
brand_color: "#000000"
display_name: "Auth0"
name: "auth0"
description: "Use SQL to query users, clients, connections, keys and more from Auth0."
og_description: "Query Auth0 with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/auth0-social-graphic.png"
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
  not email_verified
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
