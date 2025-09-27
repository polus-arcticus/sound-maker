// Package songs provides a collection of musical compositions
package songs

import (
	"fmt"
)

// Composition defines the interface for a musical composition
type Composition interface {
	Play() error
	GetName() string
	GetDescription() string
}

// CompositionFunc is a function type that creates a composition
type CompositionFunc func() Composition

// Registry stores all available compositions
var Registry = make(map[string]CompositionFunc)

// RegisterComposition adds a composition to the registry
func RegisterComposition(name string, createFn CompositionFunc) {
	Registry[name] = createFn
}

// ListCompositions returns a list of all available composition names
func ListCompositions() []string {
	var names []string
	for name := range Registry {
		names = append(names, name)
	}
	return names
}

// GetComposition returns a composition by name
func GetComposition(name string) (Composition, error) {
	createFn, exists := Registry[name]
	if !exists {
		return nil, fmt.Errorf("composition %q not found", name)
	}
	return createFn(), nil
}
