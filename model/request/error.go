package request

type OAuthError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorUri         string `json:"error_uri"`
	State            string `json:"state"`
}

type ErrorType int32

const (
	InvalidRequest          ErrorType = 0
	UnauthorizedClient      ErrorType = 1
	AccessDenied            ErrorType = 2
	UnsupportedResponseType ErrorType = 3
	InvalidScope            ErrorType = 4
	ServerError             ErrorType = 5
	TemporarilyUnavailable  ErrorType = 6
)

func (e ErrorType) String() string {
	switch e {
	case InvalidRequest:
		return "invalid_request"
	case UnauthorizedClient:
		return "unauthorized_client"
	case AccessDenied:
		return "access_denied"
	case UnsupportedResponseType:
		return "unsupported_response_type"
	case InvalidScope:
		return "invalid_scope"
	case ServerError:
		return "server_error"
	case TemporarilyUnavailable:
		return "temporarily_unavailable"
	default:
		return "unknown"
	}
}

func (e ErrorType) Message() string {
	switch e {
	case InvalidRequest:
		return "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed."
	case UnauthorizedClient:
		return "The client is not authorized to request an access token using this method."
	case AccessDenied:
		return "The resource owner or authorization server denied the request."
	case UnsupportedResponseType:
		return "The authorization server does not support obtaining an access token using this method."
	case InvalidScope:
		return "The requested scope is invalid, unknown, or malformed."
	case ServerError:
		return "The authorization server encountered an unexpected condition that prevented it from fulfilling the request. (This error code is needed because a 500 Internal Server Error HTTP status code cannot be returned to the client via an HTTP redirect.)"
	case TemporarilyUnavailable:
		return "The authorization server is currently unable to handle the request due to a temporary overloading or maintenance of the server.  (This error code is needed because a 503 Service Unavailable HTTP status code cannot be returned to the client via an HTTP redirect.)"
	default:
		return "unknown"
	}
}

func NewOAuthError(errorType ErrorType, uri string, state string) OAuthError {
	return OAuthError{
		Error:            errorType.String(),
		ErrorUri:         uri,
		State:            state,
		ErrorDescription: errorType.Message(),
	}
}
