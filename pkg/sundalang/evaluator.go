package sundalang

import (
	"bufio"
	"bytes"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
	"time"
)

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	STRING_OBJ       = "STRING"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
	BUILTIN_OBJ      = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Integer struct{ Value int64 }

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

type BooleanObject struct{ Value bool }

func (b *BooleanObject) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *BooleanObject) Type() ObjectType { return BOOLEAN_OBJ }
func (b *BooleanObject) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

type String struct{ Value string }

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

type ReturnValue struct{ Value Object }

func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Error struct{ Message string }

func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
func (e *Error) Type() ObjectType { return ERROR_OBJ }

type Function struct {
	Parameters []*Identifier
	Body       *BlockStatement
	Env        *Environment
}

func (f *Function) Inspect() string  { return "fungsi(...)" }
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type HashPair struct {
	Key   Object
	Value Object
}
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "fungsi bawaan" }

var builtins = map[string]*Builtin{
	"panjang": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: fmt.Sprintf("panjang() butuh 1 argumen, dibere %d", len(args))}
			}
			switch arg := args[0].(type) {
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			default:
				return &Error{Message: fmt.Sprintf("argumen ka panjang() teu didukung, tipe: %s", args[0].Type())}
			}
		},
	},
	"mimiti": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "mimiti() butuh 1 argumen array"}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "argumen mimiti() kudu ARRAY"}
			}
			arr := args[0].(*Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}
			return NULL
		},
	},
	"tungtung": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "tungtung() butuh 1 argumen array"}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "argumen tungtung() kudu ARRAY"}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}
			return NULL
		},
	},
	"asupkeun": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return &Error{Message: "asupkeun() butuh 2 argumen (array, isi)"}
			}
			if args[0].Type() != ARRAY_OBJ {
				return &Error{Message: "argumen kahiji asupkeun() kudu ARRAY"}
			}
			arr := args[0].(*Array)
			length := len(arr.Elements)
			newElements := make([]Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]
			return &Array{Elements: newElements}
		},
	},
	
	"kapital": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "kapital() butuh 1 argumen string"}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argumen kapital() kudu STRING"}
			}
			return &String{Value: strings.ToUpper(args[0].(*String).Value)}
		},
	},
	"leutik": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "leutik() butuh 1 argumen string"}
			}
			if args[0].Type() != STRING_OBJ {
				return &Error{Message: "argumen leutik() kudu STRING"}
			}
			return &String{Value: strings.ToLower(args[0].(*String).Value)}
		},
	},

	"kana_angka": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "kana_angka() butuh 1 argumen"}
			}
			switch arg := args[0].(type) {
			case *String:
				val, err := strconv.ParseInt(arg.Value, 0, 64)
				if err != nil {
					return &Error{Message: "gagal konversi string ka angka: " + arg.Value}
				}
				return &Integer{Value: val}
			case *Integer:
				return arg
			default:
				return &Error{Message: "tipe teu didukung keur kana_angka()"}
			}
		},
	},
	"kana_tulisan": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "kana_tulisan() butuh 1 argumen"}
			}
			return &String{Value: args[0].Inspect()}
		},
	},

	"tipe": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return &Error{Message: "tipe() butuh 1 argumen"}
			}
			return &String{Value: string(args[0].Type())}
		},
	},

	"sare": {Fn: fungsiSare},
	"reureuh": {Fn: fungsiSare}, // Sarua keneh jeung sare, ngan ganti ngaran hungkul.
}

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

var (
	TRUE  = &BooleanObject{Value: true}
	FALSE = &BooleanObject{Value: false}
	NULL  = &Null{}
)

func Eval(node Node, env *Environment) Object {
	switch node := node.(type) {
	case *Program:
		return evalProgram(node, env)
	case *BlockStatement:
		return evalBlockStatement(node, env)
	case *ExpressionStatement:
		return Eval(node.Expression, env)
	case *VarStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &ReturnValue{Value: val}
	case *Identifier:
		return evalIdentifier(node, env)
	case *IntegerLiteral:
		return &Integer{Value: node.Value}
	case *Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *StringLiteral:
		return &String{Value: node.Value}
	case *FunctionLiteral:
		return &Function{Parameters: node.Parameters, Env: env, Body: node.Body}
	case *ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &Array{Elements: elements}
	case *HashLiteral:
		return evalHashLiteral(node, env)
	case *IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *CallExpression:
		function := Eval(node.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)
	case *PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)
	case *IfExpression:
		return evalIfExpression(node, env)
	case *WhileStatement:
		return evalWhileStatement(node, env)
	case *PrintStatement:
		val := Eval(node.Expression, env)
		if !isError(val) {
			fmt.Println(val.Inspect())
		}
		return NULL
	case *InputExpression:
		if node.Prompt != "" {
			fmt.Print(node.Prompt)
		}
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
			if result.Type() == ERROR_OBJ {
				return result
			}
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
		if isError(evaluated) {
			return []Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func applyFunction(fn Object, args []Object) Object {
	switch fn := fn.(type) {
	case *Function:
		extendedEnv := NewEnclosedEnvironment(fn.Env)
		for i, param := range fn.Parameters {
			extendedEnv.Set(param.Value, args[i])
		}
		evaluated := Eval(fn.Body, extendedEnv)
		if returnValue, ok := evaluated.(*ReturnValue); ok {
			return returnValue.Value
		}
		return evaluated
	case *Builtin:
		return fn.Fn(args...)
	default:
		return &Error{Message: "lain fungsi (not a function): " + string(fn.Type())}
	}
}

func evalHashLiteral(node *HashLiteral, env *Environment) Object {
	pairs := make(map[HashKey]HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(Hashable)
		if !ok {
			return &Error{Message: "konci teu valid (unusable as hash key): " + string(key.Type())}
		}
		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = HashPair{Key: key, Value: value}
	}
	return &Hash{Pairs: pairs}
}

func evalIndexExpression(left, index Object) Object {
	if left.Type() == ARRAY_OBJ && index.Type() == INTEGER_OBJ {
		return evalArrayIndexExpression(left, index)
	}
	if left.Type() == HASH_OBJ {
		return evalHashIndexExpression(left, index)
	}
	return &Error{Message: fmt.Sprintf("indeks operasi teu didukung: %s", left.Type())}
}

func fungsiSare(args ...Object) Object {
	if len(args) != 1 {
		return &Error{Message: "sare() atawa reureuh() butuh 1 argumen (waktu dina milidetik contohna 1000 artina sadetik)"}
	}
	if args[0].Type() != INTEGER_OBJ {
		return &Error{Message: "argumen waktu kudu INTEGER"}
	}

	ms := args[0].(*Integer).Value
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return NULL
}

func evalArrayIndexExpression(array, index Object) Object {
	arrayObject := array.(*Array)
	idx := index.(*Integer).Value
	max := int64(len(arrayObject.Elements) - 1)
	if idx < 0 || idx > max {
		return NULL
	}
	return arrayObject.Elements[idx]
}

func evalHashIndexExpression(hash, index Object) Object {
	hashObject := hash.(*Hash)
	key, ok := index.(Hashable)
	if !ok {
		return &Error{Message: "konci teu valid (unusable as hash key): " + string(index.Type())}
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}
	return pair.Value
}

func evalIdentifier(node *Identifier, env *Environment) Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	return &Error{Message: "identifier teu kapanggih (not found): " + node.Value}
}

func nativeBoolToBooleanObject(input bool) *BooleanObject {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right Object) Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return &Error{Message: fmt.Sprintf("operator teu dikenal: %s%s", operator, right.Type())}
	}
}

func evalBangOperatorExpression(right Object) Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right Object) Object {
	if right.Type() != INTEGER_OBJ {
		return &Error{Message: fmt.Sprintf("operator teu dikenal: -%s", right.Type())}
	}
	return &Integer{Value: -right.(*Integer).Value}
}

func evalInfixExpression(operator string, left, right Object) Object {
	if operator == "&&" {
		return nativeBoolToBooleanObject(isTruthy(left) && isTruthy(right))
	}
	if operator == "||" {
		return nativeBoolToBooleanObject(isTruthy(left) || isTruthy(right))
	}

	if left.Type() == INTEGER_OBJ && right.Type() == INTEGER_OBJ {
		return evalIntegerInfixExpression(operator, left.(*Integer), right.(*Integer))
	}

	if operator == "+" {
		if left.Type() == STRING_OBJ && right.Type() == STRING_OBJ {
			return &String{Value: left.(*String).Value + right.(*String).Value}
		}
		if left.Type() == STRING_OBJ && right.Type() == INTEGER_OBJ {
			return &String{Value: left.(*String).Value + fmt.Sprintf("%d", right.(*Integer).Value)}
		}
		if left.Type() == INTEGER_OBJ && right.Type() == STRING_OBJ {
			return &String{Value: fmt.Sprintf("%d", left.(*Integer).Value) + right.(*String).Value}
		}
	}


	if operator == "==" {
		return nativeBoolToBooleanObject(left == right)
	}
	if operator == "!=" {
		return nativeBoolToBooleanObject(left != right)
	}

	return &Error{Message: fmt.Sprintf("tipe teu cocok (mismatch): %s %s %s", left.Type(), operator, right.Type())}
}

func evalIntegerInfixExpression(operator string, left, right *Integer) Object {
	lVal := left.Value
	rVal := right.Value
	switch operator {
	case "+":
		return &Integer{Value: lVal + rVal}
	case "-":
		return &Integer{Value: lVal - rVal}
	case "*":
		return &Integer{Value: lVal * rVal}
	case "/":
		if rVal == 0 {
			return &Error{Message: "teu bisa ngabagi jeung nol (division by zero)"}
		}
		return &Integer{Value: lVal / rVal}
	case "%":
		return &Integer{Value: lVal % rVal}
	case "<":
		return nativeBoolToBooleanObject(lVal < rVal)
	case ">":
		return nativeBoolToBooleanObject(lVal > rVal)
	case "<=":
		return nativeBoolToBooleanObject(lVal <= rVal)
	case ">=":
		return nativeBoolToBooleanObject(lVal >= rVal)
	case "==":
		return nativeBoolToBooleanObject(lVal == rVal)
	case "!=":
		return nativeBoolToBooleanObject(lVal != rVal)
	default:
		return &Error{Message: "operator teu dikenal jang integer"}
	}
}

func evalIfExpression(ie *IfExpression, env *Environment) Object {
	condition := Eval(ie.Condition, env)
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	}
	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}
	return NULL
}

func evalWhileStatement(ws *WhileStatement, env *Environment) Object {
	for {
		condition := Eval(ws.Condition, env)
		if isError(condition) {
			return condition
		}
		if !isTruthy(condition) {
			break
		}
		Eval(ws.Body, env)
	}
	return NULL
}

func isTruthy(obj Object) bool {
	if obj == NULL {
		return false
	}
	if obj == TRUE {
		return true
	}
	if obj == FALSE {
		return false
	}
	return true
}

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}