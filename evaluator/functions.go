package evaluator

import (
	"pomidor/builtins"
	"pomidor/object"
)

func GetBuiltins() map[string]*object.Builtin {
	var b = map[string]*object.Builtin{}

	for name, fn := range builtins.CommonFunctions {
		b[name] = fn
	}
	for name, fn := range builtins.ArrayFunctions {
		b[name] = fn
	}

	return b
}
