connection "auth0" {
  plugin = "auth0"

  # `domain` (required) - Your Auth0 domain name.
  # This can also be set via the `AUTH0_DOMAIN` environment variable.
  # domain = "<your_auth0_domain>.<region>.auth0.com"

  # Either api_token or client_id + client_secret is also required

  # Get your API token from Auth0 https://auth0.com/docs/secure/tokens/access-tokens/management-api-access-tokens
  # This can also be set via the `AUTH0_API_TOKEN` environment variable.
  api_token = "fyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InZYbjJoN251dFRFbS1ZSndPSEdFdiJ9.eyJpc3MiOiJodHRwczovL2Rldi1zdGVhZHktbGFyay51cy5hdXRoMC5jb20vIiwic3ViIjoickNRWTF6UndJOEFOTzM4RkF3NE5nRFg2dzJIVHJOUWZAY1xp"

  # The below client_id and client_secret can be used instead of api_token. If both are specified, api_token will be used over client_id + client_secret.
  # This can also be set via the `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET` environment variables.
  # client_id = "rCQY1zRwI8ANO38FAw4NgDX6w2HTrNQh"
  # client_secret = "p8vxBHRRLiYDRNAQ9sk37sh2-6k_9XY25YgC2YY-mYcw715hvAl9olXg2Iqqpa2o"
}
