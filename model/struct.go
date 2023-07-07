package model

import (
	"diana/model/request"
	"diana/model/response"
	"diana/utils"
	"errors"
	"strings"
)

type Client struct {
	Id           int64    `json:"id"`
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

// VerifyScope verifies the request scopes is allowed.
func (c *Client) VerifyScope(req *request.AuthorizationRequest) error {
	if len(utils.Union(c.Scopes, strings.Split(req.Scope, " "))) > 0 {
		return nil
	}
	return errors.New(response.InvalidScope.String())
}

func (c *Client) VerifyGrantType(req *request.AuthorizationRequest) error {
	if utils.Contains(c.GrantTypes, req.ResponseType) {
		return nil
	}
	return errors.New(response.UnsupportedResponseType.String())
}

func (c *Client) VerifySecret(secret string) error {
	if c.ClientSecret == secret {
		return nil
	}
	return errors.New(response.UnauthorizedClient.String())
}

func (c *Client) VerifyRedirectUri(req *request.AuthorizationRequest) error {
	if c.RedirectUri != req.RedirectUri {
		return nil
	}
	return errors.New(response.InvalidRequest.String())
}

func (c *Client) codeVerify(req *request.AuthorizationRequest) error {
	funcs := []func(req *request.AuthorizationRequest) error{
		c.VerifyGrantType,
		c.VerifyScope,
	}
	for _, fn := range funcs {
		if err := fn(req); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Verify(req *request.AuthorizationRequest) error {
	switch req.ResponseType {
	case request.CodeGrant.String():
		return c.codeVerify(req)
	}
	return nil
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
