package gconcat

import "github.com/jeffotoni/gconcat"

// Concat encapsulamento da lib gconcat
func Concat(values ...interface{}) string {
	return gconcat.Build(values)
}
