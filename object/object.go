package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/leogtzr/monkeylango/ast"
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
	BUILTIN_OBJ      = "BUILTIN"
)

// Error ...
type Error struct {
	Message string
}

func (e *Error) Type() string {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

// Object ...
type Object interface {
	Type() string
	Inspect() string
}

// Integer ...
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() string {
	return INTEGER_OBJ
}

// Boolean ...
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() string {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null ...
type Null struct{}

func (n *Null) Type() string {
	return NULL_OBJ
}
func (n *Null) Inspect() string {
	return "null"
}

// ReturnValue ...
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() string {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Function ...
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() string {
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

type String struct {
	Value string
}

func (s *String) Type() string {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

// BuiltinFunction ...
type BuiltinFunction func(args ...Object) Object

// Builtin ...
type Builtin struct {
	Fn BuiltinFunction
	// Fn func(args ...Object) Object
}

func (b *Builtin) Type() string {
	return BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}
