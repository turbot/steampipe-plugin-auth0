---
title: "Steampipe Table: auth0_organization - Query Auth0 Organizations using SQL"
description: "Allows users to query Auth0 Organizations, providing detailed information about each organization's settings, metadata, and enabled connections."
---

# Table: auth0_organization - Query Auth0 Organizations using SQL

An Auth0 Organization is a logical grouping of users and applications in Auth0. It represents a real-world organization, such as a company or a school, and allows for easier user management and application configuration within that organization. Organizations in Auth0 can have multiple applications, connections, and members, each with their own settings and permissions.

## Table Usage Guide

The `auth0_organization` table provides insights into organizations within Auth0. As a security engineer, explore organization-specific details through this table, including settings, metadata, and enabled connections. Utilize it to uncover information about organizations, such as the applications they use, the connections they have enabled, and the members who are part of the organization.

## Examples

### List of my organizations
Explore the organizations you're associated with, including their names, display names, and logos, to gain a better understanding of your involvement and connections. This can be particularly useful for managing and organizing various collaborations and partnerships.

```sql+postgres
select
  name,
  display_name,
  branding ->> 'logo_url' as logo_url,
  metadata
from
  auth0_organization
order by
  name;
```

```sql+sqlite
select
  name,
  display_name,
  json_extract(branding, '$.logo_url') as logo_url,
  metadata
from
  auth0_organization
order by
  name;
```

### Filter my organizations by metadata tags
Identify specific organizations based on their metadata tags to streamline management and cost tracking. This is useful in scenarios where organizations are categorized by unique identifiers for better financial or operational control.

```sql+postgres
select
  name,
  display_name,
  branding ->> 'logo_url' as logo_url,
  metadata
from
  auth0_organization
where
  metadata ->> 'cost_id' = 'e42345'
order by
  name;
```

```sql+sqlite
select
  name,
  display_name,
  json_extract(branding, '$.logo_url') as logo_url,
  metadata
from
  auth0_organization
where
  json_extract(metadata, '$.cost_id') = 'e42345'
order by
  name;
```