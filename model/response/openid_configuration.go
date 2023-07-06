package response

type OpenidConfiguration struct {
	AuthorizationEndpoint                      string   `json:"authorization_endpoint"`
	DeviceAuthorizationEndpoint                string   `json:"device_authorization_endpoint"`
	ClaimsParameterSupported                   bool     `json:"claims_parameter_supported"`
	ClaimsSupported                            []string `json:"claims_supported"`
	CodeChallengeMethodsSupported              []string `json:"code_challenge_methods_supported"`
	EndSessionEndpoint                         string   `json:"end_session_endpoint"`
	GrantTypesSupported                        []string `json:"grant_types_supported"`
	Issuer                                     string   `json:"issuer"`
	JwksUri                                    string   `json:"jwks_uri"`
	AuthorizationResponseIssParameterSupported bool     `json:"authorization_response_iss_parameter_supported"`
	ResponseModesSupported                     []string `json:"response_modes_supported"`
	ResponseTypesSupported                     []string `json:"response_types_supported"`
	ScopesSupported                            []string `json:"scopes_supported"`
	SubjectTypesSupported                      []string `json:"subject_types_supported"`
	TokenEndpointAuthMethodsSupported          []string `json:"token_endpoint_auth_methods_supported"`
	TokenEndpointAuthSigningAlgValuesSupported []string `json:"token_endpoint_auth_signing_alg_values_supported"`
	TokenEndpoint                              string   `json:"token_endpoint"`
	IdTokenSigningAlgValuesSupported           []string `json:"id_token_signing_alg_values_supported"`
	PushedAuthorizationRequestEndpoint         string   `json:"pushed_authorization_request_endpoint"`
	RequestParameterSupported                  bool     `json:"request_parameter_supported"`
	RequestUriParameterSupported               bool     `json:"request_uri_parameter_supported"`
	UserinfoEndpoint                           string   `json:"userinfo_endpoint"`
	RevocationEndpoint                         string   `json:"revocation_endpoint"`
	ClaimTypesSupported                        []string `json:"claim_types_supported"`
}
