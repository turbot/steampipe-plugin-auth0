---
title: "Steampipe Table: auth0_hook - Query Auth0 Hooks using SQL"
description: "Allows users to query Auth0 Hooks, specifically providing insights into the custom scripts that are executed as part of the Auth0 pipeline."
---

# Table: auth0_hook - Query Auth0 Hooks using SQL

Auth0 Hooks are customizable scripts executed as part of the Auth0 pipeline, allowing you to customize the behavior of Auth0. They are written in Node.js and can be used for a variety of purposes, such as enriching user profiles, denying access based on custom rules, or even integrating with other services. Hooks are a powerful tool for extending the functionality of Auth0.

## Table Usage Guide

The `auth0_hook` table provides insights into Hooks within Auth0. As a DevOps engineer or security analyst, explore hook-specific details through this table, including trigger, script, and associated metadata. Utilize it to uncover information about hooks, such as those with specific triggers, the content of the script, and the status of the hook.

## Examples

### Enabled hook scripts
Analyze the settings to understand the active hook scripts in your Auth0 environment. This can help in managing and troubleshooting your authentication workflows.

```sql
select
  id,
  name,
  dependencies,
  trigger_id,
  script
from
  auth0_hook
where
  enabled;
```

### Post change password script
Review the configuration for post-password-change procedures to assess the elements within your authentication system. This allows you to pinpoint the specific locations where changes have been made, enabling better security management.

```sql
select
  id,
  name,
  script,
  dependencies,
  enabled
from
  auth0_hook
where
  trigger_id = 'post-change-password';
```

### User registration scripts
Analyze the settings to understand the status and dependencies of user registration scripts in Auth0. This could be beneficial in managing and optimizing the user registration process.

```sql
select
  id,
  name,
  script,
  dependencies,
  enabled
from
  auth0_hook
where
  trigger_id in ('pre-user-registration', 'post-user-registration');
```