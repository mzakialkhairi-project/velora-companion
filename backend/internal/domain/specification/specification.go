// Package specification provides specification pattern for the domain layer.
package specification

import (
	"context"
)

// Specification is the interface for the specification pattern
type Specification[T any] interface {
	// IsSatisfiedBy checks if the given entity satisfies the specification
	IsSatisfiedBy(ctx context.Context, entity T) (bool, error)

	// And combines this specification with another using AND logic
	And(other Specification[T]) Specification[T]

	// Or combines this specification with another using OR logic
	Or(other Specification[T]) Specification[T]

	// Not negates this specification
	Not() Specification[T]
}

// AndSpecification combines two specifications with AND logic
type AndSpecification[T any] struct {
	left, right Specification[T]
}

// NewAndSpecification creates a new AndSpecification
func NewAndSpecification[T any](left, right Specification[T]) *AndSpecification[T] {
	return &AndSpecification[T]{left: left, right: right}
}

// IsSatisfiedBy implements Specification
func (s *AndSpecification[T]) IsSatisfiedBy(ctx context.Context, entity T) (bool, error) {
	leftSatisfied, err := s.left.IsSatisfiedBy(ctx, entity)
	if err != nil {
		return false, err
	}
	if !leftSatisfied {
		return false, nil
	}

	return s.right.IsSatisfiedBy(ctx, entity)
}

// And implements Specification
func (s *AndSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification(s, other)
}

// Or implements Specification
func (s *AndSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification(s, other)
}

// Not implements Specification
func (s *AndSpecification[T]) Not() Specification[T] {
	return NewNotSpecification(s)
}

// OrSpecification combines two specifications with OR logic
type OrSpecification[T any] struct {
	left, right Specification[T]
}

// NewOrSpecification creates a new OrSpecification
func NewOrSpecification[T any](left, right Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{left: left, right: right}
}

// IsSatisfiedBy implements Specification
func (s *OrSpecification[T]) IsSatisfiedBy(ctx context.Context, entity T) (bool, error) {
	leftSatisfied, err := s.left.IsSatisfiedBy(ctx, entity)
	if err != nil {
		return false, err
	}
	if leftSatisfied {
		return true, nil
	}

	return s.right.IsSatisfiedBy(ctx, entity)
}

// And implements Specification
func (s *OrSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification(s, other)
}

// Or implements Specification
func (s *OrSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification(s, other)
}

// Not implements Specification
func (s *OrSpecification[T]) Not() Specification[T] {
	return NewNotSpecification(s)
}

// NotSpecification negates a specification
type NotSpecification[T any] struct {
	spec Specification[T]
}

// NewNotSpecification creates a new NotSpecification
func NewNotSpecification[T any](spec Specification[T]) *NotSpecification[T] {
	return &NotSpecification[T]{spec: spec}
}

// IsSatisfiedBy implements Specification
func (s *NotSpecification[T]) IsSatisfiedBy(ctx context.Context, entity T) (bool, error) {
	satisfied, err := s.spec.IsSatisfiedBy(ctx, entity)
	if err != nil {
		return false, err
	}
	return !satisfied, nil
}

// And implements Specification
func (s *NotSpecification[T]) And(other Specification[T]) Specification[T] {
	return NewAndSpecification(s, other)
}

// Or implements Specification
func (s *NotSpecification[T]) Or(other Specification[T]) Specification[T] {
	return NewOrSpecification(s, other)
}

// Not implements Specification
func (s *NotSpecification[T]) Not() Specification[T] {
	return s.spec
}

// TrueSpecification always returns true
type TrueSpecification[T any] struct{}

// NewTrueSpecification creates a new TrueSpecification
func NewTrueSpecification[T any]() *TrueSpecification[T] {
	return &TrueSpecification[T]{}
}

// IsSatisfiedBy implements Specification
func (s *TrueSpecification[T]) IsSatisfiedBy(ctx context.Context, entity T) (bool, error) {
	return true, nil
}

// And implements Specification
func (s *TrueSpecification[T]) And(other Specification[T]) Specification[T] {
	return other
}

// Or implements Specification
func (s *TrueSpecification[T]) Or(other Specification[T]) Specification[T] {
	return s
}

// Not implements Specification
func (s *TrueSpecification[T]) Not() Specification[T] {
	return NewFalseSpecification[T]()
}

// FalseSpecification always returns false
type FalseSpecification[T any] struct{}

// NewFalseSpecification creates a new FalseSpecification
func NewFalseSpecification[T any]() *FalseSpecification[T] {
	return &FalseSpecification[T]{}
}

// IsSatisfiedBy implements Specification
func (s *FalseSpecification[T]) IsSatisfiedBy(ctx context.Context, entity T) (bool, error) {
	return false, nil
}

// And implements Specification
func (s *FalseSpecification[T]) And(other Specification[T]) Specification[T] {
	return s
}

// Or implements Specification
func (s *FalseSpecification[T]) Or(other Specification[T]) Specification[T] {
	return other
}

// Not implements Specification
func (s *FalseSpecification[T]) Not() Specification[T] {
	return NewTrueSpecification[T]()
}
