---
title: "Steampipe Table: auth0_role_assigned_user - Query Auth0 Assigned Users using SQL"
description: "Allows users to query Assigned Users in Auth0, specifically the users assigned to different roles, providing insights into user access and role permissions."
---

# Table: auth0_role_assigned_user - Query Auth0 Assigned Users using SQL

Auth0 is a flexible, drop-in solution to add authentication and authorization services to your applications. It offers a platform to authenticate, authorize, and secure access for applications, devices, and users. The service simplifies identity and access management for developers while making it safer and faster for users to log into the applications.

## Table Usage Guide

The `auth0_role_assigned_user` table provides insights into users assigned to roles within Auth0 Identity and Access Management. As a security analyst, explore user-specific details through this table, including user roles, permissions, and associated metadata. Utilize it to uncover information about users, such as those with specific role permissions, the relationships between users and roles, and the verification of user access.

## Examples

### List users assigned to a role
Explore which users have been assigned the 'operator' role in your Auth0 system. This can be useful to maintain security and manage user permissions effectively.

```sql+postgres
select
  u.name,
  u.email,
  u.user_id
from
  auth0_role r
  join
    auth0_role_assigned_user u
    on u.role_id = r.id
where
  r.name = 'operator'
order by
  u.name;
```

```sql+sqlite
select
  u.name,
  u.email,
  u.user_id
from
  auth0_role r
  join
    auth0_role_assigned_user u
    on u.role_id = r.id
where
  r.name = 'operator'
order by
  u.name;
```

### Admin users with unverified email
Explore which administrative users have not yet verified their emails. This is useful in ensuring all admins have completed necessary verification steps for security purposes.

```sql+postgres
select
  u.email,
  u.id,
  u.updated_at
from
  auth0_role r
  join
    auth0_role_assigned_user ru
    on ru.role_id = r.id
  join
    auth0_user u
    on u.id = ru.user_id
where
  r.name = 'admin'
  and not u.email_verified;
```

```sql+sqlite
select
  u.email,
  u.id,
  u.updated_at
from
  auth0_role r
  join
    auth0_role_assigned_user ru
    on ru.role_id = r.id
  join
    auth0_user u
    on u.id = ru.user_id
where
  r.name = 'admin'
  and u.email_verified = 0;
```