package request

type AuthorizationCodeAccessTokenRequest struct {
	GrantType   string `json:"grant_type" binding:"required"` // Value MUST be set to "authorization_code".
	Code        string `json:"code" binding:"required"`       // The authorization code received from the authorization server.
	RedirectUri string `json:"redirect_uri" binding:"required"`
	ClientId    string `json:"client_id" binding:"required"`
}

type ResourceOwnerAccessTokenRequest struct {
	GrantType string `json:"grant_type" binding:"required"` // Value MUST be set to "password".
	Username  string `json:"username" binding:"required"`   // The resource owner username.
	Password  string `json:"password" binding:"required"`   // The resource owner password.
	Scope     string `json:"scope" binding:"required"`      // The scope of the access request
}

type ClientCredentialsAccessTokenRequest struct {
	GrantType string `json:"grant_type" binding:"required"` // Value MUST be set to "client_credentials".
	Scope     string `json:"scope" binding:"required"`      // The scope of the access request
}
