package utils

type Client struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectUri  string   `json:"redirect_uri"`
	ClientName   string   `json:"client_name"`
	Description  string   `json:"description"`
	CreatedTime  string   `json:"created_time"`
	UpdatedTime  string   `json:"updated_time"`
	GrantTypes   []string `json:"grant_types"`
	ClientUri    string   `json:"client_uri"`
	Active       bool     `json:"active"`
	Scopes       []string `json:"scopes"`
}

type Account struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	CreateTime string `json:"create_time"`
}

type AuthCodeInfo struct {
	ClientId    string `json:"client_id"`
	ExpiresTime int64  `json:"expires_time"`
	CreateTime  string `json:"create_time"`
}

type TokenInfo struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresTime  int64  `json:"expires_time"`
	CreateTime   string `json:"create_time"`
}
