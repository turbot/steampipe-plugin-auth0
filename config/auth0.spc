connection "auth0" {
  plugin = "auth0"

  # Get your API token from Auth0 https://auth0.com/docs/secure/tokens/access-tokens/management-api-access-tokens

  domain = "<your_auth0_domain>.<region>.auth0.com"
  api_token = "fyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InZYbjJoN251dFRFbS1ZSndPSEdFdiJ9.eyJpc3MiOiJodHRwczovL2Rldi1zdGVhZHktbGFyay51cy5hdXRoMC5jb20vIiwic3ViIjoickNRWTF6UndJOEFOTzM4RkF3NE5nRFg2dzJIVHJOUWZAY2xp"

  # The bellow client_id and client_secret can be used instead of api_token. If both are specified, api_token will be used over client_id + client_secret.
  # client_id = "rCQY1zRwI8ANO38FAw4NgDX6w2HTrNQf"
  # client_secret = "p8vxBHRRLiYDRNAQ9sk37sh2-6k_9XY25YgC2YY-mYcw715hvAl9olXg2Iqqpa7o"
}
