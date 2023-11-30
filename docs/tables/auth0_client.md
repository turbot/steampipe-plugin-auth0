---
title: "Steampipe Table: auth0_client - Query Auth0 Clients using SQL"
description: "Allows users to query Auth0 Clients, specifically the details of applications, APIs, and third-party services, providing insights into the configuration and management of these resources."
---

# Table: auth0_client - Query Auth0 Clients using SQL

Auth0 is a flexible, drop-in solution to add authentication and authorization services to your applications. It provides a platform to authenticate, authorize, and secure access for applications, devices, and users. With Auth0, you can manage authentication of users and enable the integration of social identity providers.

## Table Usage Guide

The `auth0_client` table provides insights into the clients within Auth0. As a DevOps engineer or a security analyst, explore client-specific details through this table, including client types, grant types, and associated metadata. Utilize it to uncover information about clients, such as their callback URLs, allowed origins, and client secrets, aiding in the configuration and management of these resources.

## Examples

### Number of clients by type
Determine the distribution of client types within your application ecosystem. This can provide insights into the variety and prevalence of different client types, aiding in strategic decision-making and resource allocation.

```sql
select
  app_type,
  count(1)
from
  auth0_client
group by
  app_type;
```

### Token lifetime
Assess the elements within your Auth0 clients to understand the lifespan of their tokens. This can be useful to manage session durations and enhance security by determining the idle and active lifetimes of tokens.

```sql
select
  client_id,
  name,
  refresh_token ->> 'token_lifetime' as token_lifetime,
  refresh_token ->> 'idle_token_lifetime' as idle_token_lifetime
from
  auth0_client
order by
  name;
```

### Grant types of a client
Analyze the types of authorizations granted to a specific client in the Auth0 platform. This can be useful for assessing security settings and understanding the level of access a client has.

```sql
select
  g as grant_types
from
  auth0_client c,
  jsonb_array_elements(grant_types) g
where
  client_id = 'Jh5ap2mN94TJmZZ1sVeVmtW9Fpaim190';
```