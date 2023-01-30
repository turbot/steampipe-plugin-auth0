![image](https://hub.steampipe.io/images/plugins/turbot/auth0-social-graphic.png)

# Auth0 Plugin for Steampipe

Use SQL to query users, clients, connections, keys and more from Auth0.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/auth0)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/auth0/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-auth0/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install auth0
```

Run a query:

```sql
select
  email,
  id,
  updated_at
from
  auth0_user
where
  not email_verified
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-auth0.git
cd steampipe-plugin-auth0
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/auth0.spc
```

Try it!

```
steampipe query
> .inspect auth0
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-auth0/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Auth0 Plugin](https://github.com/turbot/steampipe-plugin-auth0/labels/help%20wanted)
