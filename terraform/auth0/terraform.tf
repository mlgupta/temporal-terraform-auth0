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
  domain        = "dev-x065zf5k0esil2fl.us.auth0.com"
  client_id     = "5t4sCYmnwZ7znshYjquXSy58vo1sdDhS"
  client_secret = "sL9yjqRkataAUlzafQWs6kIWZO2LMlpSBbub_PuPnTv5I1buHYRPPCpEbDKf0WLO"
}

