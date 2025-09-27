package vo

type Vo[T any] interface {
	Value() T
}

type voImpl[T any] struct {
	value T
}

func (v *voImpl[T]) Value() T {
	return v.value
}
