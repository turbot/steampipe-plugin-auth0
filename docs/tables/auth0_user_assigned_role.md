---
title: "Steampipe Table: auth0_user_assigned_role - Query Auth0 User Assigned Roles using SQL"
description: "Allows users to query User Assigned Roles in Auth0, specifically the roles assigned to each user, providing insights into user access and permissions."
---

# Table: auth0_user_assigned_role - Query Auth0 User Assigned Roles using SQL

Auth0 User Assigned Roles is a feature within Auth0 that allows you to manage and assign roles to users in your applications. It provides a centralized way to manage user permissions and access to various resources in your applications. Auth0 User Assigned Roles helps you ensure the right level of access is given to the right users, thereby enhancing the security and management of your applications.

## Table Usage Guide

The `auth0_user_assigned_role` table provides insights into user roles within Auth0. As a security analyst, explore user-specific role details through this table, including the roles assigned to each user, and their associated permissions. Utilize it to uncover information about user access, such as users with high-level permissions, and to verify the appropriateness of assigned roles.

## Examples

### Admin users without MFA
Explore which admin users have not enabled multi-factor authentication. This could be useful for identifying potential security vulnerabilities within your system.

```sql+postgres
select
  u.email,
  u.id,
  u.updated_at
from
  auth0_user u
  join
    auth0_user_assigned_role r
    on r.user_id = u.id
where
  r.name = 'admin'
  and u.multifactor is null;
```

```sql+sqlite
select
  u.email,
  u.id,
  u.updated_at
from
  auth0_user u
  join
    auth0_user_assigned_role r
    on r.user_id = u.id
where
  r.name = 'admin'
  and u.multifactor is null;
```

### Roles a user is assigned to
Uncover the details of different roles assigned to a specific user in the system, which is crucial for managing user permissions and access control. This can be used to ensure that users have the appropriate roles for their needs and to prevent unauthorized access to certain areas of the system.

```sql+postgres
select
  r.role_id,
  r.name,
  r.description
from
  auth0_user u
  join
    auth0_user_assigned_role r
    on r.user_id = u.id
where
  email = 'select-joey@mail.com';
```

```sql+sqlite
select
  r.role_id,
  r.name,
  r.description
from
  auth0_user u
  join
    auth0_user_assigned_role r
    on r.user_id = u.id
where
  email = 'select-joey@mail.com';
```