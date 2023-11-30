---
title: "Steampipe Table: auth0_signing_key - Query Auth0 Signing Keys using SQL"
description: "Allows users to query Auth0 Signing Keys, providing insights into the keys used for verifying the signature of incoming JWT tokens."
---

# Table: auth0_signing_key - Query Auth0 Signing Keys using SQL

Auth0 Signing Keys are crucial components in the Auth0 security model, used for verifying the signature of incoming JWT tokens. These keys are primarily used in the process of authentication and authorization in applications. They are part of the wider Auth0 platform, a flexible and scalable solution for identity and access management.

## Table Usage Guide

The `auth0_signing_key` table offers valuable insights into the signing keys within the Auth0 platform. If you're a security analyst or a developer, you can use this table to explore key-specific details, including the key ID, certificate, and associated metadata. This can be particularly useful for verifying the integrity of JWT tokens, ensuring secure user authentication, and maintaining the overall security posture of your applications.

## Examples

### For how long has current signing key been available
Explore the duration for which the current signing key has been active. This can help in identifying potential security risks and maintaining good practices by regularly updating keys.

```sql
select
  current_date - current_since as current_for
from
  auth0_signing_key
where
  current;
```

### Next signing key
Determine the upcoming signing key in your Auth0 environment to ensure smooth transitions of authentication processes and avoid unexpected service disruptions.

```sql
select
  kid,
  fingerprint,
  thumbprint
from
  auth0_signing_key
where
  next;
```

### Previous signing key
Explore the history of signing keys to understand when a particular key was in use. This can be beneficial for auditing purposes or to trace back any security-related issues.

```sql
select
  kid,
  fingerprint,
  thumbprint,
  current_since,
  current_until
from
  auth0_signing_key
where
  previous;
```

### Average time for which the previous signing keys were available
Determine the average duration for which previous authentication keys were available. This is useful for understanding the typical lifespan of keys, aiding in planning for key rotation schedules.

```sql
select
  avg(current_until - current_since) as average_duration
from
  auth0_signing_key
where
  previous;
```

### Revoked signing keys
Assess the elements within your Auth0 system to identify and prioritize revoked signing keys. This allows you to maintain system integrity by focusing on keys that have been revoked, especially useful in high-security environments.

```sql
select
  kid,
  fingerprint,
  thumbprint,
  revoked_at
from
  auth0_signing_key
where
  revoked
order by
  revoked_at desc;
```