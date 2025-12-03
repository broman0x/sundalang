package sundalang

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// OBJECT SYSTEM
type ObjectType string
const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	STRING_OBJ  = "STRING"
	NULL_OBJ    = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ   = "ERROR"
	FUNCTION_OBJ = "FUNCTION"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct { Value int64 }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type BooleanObject struct { Value bool }
func (b *BooleanObject) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *BooleanObject) Type() ObjectType { return BOOLEAN_OBJ }

type String struct { Value string }
func (s *String) Inspect() string { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }

type Null struct{}
func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

type ReturnValue struct { Value Object }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Error struct { Message string }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

type Function struct {
	Parameters []*Identifier
	Body       *BlockStatement
	Env        *Environment
}
func (f *Function) Inspect() string { return "fungsi(...)" }
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// ENVIRONMENT
type Environment struct {
	store map[string]Object
	outer *Environment
}
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
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

// EVALUATOR
var (
	TRUE  = &BooleanObject{Value: true}
	FALSE = &BooleanObject{Value: false}
	NULL  = &Null{}
)

func Eval(node Node, env *Environment) Object {
	switch node := node.(type) {
	case *Program: return evalProgram(node, env)
	case *BlockStatement: return evalBlockStatement(node, env)
	case *VarStatement:
		val := Eval(node.Value, env)
		if isError(val) { return val }
		env.Set(node.Name.Value, val)
	case *ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) { return val }
		return &ReturnValue{Value: val}
	case *Identifier: return evalIdentifier(node, env)
	case *IntegerLiteral: return &Integer{Value: node.Value}
	case *Boolean: return nativeBoolToBooleanObject(node.Value)
	case *StringLiteral: return &String{Value: node.Value}
	case *FunctionLiteral:
		return &Function{Parameters: node.Parameters, Env: env, Body: node.Body}
	case *CallExpression:
		function := Eval(node.Function, env)
		if isError(function) { return function }
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) { return args[0] }
		return applyFunction(function, args)
	case *PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) { return right }
		return evalPrefixExpression(node.Operator, right)
	case *InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) { return left }
		right := Eval(node.Right, env)
		if isError(right) { return right }
		return evalInfixExpression(node.Operator, left, right)
	case *IfExpression: return evalIfExpression(node, env)
	case *WhileStatement: return evalWhileStatement(node, env)
	case *PrintStatement:
		val := Eval(node.Expression, env)
		if !isError(val) { fmt.Println(val.Inspect()) }
		return NULL
	case *InputExpression:
		if node.Prompt != "" { fmt.Print(node.Prompt) }
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if i, err := strconv.ParseInt(text, 10, 64); err == nil {
			return &Integer{Value: i}
		}
		return &String{Value: text}
	}
	return nil
}

func evalProgram(program *Program, env *Environment) Object {
	var result Object
	for _, statement := range program.Statements {
		result = Eval(statement, env)
		if result != nil {
			if result.Type() == RETURN_VALUE_OBJ {
				return result.(*ReturnValue).Value
			}
			if result.Type() == ERROR_OBJ { return result }
		}
	}
	return result
}

func evalBlockStatement(block *BlockStatement, env *Environment) Object {
	var result Object
	for _, statement := range block.Statements {
		result = Eval(statement, env)
		if result != nil {
			if result.Type() == RETURN_VALUE_OBJ || result.Type() == ERROR_OBJ {
				return result
			}
		}
	}
	return result
}

func evalExpressions(exps []Expression, env *Environment) []Object {
	var result []Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) { return []Object{evaluated} }
		result = append(result, evaluated)
	}
	return result
}

func applyFunction(fn Object, args []Object) Object {
	function, ok := fn.(*Function)
	if !ok { return &Error{Message: "not a function: " + string(fn.Type()) } }
	
	extendedEnv := NewEnclosedEnvironment(function.Env)
	for i, param := range function.Parameters {
		extendedEnv.Set(param.Value, args[i])
	}
	evaluated := Eval(function.Body, extendedEnv)
	if returnValue, ok := evaluated.(*ReturnValue); ok {
		return returnValue.Value
	}
	return evaluated
}

func evalIdentifier(node *Identifier, env *Environment) Object {
	val, ok := env.Get(node.Value)
	if !ok { return &Error{Message: "identifier not found: " + node.Value} }
	return val
}

func nativeBoolToBooleanObject(input bool) *BooleanObject {
	if input { return TRUE }
	return FALSE
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!": return evalBangOperatorExpression(right)
	case "-": return evalMinusPrefixOperatorExpression(right)
	default: return &Error{Message: fmt.Sprintf("unknown operator: %s%s", operator, right.Type())}
	}
}

func evalBangOperatorExpression(right Object) Object {
	switch right {
	case TRUE: return FALSE
	case FALSE: return TRUE
	case NULL: return TRUE
	default: return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right Object) Object {
	if right.Type() != INTEGER_OBJ { return &Error{Message: fmt.Sprintf("unknown operator: -%s", right.Type())} }
	return &Integer{Value: -right.(*Integer).Value}
}

func evalInfixExpression(operator string, left, right Object) Object {
	if left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left.(*Integer), right.(*Integer))
	}
	if operator == "==" {
		return nativeBoolToBooleanObject(left == right)
	}
	if operator == "!=" {
		return nativeBoolToBooleanObject(left != right)
	}
	if left.Type() == STRING_OBJ && right.Type() == STRING_OBJ && operator == "+" {
		return &String{Value: left.(*String).Value + right.(*String).Value}
	}
	return &Error{Message: fmt.Sprintf("type mismatch: %s %s %s", left.Type(), operator, right.Type())}
}

func evalIntegerInfixExpression(operator string, left, right *Integer) Object {
	lVal := left.Value
	rVal := right.Value
	switch operator {
	case "+": return &Integer{Value: lVal + rVal}
	case "-": return &Integer{Value: lVal - rVal}
	case "*": return &Integer{Value: lVal * rVal}
	case "/": return &Integer{Value: lVal / rVal}
	case "%": return &Integer{Value: lVal % rVal} 
	case "<": return nativeBoolToBooleanObject(lVal < rVal)
	case ">": return nativeBoolToBooleanObject(lVal > rVal)
	case "==": return nativeBoolToBooleanObject(lVal == rVal)
	case "!=": return nativeBoolToBooleanObject(lVal != rVal)
	default: return &Error{Message: "unknown operator for ints"}
	}
}

func evalIfExpression(ie *IfExpression, env *Environment) Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) { return condition }
	if isTruthy(condition) { return Eval(ie.Consequence, env) }
	if ie.Alternative != nil { return Eval(ie.Alternative, env) }
	return NULL
}

func evalWhileStatement(ws *WhileStatement, env *Environment) Object {
	for {
		condition := Eval(ws.Condition, env)
		if isError(condition) { return condition }
		if !isTruthy(condition) { break }
		Eval(ws.Body, env)
	}
	return NULL
}

func isTruthy(obj Object) bool {
	if obj == NULL { return false }
	if obj == TRUE { return true }
	if obj == FALSE { return false }
	return true
}

func isError(obj Object) bool {
	if obj != nil { return obj.Type() == ERROR_OBJ }
	return false
}