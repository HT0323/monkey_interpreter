package object

import (
	"bytes"
	"fmt"
	"github.com/HT0323/monkey_interpreter/ast"
	"strings"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
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

// 真偽値を表現するオブジェクト
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// nullを表現するオブジェクト
type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

// 返値のオブジェクトを保持するオブジェクト
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Error struct {
	Message string
}

func (e *Error) Inspect() string  { return "ERROR:" + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// 関数を表現するオブジェクト
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// 文字列を表現するオブジェクト
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }

func (s *String) Inspect() string { return s.Value }
