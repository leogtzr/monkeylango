package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

/*
	This new behaviour mirrors how we think about variable scopes. There is an inner scope and
	an outer scope. If something is not found in the inner scope, it’s looked up in the outer scope.
	The outer scope encloses the inner scope. And the inner scope extends the outer one.
*/
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}
