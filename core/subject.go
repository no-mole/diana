package core

type SubjectType int32

const (
	User SubjectType = iota
	Group
	Department
	Role
)

type Subject interface {
	GetType() SubjectType
	GetId() string
	GetName() string
	GetPolicy()

	SetType(subjectType SubjectType)
	SetId(id string)
	SetName() string
	SetPolicy()
}
