---
title: "Steampipe Table: auth0_user_permission - Query Auth0 User Permissions using SQL"
description: "Allows users to query User Permissions in Auth0, specifically the permissions granted to a user in an application, providing insights into user access control and potential security risks."
---

# Table: auth0_user_permission - Query Auth0 User Permissions using SQL

Auth0 User Permissions is a feature within the Auth0 Identity Platform that allows you to manage and assign permissions to users in your applications. It provides a centralized way to set up and manage user permissions, enhancing the security and access control of your applications. Auth0 User Permissions helps you stay informed about the access rights of your users and take appropriate actions when necessary.

## Table Usage Guide

The `auth0_user_permission` table provides insights into user permissions within Auth0 Identity Platform. As a security engineer, explore user-specific permissions through this table, including the applications they have access to, and the level of access they have. Utilize it to uncover information about users, such as those with high-level permissions, the access rights of individual users, and the verification of access controls.

## Examples

### All permissions granted by a user
Explore the range of permissions granted by a specific user to understand their level of access and control within the system. This can be useful for auditing user privileges and ensuring appropriate access levels.

```sql+postgres
select
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_user u
  join
    auth0_user_permission p
    on p.user_id = u.id
where
  email = 'select-joey@mail.com'
order by
  p.resource_server_name,
  p.permission_name;
```

```sql+sqlite
select
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_user u
  join
    auth0_user_permission p
    on p.user_id = u.id
where
  email = 'select-joey@mail.com'
order by
  p.resource_server_name,
  p.permission_name;
```

### User granted permissions in a resource server
Explore which permissions have been granted to a specific user within a particular resource server. This is useful for managing user access rights and ensuring appropriate permissions are assigned.

```sql+postgres
select
  u.email,
  p.permission_name,
  p.description
from
  auth0_user u
  join
    auth0_user_permission p
    on p.user_id = u.id
where
  u.email = 'select-joey@mail.com'
  and p.resource_server_name = 'novel-mutt';
```

```sql+sqlite
select
  u.email,
  p.permission_name,
  p.description
from
  auth0_user u
  join
    auth0_user_permission p
    on p.user_id = u.id
where
  u.email = 'select-joey@mail.com'
  and p.resource_server_name = 'novel-mutt';
```