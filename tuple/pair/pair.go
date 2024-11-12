package pair

import "fmt"

type Pair[K any, V any] struct {
	Key   K
	Value V
}

func (pair *Pair[K, V]) String() string {
	return fmt.Sprintf()
}
