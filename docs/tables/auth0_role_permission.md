---
title: "Steampipe Table: auth0_role_permission - Query Auth0 Role Permissions using SQL"
description: "Allows users to query Role Permissions in Auth0, specifically the permissions assigned to each role, providing insights into role-based access control and potential security risks."
---

# Table: auth0_role_permission - Query Auth0 Role Permissions using SQL

Auth0 Role Permission is a feature within Auth0's Role-Based Access Control (RBAC) that allows the assignment of permissions to roles. It provides a structured way to manage permissions for various roles, including users, groups, and applications. Auth0 Role Permission helps maintain the security and integrity of your resources by ensuring only authorized access.

## Table Usage Guide

The `auth0_role_permission` table provides insights into role permissions within Auth0's Role-Based Access Control (RBAC). As a security analyst, explore permission-specific details through this table, including the roles assigned, associated resources, and access levels. Utilize it to uncover information about permissions, such as those with broad access, the relationships between roles and permissions, and the verification of access controls.

## Examples

### List all the permissions assigned to a role
Discover the segments that have been granted specific permissions within a designated role. This is particularly useful in managing user access and ensuring appropriate security measures are in place.

```sql
select
  permission_name,
  description,
  resource_server_name
from
  auth0_role_permission
where
  role_id = 'rol_VkaG05dncCpNN3oI'
order by
  resource_server_name,
  permission_name;
```

### List roles with assigned permission of a resource server
Determine the areas in which specific roles are assigned permissions within a particular resource server. This can be useful for managing access control and ensuring appropriate permissions are allocated.

```sql
select
  r.name
from
  auth0_role r
  join
    auth0_role_permission p
    on p.role_id = r.id
where
  p.resource_server_name = 'novel-mutt'
group by
  r.name
order by
  r.name;
```