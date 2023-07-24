package app

import (
	"diana/core"
)

type GrantType int32
type CredentialType int32

const (
	AuthorizationCode GrantType = 0
)

const (
	Email CredentialType = iota
	Password
	Phone
	LDAP
)

type App struct {
	Id                   string
	AppId                string
	Name                 string
	Secret               string
	Description          string
	AuthenticationConfig AuthenticationConfig
	AuthorizationConfig  AuthorizationConfig

	addedSubject   []core.Subject
	removedSubject []core.Subject

	addedObjects   []core.Object
	removedObjects []core.Object

	Namespace string
}

func NewApp(appName string, description string) (*App, error) {
	return &App{Name: appName, Description: description}, nil
}

func (app *App) AddSubject(subject core.Subject) {
	app.addedSubject = append(app.addedSubject, subject)
}

func (app *App) GetAddedSubject() []core.Subject {
	return app.addedSubject
}

func (app *App) RemoveSubject(subject core.Subject) {
	app.removedSubject = append(app.removedSubject, subject)
}

func (app *App) GetRemovedSubject() []core.Subject {
	return app.removedSubject
}

func (app *App) AddObject(obj core.Object) {
	app.addedObjects = append(app.addedObjects, obj)
}

func (app *App) GetAddedObjects() []core.Object {
	return app.addedObjects
}

func (app *App) GetRemovedObjects() []core.Object {
	return app.removedObjects
}

func (app *App) RemoveObject(obj core.Object) {
	app.removedObjects = append(app.removedObjects, obj)
}

type AuthenticationConfig struct {
	identitySources []IdentitySource
	credentials     []CredentialType
	callbackUri     string
}

// AuthorizationConfig effects OAuth 2.0 authorization
type AuthorizationConfig struct {
	grantTypes                  []GrantType
	scope                       []string
	authorizationCodeExpiration int32
	idTokenExpiration           int32
	issuer                      string
	serviceDiscoveryAddress     string
}

type IdentitySource interface {
	GetAppId() string
	GetAppSecret() string
	GetName() string
	GetLogo() string
}
