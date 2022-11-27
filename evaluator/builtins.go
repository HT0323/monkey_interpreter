package evaluator

import "github.com/HT0323/monkey_interpreter/object"

//組み込み関数名とそれに対応する処理を保持
var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		func(args ...object.Object) object.Object {
			//len関数の引数は一つだけなのでそれ以外の場合はエラー
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				//有効な引数はString型のみ
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
			return NULL
		},
	},
}
