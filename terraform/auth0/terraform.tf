terraform {
  cloud {
    organization = "DBSentry"

    workspaces {
      name = "Auth0Temporal"
    }
  }

  required_providers {
    auth0 = {
      source  = "auth0/auth0"
      version = "~> 1.1"
    }
  }

  required_version = "~> 1.2"
}

# These needs to come from secrets
provider "auth0" {
  domain        = "xxx.auth0.com"
  client_id     = "xxx"
  client_secret = "xxxxx"
}

