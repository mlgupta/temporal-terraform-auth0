locals  {
}


# resource "auth0_tenant" "dbsentry-dev1" {
#     friendly_name = "DBSentry Dev1"
#     support_email = "support@dbsentry.com"
#     support_url = "https://dbsentry.com"
#     session_lifetime = 8760
#     sandbox_version = "12"
#     enabled_locales = ["en"]
#     default_redirection_uri = "https://dbsentry.com"
  
#     flags {
#         disable_clickjack_protection_headers   = true
#         enable_public_signup_user_exists_error = true
#         use_scope_descriptions_for_consent     = true
#         no_disclose_enterprise_connections     = false
#         disable_management_api_sms_obfuscation = false
#         disable_fields_map_fix                 = false
#     }

#     session_cookie {
#         mode = "non-persistent"
#     }

#     sessions {
#         oidc_logout_prompt_enabled = false
#     }
# }

resource "auth0_custom_domain" "my_custom_domain" {
  domain = "auth.dbsentry.com"
  type   = "auth0_managed_certs"
}