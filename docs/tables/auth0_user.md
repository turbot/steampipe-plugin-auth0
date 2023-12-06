---
title: "Steampipe Table: auth0_user - Query Auth0 Users using SQL"
description: "Allows users to query Auth0 Users, specifically enabling retrieval of user profile data, including identifiers, names, picture URLs, and user metadata."
---

# Table: auth0_user - Query Auth0 Users using SQL

Auth0 is a flexible, drop-in solution to add authentication and authorization services to your applications. It provides a universal authentication & authorization platform for web, mobile and legacy applications, and makes it easy to implement even the most complex identity solutions for your applications. Auth0 allows you to authenticate and authorize applications and APIs with any identity provider running on any stack, device, or cloud.

## Table Usage Guide

The `auth0_user` table provides insights into user profiles within Auth0. As a security engineer, explore user-specific details through this table, including identifiers, names, picture URLs, and user metadata. Utilize it to uncover information about users, such as their identifiers, metadata, and other profile details.

## Examples

### Users without MFA
Identify users who haven't activated multi-factor authentication. This is useful for enhancing security measures by pinpointing potential vulnerabilities.

```sql+postgres
select
  email,
  id,
  updated_at
from
  auth0_user
where
  multifactor is null;
```

```sql+sqlite
select
  email,
  id,
  updated_at
from
  auth0_user
where
  multifactor is null;
```

### Users with unverified email
Analyze user data to identify accounts with unverified email addresses. This can be used to pinpoint potential security risks or areas for user outreach.

```sql+postgres
select
  email,
  id,
  updated_at
from
  auth0_user
where
  not email_verified;
```

```sql+sqlite
select
  email,
  id,
  updated_at
from
  auth0_user
where
  email_verified = 0;
```

### Ranking of highly used auth0 connections
Explore the frequency of different Auth0 connections to understand the most commonly used ones. This can help in identifying popular connection methods, aiding in strategic decision-making for resource allocation or optimization efforts.

```sql+postgres
select
  i ->> 'connection' as "connection",
  count(1)
from
  auth0_user u,
  jsonb_array_elements(u.identities) i
group by
  i ->> 'connection'
order by
  count desc;
```

```sql+sqlite
select
  json_extract(i.value, '$.connection') as "connection",
  count(1)
from
  auth0_user u,
  json_each(u.identities) i
group by
  json_extract(i.value, '$.connection')
order by
  count(1) desc;
```

### Users signed up through GitHub
Explore which users have signed up through GitHub to gain insights into the user base and their login activity. This can help you understand the popularity of different signup methods and identify trends in user behavior.

```sql+postgres
select
  nickname,
  id,
  last_login
from
  auth0_user
where
  identities -> 0 ->> 'connection' = 'github';
```

```sql+sqlite
select
  nickname,
  id,
  last_login
from
  auth0_user
where
  json_extract(json_extract(identities, '$[0]'), '$.connection') = 'github';
```