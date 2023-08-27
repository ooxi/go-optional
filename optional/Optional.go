package optional

// A container object which may or may not contain a non-nil value. If a value
// is present, [Optional.IsPresent] returns `true`. If no value is present, the
// object is considered empty and [Optional.IsPresent] returns `false`.
type Optional[T any] struct {
	value *T
}

// Returns an empty [Optional] instance.
func Empty[T any]() Optional[T] {
	return Optional[T]{nil}
}

// Returns an [Optional] describing the given non-nil value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{&value}
}

// Returns an [Optional] describing the given value, if non-nil, otherwise
// returns an empty [Optional].
func OfNilable[T any](value *T) Optional[T] {
	if nil == value {
		return Empty[T]()
	}
	return Of[T](*value)
}

// If a value is not present, returns `true`, otherwise `false`.
func (self *Optional[T]) IsEmpty() bool {
	return nil == self.value
}

// If a value is present, returns `true`, otherwise `false`.
func (self *Optional[T]) IsPresent() bool {
	return nil != self.value
}

// If a value is present, returns the value, otherwise returns `other`.
func (self *Optional[T]) OrElse(other T) T {
	if self.IsEmpty() {
		return other
	}

	return *self.value
}

// If a value is present, returns the value, otherwise panics
// [NoSuchElementError].
func (self *Optional[T]) OrElseThrow() T {
	if !self.IsPresent() {
		panic(NoSuchElementError)
	}

	return *self.value
}
