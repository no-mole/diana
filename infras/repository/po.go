package repository

import (
	"diana/core"
	"gorm.io/gorm"
)

type Meta struct {
	gorm.Model
	creator string
}

type App struct {
	id          int64
	appId       string
	appName     string
	appSecret   string
	description string
	Meta
}

// AppSubjectRef defines which subject can sign in app.
type AppSubjectRef struct {
	id          int64
	appId       string
	subjectId   string
	subjectType int32
	Meta
}

func (ref *AppSubjectRef) ConvertTo() (appId string, subject core.Subject) {
	appId = ref.appId
	subject.SetId(ref.subjectId)
	subject.SetType(core.SubjectType(ref.subjectType))
	return
}

func (ref *AppSubjectRef) ConvertFrom(appId string, subject core.Subject) *AppSubjectRef {
	return &AppSubjectRef{
		subjectId:   subject.GetId(),
		subjectType: int32(subject.GetType()),
		appId:       appId,
	}
}

// AppObjectRef defines which object is belong to app.
type AppObjectRef struct {
	id         int64
	appId      string
	objectId   string
	objectType int32
	Meta
}

type AuthenticationConfig struct {
	id                int64
	appId             string
	providerAppId     string
	providerAppSecret string
	Meta
}

type AuthorizationConfig struct {
	id        int64
	appId     string
	grantType string
	Meta
}
