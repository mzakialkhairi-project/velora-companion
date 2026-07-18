// Package provider provides AI provider abstraction layer.
package provider

import (
	"fmt"
	"sync"
)

// globalRegistry is the global provider registry instance.
var globalRegistry = &Registry{
	providers: make(map[string]Provider),
}

// Registry manages AI provider instances.
type Registry struct {
	mu        sync.RWMutex
	providers map[string]Provider
	defaultP  string
}

// Register registers a provider with the global registry.
func Register(name string, p Provider) {
	globalRegistry.Register(name, p)
}

// Register registers a provider with the registry.
func (r *Registry) Register(name string, p Provider) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if name == "" {
		panic("provider name cannot be empty")
	}
	if p == nil {
		panic("provider cannot be nil")
	}

	r.providers[name] = p

	// If this is the first provider, set it as default
	if len(r.providers) == 1 {
		r.defaultP = name
	}
}

// Get retrieves a provider by name from the global registry.
func Get(name string) (Provider, bool) {
	return globalRegistry.Get(name)
}

// Get retrieves a provider by name from the registry.
func (r *Registry) Get(name string) (Provider, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.providers[name]
	return p, ok
}

// MustGet retrieves a provider by name from the global registry.
// It panics if the provider is not found.
func MustGet(name string) Provider {
	p, ok := globalRegistry.Get(name)
	if !ok {
		panic(fmt.Sprintf("provider not found: %s", name))
	}
	return p
}

// List returns a list of registered provider names from the global registry.
func List() []string {
	return globalRegistry.List()
}

// List returns a list of registered provider names.
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.providers))
	for name := range r.providers {
		names = append(names, name)
	}
	return names
}

// Default returns the default provider from the global registry.
func Default() (Provider, bool) {
	return globalRegistry.Default()
}

// Default returns the default provider.
func (r *Registry) Default() (Provider, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.defaultP == "" {
		return nil, false
	}

	p, ok := r.providers[r.defaultP]
	return p, ok
}

// SetDefault sets the default provider by name.
func SetDefault(name string) {
	globalRegistry.SetDefault(name)
}

// SetDefault sets the default provider by name.
func (r *Registry) SetDefault(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.providers[name]; !ok {
		panic(fmt.Sprintf("provider not found: %s", name))
	}

	r.defaultP = name
}

// Reset clears all registered providers (useful for testing).
func Reset() {
	globalRegistry.Reset()
}

// Reset clears all registered providers.
func (r *Registry) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.providers = make(map[string]Provider)
	r.defaultP = ""
}
