package object

type Environment struct {
	store map[string]Object
	outer *Environment //  新たにEnvironment構造体を作成するとる時に既存の構造体を格納する
}

func NewEnviroment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

//既存のEnvironment構造体を引き継いで新たに構造体を作成
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnviroment()
	env.outer = outer
	return env
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
