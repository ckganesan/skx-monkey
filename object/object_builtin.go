package object

// BuiltinFunction holds the type of a built-in function.
type BuiltinFunction func(env *Environment, args ...Object) Object

// Builtin wraps func and implements Object interface.
type Builtin struct {
	// Value holds the function we wrap.
	Fn BuiltinFunction
}

// Type returns the type of this object.
func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

// Inspect returns a string-representation of the given object.
func (b *Builtin) Inspect() string {
	return "builtin function"
}

// InvokeMethod invokes a method against the object.
// (Built-in methods only.)
func (b *Builtin) InvokeMethod(method string, env Environment, args ...Object) Object {
	if method == "methods" {
		names := []string{"methods", "type"}

		result := make([]Object, len(names), len(names))
		for i, txt := range names {
			result[i] = &String{Value: txt}
		}
		return &Array{Elements: result}
	}
	if method == "type" {
		return &String{Value: "builtin"}
	}
	return nil
}
