package request

type AuthorizationCodeGrantRequest struct {
	ResponseType string `json:"response_type" form:"response_type" binding:"required"` // Value MUST be set to "code".
	ClientId     string `json:"client_id" form:"client_id" binding:"required"`         // The client identifier as described in RFC6749 Section 2.2
	RedirectUri  string `json:"redirect_uri" form:"redirect_uri"`                      // As described in RFC6749 Section 3.1.2
	Scope        string `json:"scope" form:"scope"`                                    // The scope of the access request as described by RFC6749 Section 3.3
	State        string `json:"state" form:"state"`                                    // An opaque value used by the client to maintain state between the request and callback.
}

type ImplicitGrantRequest struct {
	ResponseType string `json:"response_type" binding:"required"` // Value MUST be set to "code".
	ClientId     string `json:"client_id" binding:"required"`     // The client identifier as described in RFC6749 Section 2.2
	RedirectUri  string `json:"redirect_uri"`                     // As described in RFC6749 Section 3.1.2
	Scope        string `json:"scope"`                            // The scope of the access request as described by RFC6749 Section 3.3
	State        string `json:"state"`                            // An opaque value used by the client to maintain state between the request and callback.
}
