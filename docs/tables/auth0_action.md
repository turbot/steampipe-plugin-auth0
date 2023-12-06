---
title: "Steampipe Table: auth0_action - Query Auth0 Actions using SQL"
description: "Allows users to query Actions in Auth0, specifically the details of each action, providing insights into the actions' configurations and status."
---

# Table: auth0_action - Query Auth0 Actions using SQL

Auth0 Actions are code snippets that execute during the transaction lifecycle of Auth0's authentication and authorization process. They allow you to customize and extend Auth0's capabilities by inserting custom logic at specific points, known as extensibility points, in the transaction flow. With Auth0 Actions, you can modify user profiles, perform conditional MFA, enrich tokens, and more.

## Table Usage Guide

The `auth0_action` table provides insights into Actions within Auth0. As a security engineer, explore action-specific details through this table, including the action's script, dependencies, and associated triggers. Utilize it to uncover information about actions, such as their current status, the runtime they use, and the version of the action.

## Examples

### Deprecated NodeJS 12 based code deployed
Discover the segments that contain outdated NodeJS 12 based code that has been deployed. This can be beneficial in identifying areas that may require updates or migration to newer versions for improved security and performance.

```sql+postgres
select
  id,
  name,
  supported_triggers,
  updated_at
from
  auth0_action
where
  runtime = 'node12'
  and all_changes_deployed;
```

```sql+sqlite
select
  id,
  name,
  supported_triggers,
  updated_at
from
  auth0_action
where
  runtime = 'node12'
  and all_changes_deployed = 1;
```

### Actions triggered post a password change
Explore which actions are initiated following a password change. This can be helpful in understanding and managing security protocols.

```sql+postgres
select
  id,
  name,
  updated_at
from
  auth0_action
where
  supported_triggers -> 0 ->> 'id' = 'post-change-password';
```

```sql+sqlite
select
  id,
  name,
  updated_at
from
  auth0_action
where
  json_extract(supported_triggers, '$[0].id') = 'post-change-password';
```

### Action code by name
Analyze the settings to understand the specific code associated with an action, such as sending a notification. This can be useful in assessing the elements within your authentication process, particularly in identifying instances where specific actions are triggered.

```sql+postgres
select
  code
from
  auth0_action
where
  name = 'send-notification';
```

```sql+sqlite
select
  code
from
  auth0_action
where
  name = 'send-notification';
```

### Deployed actions
Discover the segments that have all their changes deployed in Auth0. This query is useful to understand which areas have the most recent updates, aiding in system management and maintenance.

```sql+postgres
select
  id,
  name,
  deployed_version ->> 'number' version,
  supported_triggers,
  updated_at
from
  auth0_action
where
  all_changes_deployed;
```

```sql+sqlite
select
  id,
  name,
  json_extract(deployed_version, '$.number') as version,
  supported_triggers,
  updated_at
from
  auth0_action
where
  all_changes_deployed;
```