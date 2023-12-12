---
title: "Steampipe Table: auth0_role - Query Auth0 Roles using SQL"
description: "Allows users to query Roles in Auth0, specifically the role details including permissions and associated users, providing insights into role management and access control."
---

# Table: auth0_role - Query Auth0 Roles using SQL

Auth0 is a flexible, drop-in solution to add authentication and authorization services to your applications. It provides a universal authentication & authorization platform for web, mobile and legacy applications, and allows you to authenticate and authorize apps and APIs with any identity provider running on any stack, any device or cloud. With Auth0, you can manage roles and permissions, assign users to roles, and control who can access your APIs.

## Table Usage Guide

The `auth0_role` table provides insights into roles within Auth0. As a Security engineer, explore role-specific details through this table, including permissions, associated users, and other metadata. Utilize it to uncover information about roles, such as those with specific permissions, the users associated with each role, and the overall management of access control.

## Examples

### Non-admin roles with 'all:*' permissions
Determine the roles, excluding the 'admin', that have been granted all permissions. This allows for a review of potential security vulnerabilities by identifying roles with overly broad access rights.

```sql+postgres
select
  r.name as role_name,
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_role r
  join
    auth0_role_permission p
    on p.role_id = r.id
where
  r.name <> 'admin'
  and p.permission_name like 'all:%';
```

```sql+sqlite
select
  r.name as role_name,
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_role r
  join
    auth0_role_permission p
    on p.role_id = r.id
where
  r.name <> 'admin'
  and p.permission_name like 'all:%';
```

### List all permissions assigned to an admin role
Explore which permissions are associated with an administrative role to better understand the access rights and potential security implications. This can be useful when auditing system access or planning role changes.

```sql+postgres
select
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_role r
  join
    auth0_role_permission p
    on p.role_id = r.id
where
  r.name = 'admin';
```

```sql+sqlite
select
  p.permission_name,
  p.description,
  p.resource_server_name
from
  auth0_role r
  join
    auth0_role_permission p
    on p.role_id = r.id
where
  r.name = 'admin';
```