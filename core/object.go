package core

type Object interface {
	GetType() SubjectType
	GetId() string
	GetName() string
	GetPolicyList()
}
