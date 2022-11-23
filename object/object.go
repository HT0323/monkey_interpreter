package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// 整数型を表現するオブジェクト
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
