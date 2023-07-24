package core

type Policy struct {
	subject Subject
	object  Object
	actions []string
}
