// Package engine provides AI chat orchestration.
package engine

import (
	"fmt"

	"github.com/mzakiaklhairi/velora/internal/ai/provider"
)

// DefaultProviderResolver resolves the default provider from registry.
type DefaultProviderResolver struct{}

// NewDefaultProviderResolver creates a new DefaultProviderResolver.
func NewDefaultProviderResolver() *DefaultProviderResolver {
	return &DefaultProviderResolver{}
}

// Resolve returns the default provider from the registry.
func (r *DefaultProviderResolver) Resolve() (provider.Provider, error) {
	p, ok := provider.Default()
	if !ok {
		return nil, fmt.Errorf("%w: no default provider registered", ErrProviderNotAvailable)
	}
	return p, nil
}
