// Package module provides some simple functions
package module

import "github.com/cheatsnake/module/internal/private"

// Add adds two integers
func Add(a, b int) int {
	return a + b
}

// Multiply multiplies two integers using the private module
func Multiply(a, b int) int {
	// private module is accessible only in this level
	return private.Multiply(a, b)
}
