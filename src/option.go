package safego

// Option is a type that represents an optional value.
type Option[T any] struct {
	Val    T
	Exists bool
}

// NewOption creates a new Option[T] with the given value.
func NewOption[T any](val T) *Option[T] {
	return &Option[T]{
		Val:    val,
		Exists: true,
	}
}

// IsSome returns true if the option is a Some value.
func (o *Option[T]) IsSome() bool {
	return o.Exists
}

// IsNone returns true if the option is a None value.
func (o *Option[T]) IsNone() bool {
	return !o.Exists
}

// Expect panics with the given message if the option is a None value.
func (o *Option[T]) Expect(msg string) T {
	if o.Exists {
		return o.Val
	}

	panic(msg)
}

// Unwrap panics if the option is a None value.
func (o *Option[T]) Unwrap() T {
	if o.Exists {
		return o.Val
	}

	panic("unwrap a none option")
}

// UnwrapOr returns the value of the option if it is a Some value, otherwise it returns the given default value.
func (o *Option[T]) UnwrapOr(def T) T {
	if o.Exists {
		return o.Val
	}

	return def
}

// UnwrapOrElse returns the value of the option if it is a Some value, otherwise it returns the result of the given function.
func (o *Option[T]) UnwrapOrElse(fn func() T) T {
	if o.Exists {
		return o.Val
	}

	return fn()
}
